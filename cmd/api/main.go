package main

import (
	"github.com/mfamador/api/v1/internal"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Example gRPC gateway API")

	echo.Run()
}
