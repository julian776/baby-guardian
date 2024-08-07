package monitor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/julian776/baby-guardian/analytics/internal/analyzer"
	"github.com/julian776/baby-guardian/analytics/pkg/alerts"
	"github.com/julian776/baby-guardian/monitor/pkg/streamers"
	pb "github.com/julian776/baby-guardian/protos"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

type Monitor struct {
	streamer streamers.Streamer
	analyzer *analyzer.Analyzer

	lock                *sync.RWMutex
	lastDangerousSignal *pb.Signal

	alerts []alerts.Alerter
	logger *zerolog.Logger
}

func NewMonitor(
	streamer streamers.Streamer,
	analyzer *analyzer.Analyzer,
	logger *zerolog.Logger,
	alerts ...alerts.Alerter,
) *Monitor {
	return &Monitor{
		streamer: streamer,
		analyzer: analyzer,
		lock:     &sync.RWMutex{},
		alerts:   alerts,
		logger:   logger,
	}
}

func (m *Monitor) Start(ctx context.Context) error {
	if s, ok := m.streamer.(streamers.Starter); ok {
		if err := s.Start(ctx); err != nil {
			return fmt.Errorf("failed to start streamer: %w", err)
		}
	}

	return m.consumue(ctx)
}

func (m *Monitor) consumue(ctx context.Context) error {
	// Create a new context for the sensor restart
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		recs, err := m.streamer.GetRecords(ctx)
		if err != nil {
			return fmt.Errorf("failed to get records: %w", err)
		}

		for _, rec := range recs {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			signal := &pb.Signal{}
			err := proto.Unmarshal(rec, signal)
			if err != nil {
				m.logger.Error().Err(err).Msg("failed to unmarshal signal")
				continue
			}

			m.logger.Info().Msgf("received signal: %s", signal.String())

			r := m.analyzer.Analyze(signal)
			if r.IsDangerous() {
				m.alert(r.String())
				m.lock.Lock()
				m.lastDangerousSignal = signal
				m.lock.Unlock()
			}
		}

		time.Sleep(1 * time.Second)
	}

}

func (m *Monitor) LastDangerousSignal() *pb.Signal {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.lastDangerousSignal
}

func (m *Monitor) alert(message string) {
	for _, alert := range m.alerts {
		err := alert.Alert(message)
		if err != nil {
			m.logger.Error().Err(err).Msg("failed to alert: " + message)
		}
	}
}
