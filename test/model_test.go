package model

import (
	"testing"
	"./models/user"
)
const(
	username string = "eduardo_gpg"
	password string = "eduardo123"
	email string = "eduardo78d@gmail.com"
)

func TestNewUser(t *testing.T) {
	user, err := user.NewUser(username, password, email)
	if err != {
		t.Error("No es posible crear al usuario", err)
	}
}


func init(){
	
}

func main(){

}

