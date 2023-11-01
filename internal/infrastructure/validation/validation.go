package validation

import "github.com/go-playground/validator/v10"

type (
	appValidate struct {
		v *validator.Validate
	}

	AppValidate interface {
		StructExceptRules(s interface{}, rules ...string) error
	}
)

func NewValidator() AppValidate {
	v := validator.New(validator.WithRequiredStructEnabled())

	return &appValidate{
		v,
	}
}

func (ap appValidate) StructExceptRules(s interface{}, rules ...string) error {
	err := ap.v.Struct(s)

	if err == nil {
		return nil
	}

	rulesMap := map[string]bool{}

	for _, rule := range rules {
		rulesMap[rule] = true
	}

	var ve validator.ValidationErrors

	for _, validationErr := range err.(validator.ValidationErrors) {
		if !rulesMap[validationErr.Tag()] {
			ve = append(ve, validationErr)
		}
	}

	return ve
}
