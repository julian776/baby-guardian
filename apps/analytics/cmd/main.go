package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog"
)

func main() {
	rootCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Msg("analyzing baby data")

	<-rootCtx.Done()
}
