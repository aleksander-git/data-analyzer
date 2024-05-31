package book

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type API struct {
	logger     *zerolog.Logger
	validator  *validator.Validate
	repository *Repository
}

func New(logger *zerolog.Logger, validator *validator.Validate, db *sql.DB) *API {
	return &API{
		logger:     logger,
		validator:  validator,
		repository: NewRepository(db),
	}
}

func (a *API) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		return
	}
}
