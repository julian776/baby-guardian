package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/julian776/baby-guardian/monitor/internal/monitor"
	"github.com/julian776/baby-guardian/monitor/pkg/sensors"
	"github.com/julian776/baby-guardian/monitor/pkg/streamers"
	"github.com/rs/zerolog"
)

func main() {
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

	monitor := monitor.NewMonitor(
		streamer,
		&logger,
	)
	monitor.AddSensors(
		sensors.NewAudio(
			"audio",
			time.Second*5,
		),
		sensors.NewTemperature(
			"temperature",
			time.Second*5,
		),
	)

	err := monitor.Start(rootCtx)
	if err != nil {
		logger.Panic().Err(err).Msg("failed to start monitor")
	}

	<-rootCtx.Done()
}
