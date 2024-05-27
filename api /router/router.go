package router

import (
	"database/sql"

	book "command-line-arguments/Users/user/Desktop/data-analyzer/api /resource/book/handler.go"

	"github.com/aleksander-git/data-analyzer/api /resource/book"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func New(l *zerolog.Logger, v *validator.Validate, db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(requestIDMiddleware)
	v1.Use(contentTypeJSONMiddleware)

	bookAPI := book.New(l, w, db)
}