package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

func (worker *Worker) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if worker.Email == "" {
			err = errors.New("Email requerido")
			errorMessages["Required_email"] = err.Error()
		}
		if worker.Email != "" {
			if err = checkmail.ValidateFormat(worker.Email); err != nil {
				err = errors.New("Email invalido")
				errorMessages["Email_invalido"] = err.Error()
			}
		}

	case "login":
		if worker.Password == "" {
			err = errors.New("Contraseña requerida")
			errorMessages["Required_password"] = err.Error()
		}
		if worker.Email == "" {
			err = errors.New("Email requerido")
			errorMessages["Required_email"] = err.Error()
		}
		if worker.Email != "" {
			if err = checkmail.ValidateFormat(worker.Email); err != nil {
				err = errors.New("invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if worker.Name == "" {
			err = errors.New("Username requerido")
			errorMessages["Required_username"] = err.Error()
		}
		if worker.Password == "" {
			err = errors.New("Contraseña requerida")
			errorMessages["Required_password"] = err.Error()
		}
		if worker.Password != "" && len(worker.Password) < 6 {
			err = errors.New("La contraseña tiene que tener minimo 6 letras")
			errorMessages["Invalid_password"] = err.Error()
		}
		if worker.Email == "" {
			err = errors.New("Email requerido")
			errorMessages["Required_email"] = err.Error()

		}
		if worker.Email != "" {
			if err = checkmail.ValidateFormat(worker.Email); err != nil {
				err = errors.New("Email invalido")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}
	return errorMessages
}
