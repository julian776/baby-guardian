package sensors

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/julian776/baby-guardian/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Audio struct {
	name     string
	interval time.Duration
}

func NewAudio(name string, interval time.Duration) *Audio {
	return &Audio{
		name:     name,
		interval: interval,
	}
}

func (a *Audio) Name() string {
	return a.name
}

func (a *Audio) Interval() time.Duration {
	return a.interval
}

func (a *Audio) Stop(ctx context.Context) error {
	return nil
}

func (a *Audio) Start(ctx context.Context) (<-chan *pb.Signal, error) {
	signalChan := make(chan *pb.Signal)

	go func() {
		timer := time.NewTicker(a.interval)
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				close(signalChan)
				return

			case now := <-timer.C:
				signalChan <- &pb.Signal{
					Type:      pb.Type_AUDIO,
					Timestamp: timestamppb.New(now),
					Value:     a.GenerateAudioData(),
				}
			}
		}
	}()

	return signalChan, nil
}

func (a *Audio) GenerateAudioData() float64 {
	multiplier := 1.0

	isCrying := rand.Intn(100) <= 2
	if isCrying {
		multiplier = 20.0
	}

	return rand.Float64() * multiplier
}
