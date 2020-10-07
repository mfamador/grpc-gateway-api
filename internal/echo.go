package echo

import (
	"context"
	"github.com/mfamador/api/v1/internal/gen"
	"github.com/rs/zerolog/log"
)

type echoService struct{}

func NewEchoService() *echoService {
	return &echoService{}
}

func (s echoService) Echo(ctx context.Context, message *echo.StringMessage) (*echo.StringMessage, error) {
	msg := &echo.StringMessage{Value: "a response"}
	log.Info().Interface("Server Implementation", message)
	return msg, nil
}
