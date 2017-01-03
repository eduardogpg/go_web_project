package models

import(
	"errors"
	//"golang.org/x/crypto/bcrypt"
)

type User struct{
	Username string 
	Email string
	EncryptedPassword string
}

type ValidateError error

var(
	errorUsername = ValidateError(errors.New("You Must supply a username"))
	errorEmail = ValidateError(errors.New("You must supply a email"))
	errorPassword = ValidateError(errors.New("You must supply a password"))
)

func (this *User) setPassword() {
    this.EncryptedPassword = "Eduardo"
}

func NewUser(username, email, password string) (User, ValidateError){
	user := User{ username, email, password }
	if username == ""{
		return user, errorUsername
	}

	if email == ""{
		return user, errorEmail
	}

	if password == ""{
		return user, errorPassword
	}
	user.setPassword()
	return user, nil
}
