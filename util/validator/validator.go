package validator

import (
	"reflect"

	"github.com/aleksander-git/data-analyzer/util/validator"
	"github.com/go-playground/validator/v10"
)

const (
	alphaSpaceRegexString string = "^[a-zA-Z ]+$"
)

type ErrResponce struct {
	Errors []string `json:"errors"`
}

func New() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("form")

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("alpha_space", isAlphaSpace)

	return validate
}