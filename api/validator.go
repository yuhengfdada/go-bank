package api

import "github.com/go-playground/validator/v10"

func validCurrency(fl validator.FieldLevel) bool {
	currency := fl.Field().Interface().(string)
	for _, c := range []string{"USD", "EUR"} {
		if currency == c {
			return true
		}
	}
	return false
}
