package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/julian776/baby-guardian/analytics/internal/analyzer"
	"github.com/julian776/baby-guardian/analytics/internal/monitor"
	"github.com/julian776/baby-guardian/analytics/pkg/alerts"
	"github.com/julian776/baby-guardian/monitor/pkg/streamers"
	pb "github.com/julian776/baby-guardian/protos"
	"github.com/rs/zerolog"
)

func main() {
	rootCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Msg("analyzing baby data")

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

	err := monitor.Start(rootCtx)
	if err != nil {
		logger.Panic().Err(err).Msg("failed to start monitor")
	}

	<-rootCtx.Done()
}
