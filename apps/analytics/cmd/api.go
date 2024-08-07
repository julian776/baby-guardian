package main

import (
	"context"

	"github.com/julian776/baby-guardian/analytics/internal/monitor"
	pb "github.com/julian776/baby-guardian/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
