package aggregator

import (
	"context"
	"errors"
	"math"
)

var (
	ErrInvalidMaxSize = errors.New("invalid max size")
)

// Aggregate aggregates the data from the source channel into chunks
// of at most maxSize bytes.
// The function returns a channel of chunks. Sending to the output
// channel when the chunk size reaches maxSize.
// Note: This implementation is useful for small data sizes, for
// large data sizes consider using another implementation like [io.Reader]
func Aggregate(
	ctx context.Context,
	source chan []byte,
	maxSize int,
) (chan [][]byte, error) {
	if maxSize <= 0 {
		return nil, ErrInvalidMaxSize
	}

	out := make(chan [][]byte)

	go func() {
		// estimatedSize is used to pre-allocate the chunk slice
		// is calculated based on the number of elements in the last chunk
		estimatedSize := 1

		// leftData is used to store the data that didn't fit in the last chunk
		var leftData []byte
	nextChunk:
		totalSize := 0
		chunk := make([][]byte, 0, estimatedSize)
		if leftData != nil {
			chunk = append(chunk, leftData)
			totalSize += len(leftData)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-source:
				if !ok {
					return
				}

				if totalSize+len(d) > maxSize {
					leftData = d
					estimatedSize = CalculateEstimatedSize(len(chunk))
					out <- chunk
					goto nextChunk
				}

				leftData = nil
				chunk = append(chunk, d)
			}
		}
	}()

	return out, nil
}

// CalculateEstimatedSize calculates the estimated size of the chunk slice
func CalculateEstimatedSize(
	l int,
) int {
	tenPer := float64(l) * 0.1

	return int(math.Floor(float64(l) - tenPer))
}
