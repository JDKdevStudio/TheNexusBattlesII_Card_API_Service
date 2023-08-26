package models

import "github.com/go-playground/validator/v10"

type QueryPagination struct {
	Page_size   int `query:"page_size" validate:"required"`
	Page_number int `query:"page_number" validate:"required"`
}

func (e QueryPagination) Validate() bool {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return false
	}
	return true
}
