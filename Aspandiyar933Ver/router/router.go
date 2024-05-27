package router

import (
	"database/sql"
	"net/http"

	"github.com/aleksander-git/data-analyzer/resource/book"
	middleware "github.com/aleksander-git/data-analyzer/util/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func New(l *zerolog.Logger, v *validator.Validate, db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.RequestIDMiddleware)
	v1.Use(middleware.ContentTypeJSONMiddleware)

	bookAPI := book.New(l, v, db)
	v1.HandleFunc("/hello", bookAPI.HelloHandler).Methods(http.MethodGet)

	return r
}
