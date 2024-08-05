package analayzer

import (
	"fmt"
	"math"

	"github.com/julian776/baby-guardian/analytics/pkg/alerts"
	pb "github.com/julian776/baby-guardian/protos"
	"github.com/rs/zerolog"
)

type Analyzer struct {
	thresholds  map[pb.Type]Threshold
	lastSignals map[pb.Type]*pb.Signal
	alerts      []alerts.Alerter
	logger      *zerolog.Logger
}

func NewAnalyzer(
	logger *zerolog.Logger,
	alerts ...alerts.Alerter,
) *Analyzer {
	return &Analyzer{
		thresholds:  make(map[pb.Type]Threshold),
		lastSignals: make(map[pb.Type]*pb.Signal),
		alerts:      alerts,
		logger:      logger,
	}
}

func (a *Analyzer) AddThreshold(typ pb.Type, t Threshold) {
	a.thresholds[typ] = t
}

func (a *Analyzer) Analyze(signal *pb.Signal) {
	t := a.thresholds[signal.Type]

	if signal.Value < t.Min {
		a.alert(fmt.Sprintf("value below threshold detected, value: %f, min: %f for type: %s", signal.Value, t.Min, signal.Type.String()))
	}

	if signal.Value > t.Max {
		a.alert(fmt.Sprintf("value above threshold detected, value: %f, max: %f for type: %s", signal.Value, t.Max, signal.Type.String()))
	}

	if a.hasDangerousVariation(a.lastSignals[signal.Type], signal) {
		a.alert(fmt.Sprintf("dangerous variation detected, value: %f for type: %s", signal.Value, signal.Type.String()))
	}

	a.lastSignals[signal.Type] = signal
}

func (a *Analyzer) hasDangerousVariation(
	lastSignal *pb.Signal,
	signal *pb.Signal,
) bool {
	if lastSignal == nil {
		return false
	}

	variation := math.Abs(signal.Value - lastSignal.Value)
	twentyPercent := lastSignal.Value * 0.2

	return variation > twentyPercent
}

func (a *Analyzer) alert(message string) {
	for _, alert := range a.alerts {
		err := alert.Alert(message)
		if err != nil {
			a.logger.Error().Err(err).Msg("failed to alert: " + message)
		}
	}
}
