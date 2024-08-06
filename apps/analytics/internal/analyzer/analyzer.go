package analyzer

import (
	"math"

	pb "github.com/julian776/baby-guardian/protos"
	"github.com/rs/zerolog"
)

type Analyzer struct {
	thresholds  map[pb.Type]Threshold
	lastSignals map[pb.Type]*pb.Signal
	logger      *zerolog.Logger
}

func NewAnalyzer(
	logger *zerolog.Logger,
) *Analyzer {
	return &Analyzer{
		thresholds:  make(map[pb.Type]Threshold),
		lastSignals: make(map[pb.Type]*pb.Signal),
		logger:      logger,
	}
}

func (a *Analyzer) AddThreshold(typ pb.Type, t Threshold) {
	a.thresholds[typ] = t
}

func (a *Analyzer) Analyze(signal *pb.Signal) Result {
	r := Result{}
	t, ok := a.thresholds[signal.Type]
	if !ok {
		a.logger.Warn().Msgf("no threshold for signal type %s", signal.Type)
		return r
	}

	if signal.Value < t.Min || signal.Value > t.Max {
		r.DangerousValue = signal.Value
		r.Threshold = t
	}

	if a.hasDangerousVariation(a.lastSignals[signal.Type], signal) {
		r.DangerousVariation = signal.Value
	}

	a.lastSignals[signal.Type] = signal
	return r
}

func (a *Analyzer) hasDangerousVariation(
	lastSignal *pb.Signal,
	signal *pb.Signal,
) bool {
	if lastSignal == nil {
		return false
	}

	variation := math.Abs(signal.Value - lastSignal.Value)
	if variation < 0.5 {
		return false
	}

	fortyPercent := lastSignal.Value * 0.2

	return variation > fortyPercent
}
