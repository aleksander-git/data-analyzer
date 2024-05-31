package main

import (
	"net/http"
	"os"

	server "github.com/aleksander-git/data-analyzer/internal/services"
	"github.com/rs/zerolog"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	zeroLogger := zerolog.New(output).With().Timestamp().Logger()

	zeroLogger.Info().Msg("Logger initialized")
	zeroLogger.Info().Msg("Starting server")

	router := server.NewRouter(zeroLogger)
	if err := http.ListenAndServe(":8080", router); err != nil {
		zeroLogger.Error().Err(err).Msg("could not start server")
	}
}
