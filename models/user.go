package models

import(
	"errors"
	"golang.org/x/crypto/bcrypt" //go get golang.org/x/crypto/bcrypt
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

func (this *User) SetPassword(text string){
    hashedPassword, err := bcrypt.GenerateFromPassword(toByte(text), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    this.EncryptedPassword = string(hashedPassword)
}

func (this *User) CheckPassword(text string) bool {
	err := bcrypt.CompareHashAndPassword(toByte(this.EncryptedPassword), toByte(text))
	return err != nil
}

func toByte(text string)[]byte{
	return []byte(text)
}

func NewUser(username, email, password string) (User, ValidateError){
	user := User{ 	Username:username,
					Email:email, 
				}

	if username == ""{
		return user, errorUsername
	}

	if email == ""{
		return user, errorEmail
	}

	if password == ""{
		return user, errorPassword
	}

	user.SetPassword(password)
	return user, nil
}
