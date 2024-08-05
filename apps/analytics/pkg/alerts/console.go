package alerts

import "github.com/rs/zerolog"

type Console struct {
	logger *zerolog.Logger
}

func NewConsole(logger *zerolog.Logger) *Console {
	return &Console{
		logger: logger,
	}
}

func (c *Console) Alert(message string) error {
	c.logger.Info().Msg(message)
	return nil
}
