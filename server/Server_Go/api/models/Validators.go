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
			err = errors.New("required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if worker.Email != "" {
			if err = checkmail.ValidateFormat(worker.Email); err != nil {
				err = errors.New("invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}

	case "login":
		if worker.Password == "" {
			err = errors.New("required password")
			errorMessages["Required_password"] = err.Error()
		}
		if worker.Email == "" {
			err = errors.New("required email")
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
			err = errors.New("required username")
			errorMessages["Required_username"] = err.Error()
		}
		if worker.Password == "" {
			err = errors.New("required password")
			errorMessages["Required_password"] = err.Error()
		}
		if worker.Password != "" && len(worker.Password) < 6 {
			err = errors.New("password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
		}
		if worker.Email == "" {
			err = errors.New("required email")
			errorMessages["Required_email"] = err.Error()

		}
		if worker.Email != "" {
			if err = checkmail.ValidateFormat(worker.Email); err != nil {
				err = errors.New("invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}
	return errorMessages
}
