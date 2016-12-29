package main

import(
	"net/http"
	"log"
	"./template"
)

const port = ":3000"

func index(w http.ResponseWriter, r *http.Request){
	template.RenderTemplate(w, r, "index/home", nil)
}

func register(w http.ResponseWriter, r *http.Request){
	template.RenderTemplate(w, r, "user/register", nil)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/register", register)

	mux.Handle("/assets/",  http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))), )

	log.Println("Server listening in 127.0.0.1"+ port)
	http.ListenAndServe(port, mux)

}