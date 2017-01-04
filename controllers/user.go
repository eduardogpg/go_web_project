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
	user, err := models.NewUser(r.FormValue("username"), r.FormValue("password"), r.FormValue("email"))
	
	if err != nil{
		log.Println(user)
		user.ResetPassword()
		
		utils.RenderTemplate(w, r, "user/register", map[string] interface{}{
			"Error" : err.Error(),
			"User" : user,
		})
		return
	}
	
	if user.Save(){
		utils.AddCookie()
	}
	
	utils.RenderTemplate(w, r, "index/home", map[string] interface{}{
		"User" : user,
	})
}
