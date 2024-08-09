package setups

import (
	"project-user/commons/helpers"

	"github.com/go-playground/validator/v10"
)

func SetValidator() (validate *validator.Validate) {
	validate = validator.New()
	helpers.UsernameValidator(validate)
	helpers.PasswordValidator(validate)
	helpers.TelephoneValidator(validate)
	return
}
