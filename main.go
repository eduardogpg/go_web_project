package main

import(
	"net/http"
	"log"
	"fmt"
)

const port = ":3000"

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola Mundo!")
	})	

	log.Println("Server listening in 127.0.0.1"+ port)
	http.ListenAndServe(port, mux)

}