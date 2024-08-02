package monitor

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/julian776/baby-guardian/monitor/pkg/sensors"
	sensorsmocks "github.com/julian776/baby-guardian/monitor/pkg/sensors/mocks"
	streamersmocks "github.com/julian776/baby-guardian/monitor/pkg/streamers/mocks"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMonitor_monitorSensor(t *testing.T) {
	interval := time.Millisecond * 10
	ctx, cancel := context.WithTimeout(context.Background(), interval*2)
	defer cancel()

	smock := streamersmocks.NewMockStreamer(t)
	smock.EXPECT().PutRecord(
		mock.Anything,
		mock.AnythingOfType("[]uint8"),
	).Return(nil)

	b := &strings.Builder{}
	logger := log.Output(b)
	m := &Monitor{
		sensors:  []sensors.Sensor{},
		streamer: smock,
		logger:   &logger,
	}

	err := m.monitorSensor(ctx, sensors.NewAudio("audio", interval))
	assert.Error(t, err, "Monitorer.monitorSensor() did not return context.DeadlineExceeded error")

	smock.AssertCalled(t, "PutRecord", mock.Anything, mock.AnythingOfType("[]uint8"))
}

func TestMonitor_monitorSensor_StreamError(t *testing.T) {
	interval := time.Millisecond * 10
	ctx, cancel := context.WithTimeout(context.Background(), interval*2)
	defer cancel()

	smock := streamersmocks.NewMockStreamer(t)
	smock.EXPECT().PutRecord(
		mock.Anything,
		mock.AnythingOfType("[]uint8"),
	).Return(fmt.Errorf("streamer error"))

	b := &strings.Builder{}
	logger := log.Output(b)
	m := &Monitor{
		sensors:  []sensors.Sensor{},
		streamer: smock,
		logger:   &logger,
	}

	err := m.monitorSensor(ctx, sensors.NewAudio("audio", interval))
	assert.Error(t, err, "expected error from streamer")
	assert.Contains(t, b.String(), "failed to write signal to streamer:", "expected error message in log")
	assert.Contains(t, b.String(), "level\":\"error", "expected error message in log")

	smock.AssertCalled(t, "PutRecord", mock.Anything, mock.AnythingOfType("[]uint8"))
}
func TestMonitor_monitorSensor_Timeout(t *testing.T) {
	interval := time.Millisecond * 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	streamer := streamersmocks.NewMockStreamer(t)

	sensor := sensorsmocks.NewMockSensor(t)
	sensor.EXPECT().Start(
		mock.Anything,
	).Return(nil, nil)
	sensor.EXPECT().Stop(
		mock.Anything,
	).Return(nil)
	sensor.EXPECT().Interval().Return(interval)

	b := &strings.Builder{}
	logger := log.Output(b)
	m := &Monitor{
		sensors:  []sensors.Sensor{},
		streamer: streamer,
		logger:   &logger,
	}

	err := m.monitorSensor(ctx, sensor)
	assert.Error(t, err, "Monitorer.monitorSensor() did not return context.DeadlineExceeded error")

	sensor.AssertCalled(t, "Start", mock.Anything)
	sensor.AssertCalled(t, "Stop", mock.Anything)
}
