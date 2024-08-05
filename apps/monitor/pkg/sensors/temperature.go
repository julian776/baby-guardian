package sensors

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/julian776/baby-guardian/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Temperature struct {
	name     string
	interval time.Duration
}

func NewTemperature(name string, interval time.Duration) *Temperature {
	return &Temperature{
		name:     name,
		interval: interval,
	}
}

func (t *Temperature) Name() string {
	return t.name
}

func (t *Temperature) Interval() time.Duration {
	return t.interval
}

func (t *Temperature) Start(ctx context.Context) (<-chan *pb.Signal, error) {
	signalChan := make(chan *pb.Signal)

	go func() {
		ticker := time.NewTicker(t.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				close(signalChan)
				return
			case now := <-ticker.C:
				signalChan <- &pb.Signal{
					Type:      pb.Type_TEMPERATURE,
					Timestamp: timestamppb.New(now),
					Value:     t.generateTemperature(),
				}
			}
		}
	}()

	return signalChan, nil
}

func (t *Temperature) generateTemperature() float64 {
	multiplier := 1
	if rand.Intn(100) < 2 {
		multiplier = 2
	}

	return (rand.Float64() + 16) * float64(multiplier)
}

func (t *Temperature) Stop(ctx context.Context) error {
	return nil
}
