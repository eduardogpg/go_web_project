package controllers

import(
	"net/http"
	"log"
	"../utils"
	"../models"
)

func NewUser(w http.ResponseWriter, r *http.Request){
	utils.RenderTemplate(w, r, "user/register", nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	
	user, err := models.NewUser(username, password, email)
	if err != nil{
		response := models.Response{ user }
		log.Println("Problemas al crear un nuevo usuario.")
		utils.RenderTemplate(w, r, "user/register", response)
		
		return
	}else{
		utils.RenderTemplate(w, r, "user/register", nil)
	}
}
