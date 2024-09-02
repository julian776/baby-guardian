package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/julian776/baby-guardian/analytics/internal/analyzer"
	"github.com/julian776/baby-guardian/analytics/internal/monitor"
	"github.com/julian776/baby-guardian/analytics/pkg/alerts"
	"github.com/julian776/baby-guardian/analytics/pkg/auth"
	"github.com/julian776/baby-guardian/monitor/pkg/streamers"
	pb "github.com/julian776/baby-guardian/protos"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.String("port", "8080", "port to listen to")
)

func main() {
	flag.Parse()

	rootCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	streamer := streamers.NewKinesis(
		&streamers.KinesisConfig{
			StreamName: "baby-guardian",
			Region:     "us-east-1",
			StreamMode: types.StreamModeOnDemand,
		},
		&logger,
	)

	a := analyzer.NewAnalyzer(
		&logger,
	)
	a.AddThreshold(
		pb.Type_TEMPERATURE,
		analyzer.Threshold{
			Min: 15,
			Max: 22,
		},
	)
	a.AddThreshold(
		pb.Type_AUDIO,
		analyzer.Threshold{
			Min: 0,
			Max: 10,
		},
	)

	monitor := monitor.NewMonitor(
		streamer,
		a,
		&logger,
		alerts.NewConsole(&logger),
	)

	errGroup, ctx := errgroup.WithContext(rootCtx)

	errGroup.Go(func() error {
		err := monitor.Start(ctx)
		if err != nil {
			return fmt.Errorf("failed to start monitor: %w", err)
		}

		return nil
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		logger.Panic().Err(err).Msg("failed to listen")
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(auth.AuthUnaryInterceptor),
	)
	analyticsServer := NewAnalyticsServer(monitor)
	pb.RegisterAnalyticsServer(srv, analyticsServer)
	authServer := auth.NewAuthServer()
	pb.RegisterAuthServer(srv, authServer)
	reflection.Register(srv)

	errGroup.Go(func() error {
		logger.Info().Msgf("listening on port %s", *port)
		if err := srv.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve: %w", err)
		}

		return nil
	})

	<-ctx.Done()
	logger.Info().Msg("shutting down")
	srv.GracefulStop()

	if err := errGroup.Wait(); err != nil {
		logger.Error().Err(err).Msg("stopped with error")
	}
}
