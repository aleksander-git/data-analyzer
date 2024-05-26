package server

import (
	"net/http"

	"github.com/aleksander-git/data-analyzer/internal/controllers"
	"github.com/aleksander-git/data-analyzer/pkg/logger"
	"github.com/rs/zerolog"
)

func NewRouter(log zerolog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/hello", logger.LogRequest(http.HandlerFunc(controllers.HelloHandler), log))
	mux.Handle("/error", logger.LogRequest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.LogErrorHandler(w, r, log)
	}), log))

	return mux
}
