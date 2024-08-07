package monitor

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/julian776/baby-guardian/analytics/internal/analyzer"
	"github.com/julian776/baby-guardian/analytics/pkg/alerts"
	streamers "github.com/julian776/baby-guardian/monitor/pkg/streamers/mocks"
	pb "github.com/julian776/baby-guardian/protos"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestMonitor_StartAndConsume(t *testing.T) {
	signals, err := encodeMockSignals(
		&pb.Signal{Type: pb.Type_TEMPERATURE, Value: 18},
		&pb.Signal{Type: pb.Type_TEMPERATURE, Value: 20},
		&pb.Signal{Type: pb.Type_AUDIO, Value: 0.8},
		&pb.Signal{Type: pb.Type_AUDIO, Value: 0.88},
		&pb.Signal{Type: pb.Type_TEMPERATURE, Value: 38}, // dangerous temperature
	)
	if err != nil {
		t.Fatalf("Failed to encode signals: %v", err)
	}

	mockStreamer := &streamers.MockStreamer{}
	mockStreamer.EXPECT().GetRecords(
		mock.Anything,
	).Return(signals, nil)

	b := &strings.Builder{}
	logger := zerolog.New(b)
	a := analyzer.NewAnalyzer(&logger)
	a.AddThreshold(
		pb.Type_TEMPERATURE,
		analyzer.Threshold{Min: 15, Max: 22},
	)
	a.AddThreshold(
		pb.Type_AUDIO,
		analyzer.Threshold{Min: 0.0, Max: 1.8},
	)

	bAlerts := &strings.Builder{}
	loggerAlert := zerolog.New(bAlerts)
	alert := alerts.NewConsole(&loggerAlert)

	monitor := NewMonitor(mockStreamer, a, &logger, alert)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Start the Monitor
	go func() {
		if err := monitor.Start(ctx); err != nil {
			t.Errorf("Start() returned an error: %v", err)
		}
	}()

	// Wait for the Monitor to start consuming
	time.Sleep(100 * time.Millisecond)

	// Stop the Monitor
	cancel()

	assert.NotZero(t, monitor.LastDangerousSignal(), "lastDangerousSignal should not be zero")
	assert.Equal(t, pb.Type_TEMPERATURE, monitor.LastDangerousSignal().Type, "lastDangerousSignal.Type should be pb.Type_TEMPERATURE")
	assert.Equal(t, float64(38), monitor.LastDangerousSignal().Value, "lastDangerousSignal.Value should be 38")

	assert.Contains(t, bAlerts.String(), "dangerous value detected, value: 38.00 out of range [15.00, 22.00]")
}

func encodeMockSignals(signals ...*pb.Signal) ([][]byte, error) {
	byteSignals := make([][]byte, len(signals))
	for i, signal := range signals {
		b, err := proto.Marshal(signal)
		if err != nil {
			return nil, err
		}
		byteSignals[i] = b
	}

	return byteSignals, nil
}
