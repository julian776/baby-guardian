package sensors

import "time"

type Signal struct {
	Type      string
	Timestamp time.Time
	Value     float64
}
