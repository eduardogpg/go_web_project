package models

//go get golang.org/x/crypto/bcrypt
import (
	"golang.org/x/crypto/bcrypt" 
	//"fmt"
	)

type User struct{
	Id int
	Username string 
	Email string
	EncryptedPassword string
}

func (this *User) SetPassword(text string){
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    this.EncryptedPassword = string(hashedPassword)
}

func (this *User) CheckPassword(text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.EncryptedPassword), []byte(text))
	return err == nil
}

func (this *User) ResetPassword(){
	this.EncryptedPassword = ""
}

func (this *User) Save()bool{
	connection.Create(&this)
	return this.Id > 0
}

func(this *User) Delete(){
  connection.Delete(&this)
}

func (this *User) GetPassword()string{
	return this.EncryptedPassword
}

func FindUser(id int) *User{ //Refactor aqui!
	user := &User{}
  	connection.Where("id = ?", id).First(&user)
  	return user
}

func FindUserBy(field, value string) *User{
	user := &User{}
  	connection.Where(field +" = ?", value).First(&user)
  	return user
}

func NewUser(username, password, email string) (User, ValidateError){
	user := User{ Username:username, Email:email, }
 	err := UserValidate(username, password, email)
 	if err != nil{
 		return user, err
 	}
 	
	user.SetPassword(password)
	return user, nil
}


func CreateUser(username, password, email string) (User, ValidateError){
	user, err := NewUser(username, password, email)
	if err != nil{
		user.Save()
	}
	return user, err
}

