package main

import (
	"context"
	"os"
	"strings"

	"github.com/julian776/baby-guardian/analytics/internal/monitor"
	pb "github.com/julian776/baby-guardian/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrUnauthenticated = status.Error(codes.Unauthenticated, "unauthenticated")
	AuthToken          = os.Getenv("AUTH_TOKEN")
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

func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrUnauthenticated
	}

	if len(m["authorization"]) == 0 {
		return nil, ErrUnauthenticated
	}

	s := strings.TrimPrefix(m["authorization"][0], "Bearer ")
	if s != AuthToken {
		return nil, ErrUnauthenticated
	}

	return handler(ctx, req)
}
