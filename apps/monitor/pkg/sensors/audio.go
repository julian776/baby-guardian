package sensors

import (
	"context"
	"math/rand"
	"time"
)

type Audio struct {
	name     string
	interval time.Duration
}

func (a *Audio) Name() string {
	return a.name
}

func (a *Audio) Interval() time.Duration {
	return a.interval
}

func (a *Audio) Start(context.Context) (<-chan Signal, error) {
	signalChan := make(chan Signal)

	go func() {
		timer := time.NewTicker(a.interval)
		for {
			now := <-timer.C
			signalChan <- Signal{
				Type:      AudioTyp.String(),
				Timestamp: now,
				Value:     a.GenerateAudioData(),
			}
		}
	}()

	return signalChan, nil
}

func (a *Audio) GenerateAudioData() float64 {
	multiplier := 1.0

	isCrying := rand.Intn(100) <= 2
	if isCrying {
		multiplier = 10.0
	}

	return rand.Float64() * multiplier
}
