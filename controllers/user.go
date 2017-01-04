package controllers

import(
	"net/http"
	"log"
	"../utils"
	"../models"
)

func HandlerUserNew(w http.ResponseWriter, r *http.Request){
	utils.RenderTemplate(w, r, "user/register", nil)
}

func HandlerUserCreate(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	
	user, err := models.NewUser(username, password, email)

	if err != nil{
		utils.RenderTemplate(w, r, "user/register", map[string] interface{}{
			"Error" : err.Error(),
			"User" : user,
		})
		return
	}
	
	log.Println("Tambien se va a ejecutar esto D: ")
	utils.RenderTemplate(w, r, "user/register", nil)
}
