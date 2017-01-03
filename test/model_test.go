package test

import (
	"testing"
	"../models"
)

const(
	username string = "eduardo_gpg"
	password string = "eduardo123"
	email string = "eduardo78d@gmail.com"
)

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(username, password, email)
	if err != nil {
		t.Error("No es posible crear al usuario", err)
	}
}

func TestNewUserBadRequest(t *testing.T) {
	_, err := models.NewUser("", "", "")
	if err == nil {
		t.Error("", err)
	}
}

func TestEncryptedPassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if user.EncryptedPassword == password || len(user.EncryptedPassword) != 60{
		t.Error("El password no esta encriptado!", nil)
	}
}

func TestCheckPassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if !user.CheckPassword(password) {
		t.Error("Fail checkPassword", nil)
	}
}

func init(){
}

func main(){
}

