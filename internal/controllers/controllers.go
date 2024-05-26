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

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func LogErrorHandler(w http.ResponseWriter, r *http.Request, logger zerolog.Logger) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	logger.Error().Err(errors.New("this is a test error")).Msg("error occurred")
}
