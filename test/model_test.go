package test

import (
	"testing"
	"../models"
	"os"
)

const(
	username string = "eduardo_gpg"
	password string = "eduardo123"
	email string = "eduardo78d@gmail.com"
)

func TestMain(m *testing.M) { 
    setup()
    retCode := m.Run()
    teardown()
    os.Exit(retCode)
}

func setup(){
	models.InitializeDataBase()
}

func teardown(){
	models.DeleteRecords()
	models.CloseConnection()
}

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
	//Tenemos que realizar un refactor! D: que pasa
	user, _ := models.NewUser(username, password, email)
	if user.CheckPassword(password) {
		t.Error("Fail checkPassword", nil)
	}
}

func TestSaveUser(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()
	if user.Id == 0 {
		t.Error("El usuario no se almaceno en la base de datos", nil)
	}
}

func TestFindUser(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()

	user = models.Find(user.Id)
	if user.Id != 1 && user.Username != username {
		t.Error("No se pudo encontrar un registro en la base de datos.", nil)
	}
}

func TestFindByEmail(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()

	user = models.FindBy("email", email)
	if user.Email != email {
		t.Error("No se pudo encontrar un registro en la base de datos.", nil)
	}
}

func TestDeleteUser(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()
	id := user.Id

	user.Delete()

	newUser := models.Find(id)
	if newUser.Id > 0 {
		t.Error("No se pudo eliminar el registro.", nil)
	}
}

func CreateUser(t *testing.T) {
	user, _ := models.CreateUser(username, password, email)
	if user.Id == 0{
		t.Error("No se pudo crear el registro.", nil)
	}
}

func UpdateUser(t *testing.T) {
	user, _ := models.CreateUser(username, password, email)
	var newEmail = "eduardo_gpg@gmail.com"
	user.Email = newEmail
	user.Save()

	user = models.Find(user.Id)
	if user.Email == email{
		t.Error("No se pudo actualizar el registro.", nil)
	}
}

