package sensors

import (
	"context"
	"time"

	pb "github.com/julian776/baby-guardian/protos"
)

type Sensor interface {
	Name() string
	Interval() time.Duration

	// Start starts the sensor and returns a channel that will receive
	// signals from the sensor. The channel must be closed when the sensor
	// is stopped.
	//
	// A sensor must send at least one signal per interval. If no signals
	// are sent, the sensor is considered to be in an error state.
	Start(context.Context) (<-chan *pb.Signal, error)

	// Stop stops the sensor.
	Stop(context.Context) error
}
