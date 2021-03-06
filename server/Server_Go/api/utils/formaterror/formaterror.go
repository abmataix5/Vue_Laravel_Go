package formaterror

import (
	"strings"
)

var errorMessages = make(map[string]string)
var err error

func FormatError(errString string) map[string]string {
	if strings.Contains(errString, "username") {
		errorMessages["Taken_username"] = "El nombre de usuario ya existe"
	}

	if strings.Contains(errString, "email") {
		errorMessages["Taken_email"] = "El email ya esta en uso"

	}
	if strings.Contains(errString, "hashedPassword") {
		errorMessages["Incorrect_password"] = "Incorrect Password"
	}
	if strings.Contains(errString, "record not found") {
		errorMessages["No_record"] = "No Record Found"
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	if len(errorMessages) == 0 {
		errorMessages["Incorrect_details"] = "Incorrect Details"
		return errorMessages
	}

	return nil
}
