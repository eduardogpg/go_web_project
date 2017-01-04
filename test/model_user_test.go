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
	NewEmail string = "eduardo_gpg@gmail.com"
)

func TestMain(m *testing.M) { 
    setup()
    retCode := m.Run()
    teardown()
    os.Exit(retCode)
}

func setup(){
	models.InitializeDataBase()
	models.CreateTables()
}

func teardown(){
	models.DropDataTables()
	models.CloseConnection()
}

func createUser(){

}

func TestNewUserWithoutEmail(t *testing.T) {
	_, err := models.NewUser(username, password, "")
	if err.Error() != "El email es requerido." {
		t.Error("Validación incorrecta <Email>.", err)
	}
}

func TestNewUserWithoutPassword(t *testing.T) {
	_, err := models.NewUser(username, "", email)
	if err.Error() != "El password es requerido." {
		t.Error("Validación incorrecta <Password>.", err)
	}
}

func TestNewUserWithoutUsername(t *testing.T) {
	_, err := models.NewUser("", password, email)
	if err.Error() != "El username es requerido." {
		t.Error("Validación incorrecta <Username>.", err)
	}
}

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(username, password, email)
	if err != nil {
		t.Error("No es posible crear un nuevo usuario.", err)
	}
}

func TestValidatePassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if user.GetPassword() == password {
		t.Error("No es posible generar el password del usuario.", nil)
	}
}

func TestCheckPassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if !user.CheckPassword(password) {
		t.Error("No es posible validar el password.", nil)
	}
}

func TestSaveUser(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if !user.Save() { 
		t.Error("No es posible almacenar el registro en la base de datos.", nil)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	if err != nil { 
		t.Error("No es posible crear el registro en la base de datos.", nil)
	}
}

func TestFindUser(t *testing.T) {
	models.CreateUser(username, password, email)
	user := models.FindUser(1)

	if user.Username != username { 
		t.Error("No es posible encontrar el registro.", nil)
	}
}

func TestFindUserByEmail(t *testing.T) {
	models.CreateUser(username, password, email)
	user := models.FindUserBy("email", email)

	if user.Email != email { 
		t.Error("No es posible encontrar el registro por Email.", nil)
	}
}

func TestDeleteUser(t *testing.T){
	user, _ := models.CreateUser(username, password, email)
	user.Delete()
}

func UpdateUsername(t *testing.T){
	models.CreateUser(username, password, email)
	user := models.FindUser(1)
	user.Username = "Eduardo Ismael"
	user.Save()

	if user.Username == username{
		t.Error("No es posible actualizar el registro.", nil)
	}
}



