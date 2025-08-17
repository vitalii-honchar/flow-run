package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

func Struct[T any](s T) (T, error) {
	if err := validate.Struct(s); err != nil {
		return s, err
	}
	return s, nil
}
