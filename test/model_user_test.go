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
}

func teardown(){
	models.DeleteRecords()//Debemos de eliminar la tabla, no eliminar los registros en la base de datos!
	models.CloseConnection()
}

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(username, password, email)
	if err != nil {
		t.Error("No es posible crear el usuario.", err)
	}
}

func TestNewUserBadRequest(t *testing.T) {
	_, err := models.NewUser("", "", "")
	if err == nil {
		t.Error("El usuario si ha creado de forma correcta.", err)
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
		t.Error("El password es el mismo.", nil)
	}
}

func TestSaveUser(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()
	if user.Id == 0 {
		t.Error("No es posible crear el registro.", nil)
	}
}

func TestFindUser(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()

	user = models.Find(user.Id)
	if user.Id != 1 && user.Username != username {
		t.Error("No es posible encontrar el registro.", nil)
	}
}

func TestFindByEmail(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()

	user = models.FindBy("email", email)
	if user.Email != email {
		t.Error("No es posible encontrar el registro.", nil)
	}
}

func TestDeleteUser(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	user.Save()
	id := user.Id

	user.Delete()

	newUser := models.Find(id)
	if newUser.Id > 0 {
		t.Error("No es posible eliminar el registro.", nil)
	}
}

func CreateUser(t *testing.T) {
	user, _ := models.CreateUser(username, password, email)
	if user.Id == 0{
		t.Error("No es posible crear el registro.", nil)
	}
}

func UpdateUser(t *testing.T) {
	user, _ := models.CreateUser(username, password, email)
	user.Email = NewEmail
	user.Save()

	user = models.Find(user.Id)
	if user.Email == email{
		t.Error("No es posible actualizar el registro.", nil)
	}
}

