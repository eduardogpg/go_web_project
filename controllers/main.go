package controllers

import(
	"net/http"
	"../utils"
	"../models"
)

func Index(w http.ResponseWriter, r *http.Request){
	utils.RenderTemplate(w, r, "index/home", nil)
}

func Login(w http.ResponseWriter, r *http.Request){
	utils.RenderTemplate(w, r, "index/login", nil)
}

func HandleLogin(w http.ResponseWriter, r *http.Request){
	//user := User{}
	user := models.FindBy("email", r.FormValue("username"))
	if user.CheckPassword(r.FormValue("password")){
		utils.RenderTemplate(w, r, "index/home", nil)
		return
	}
	
	utils.RenderTemplate(w, r, "index/login", map[string] interface{}{
		"User" : user, 
	})
}

