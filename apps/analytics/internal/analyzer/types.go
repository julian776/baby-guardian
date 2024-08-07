package analyzer

import "fmt"

type Threshold struct {
	Min float64
	Max float64
}

type Result struct {
	DangerousVariation float64
	DangerousValue     float64
	Threshold          Threshold
}

func (r Result) String() string {
	msg := ""
	if r.DangerousVariation != 0 {
		msg += fmt.Sprintf("dangerous variation detected, value: %.2f\n", r.DangerousVariation)
	}

	if r.DangerousValue != 0 {
		msg += fmt.Sprintf("dangerous value detected, value: %.2f out of range [%.2f, %.2f]\n", r.DangerousValue, r.Threshold.Min, r.Threshold.Max)
	}

	return msg
}

func (r Result) IsDangerous() bool {
	return r.DangerousVariation != 0 || r.DangerousValue != 0
}
