package main

import (
	"context"
	"time"

	"github.com/julian776/baby-guardian/analytics/internal/monitor"
	pb "github.com/julian776/baby-guardian/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type AnalyticsServer struct {
	monitor *monitor.Monitor
	pb.UnimplementedAnalyticsServer
}

func NewAnalyticsServer(monitor *monitor.Monitor) *AnalyticsServer {
	s := &AnalyticsServer{
		monitor: monitor,
	}

	return s
}

func (s *AnalyticsServer) LastDangerousSignal(
	ctx context.Context,
	req *pb.LastDangerousSignalRequest,
) (*pb.LastDangerousSignalResponse, error) {
	signal := s.monitor.LastDangerousSignal()
	if signal == nil {
		return nil, status.Error(codes.NotFound, "no dangerous signal found")
	}

	return &pb.LastDangerousSignalResponse{
		Signal: signal,
	}, nil
}

func (s *AnalyticsServer) LastDangerousSignalStream(
	req *pb.LastDangerousSignalStreamRequest,
	stream pb.Analytics_LastDangerousSignalStreamServer,
) error {
	interval := req.GetInterval().AsDuration()
	var lastSignal *pb.Signal

	// TODO: add a way to stop the stream due grpc-gateway limitations
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
		}

		signal := s.monitor.LastDangerousSignal()
		if signal == nil {
			time.Sleep(interval)
			continue
		}

		if proto.Equal(lastSignal, signal) {
			time.Sleep(interval)
			continue
		}

		err := stream.Send(&pb.LastDangerousSignalResponse{
			Signal: signal,
		})
		if err != nil {
			return err
		}

		lastSignal = signal
		time.Sleep(interval)
	}
}
