package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/julian776/baby-guardian/monitor/pkg/sensors"
	"github.com/julian776/baby-guardian/monitor/pkg/streamers"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

type Monitor struct {
	sensors  []sensors.Sensor
	streamer streamers.Streamer
	logger   *zerolog.Logger
}

func NewMonitor(
	streamer streamers.Streamer,
	logger *zerolog.Logger,
) *Monitor {
	return &Monitor{
		sensors:  make([]sensors.Sensor, 0),
		streamer: streamer,
		logger:   logger,
	}
}

func (m *Monitor) AddSensors(sensors ...sensors.Sensor) {
	m.sensors = append(m.sensors, sensors...)
}

func (m *Monitor) Start(ctx context.Context) error {
	if s, ok := m.streamer.(streamers.Starter); ok {
		if err := s.Start(ctx); err != nil {
			return fmt.Errorf("failed to start streamer: %w", err)
		}
	}

	for _, sensor := range m.sensors {
		if err := m.monitorSensor(ctx, sensor); err != nil {
			return err
		}
	}

	return nil
}

func (m *Monitor) monitorSensor(ctx context.Context, sensor sensors.Sensor) error {
	// Create a new context for the sensor restart
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	signalChan, err := sensor.Start(ctx)
	if err != nil {
		return fmt.Errorf("failed to start sensor %s: %w", sensor.Name(), err)
	}

	interval := sensor.Interval()

	for {
		timeout := time.After(interval + time.Second)

		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-timeout:
			fmt.Println("Restarting sensor")
			cancel()

			ctx, cancel = context.WithCancel(ctx)
			defer cancel()

			err := sensor.Stop(ctx)
			if err != nil {
				return fmt.Errorf("failed to stop sensor %s: %w", sensor.Name(), err)
			}

			signalChan, err = sensor.Start(ctx)
			if err != nil {
				return fmt.Errorf("failed to restart sensor %s: %w", sensor.Name(), err)
			}

		case signal, ok := <-signalChan:
			if !ok {
				return fmt.Errorf("sensor %s closed signal channel", sensor.Name())
			}

			d, err := proto.Marshal(signal)
			if err != nil {
				m.logger.Error().Msgf("failed to marshal signal: %s", err.Error())
				continue
			}

			// TODO: Retry on error
			err = m.streamer.PutRecord(ctx, d)
			if err != nil {
				m.logger.Error().Msgf("failed to write signal to streamer: %s", err.Error())
				continue
			}
		}
	}
}
