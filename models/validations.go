package models

import "errors"

type ValidateError error

var(
	errorUsername = ValidateError(errors.New("El username es requerido."))
	errorEmail = ValidateError(errors.New("El email es requerido."))
	errorPassword = ValidateError(errors.New("El password es requerido."))
	errorLenPassword = ValidateError(errors.New("Error en la logitud del password."))
)

func UserValidate(username, email, password string) ValidateError {
	if username == ""{
		return errorUsername
	}
	
	if email == ""{
		return errorEmail
	}

	if password == ""{
		return errorPassword
	}

	lenPassword := len(password)
	if lenPassword > 50 || lenPassword < 3 {
		return errorLenPassword
	}
	return nil
}