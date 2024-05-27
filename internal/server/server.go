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
	mux.Handle("/error", logger.LogRequest(controllers.ErrorHandler(log), log))
	mux.Handle("/badrequest", logger.LogRequest(controllers.BadRequestHandler(log), log))

	return mux
}
