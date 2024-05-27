package controllers

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func ErrorHandler(logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Error().Err(errors.New("this is a test error")).Msg("error occurred")
	}
}

func BadRequestHandler(logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		logger.Warn().Msg("bad request error occurred")
	}
}
