package streamers

import "context"

type Streamer interface {
	PutRecord(ctx context.Context, data []byte) error
	GetRecords(ctx context.Context) ([][]byte, error)
}

// Starter is an interface that defines a Start method
// that should be implemented by types that need to be started.
type Starter interface {
	Start(context.Context) error
}
