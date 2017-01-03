package controllers

import(
	"net/http"
	"fmt"
	"../utils"
)

func Index(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, "Hola Mundo!")
}

func New(w http.ResponseWriter, r *http.Request){
	utils.RenderTemplate(w, r, "user/register", nil)
}
