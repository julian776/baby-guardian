package sensors

type SignalType uint8

const (
	AudioTyp SignalType = iota
	TemperatureTyp
	HeartRateTyp
)

func (s SignalType) String() string {
	switch s {
	case AudioTyp:
		return "Audio"
	case TemperatureTyp:
		return "Temperature"
	case HeartRateTyp:
		return "HeartRate"
	default:
		return "Unknown"
	}
}
