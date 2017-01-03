package main

import(
	"net/http"
	"log"
	"fmt"
	"./config"
	"./models"
)

const port = ":3000"

func init(){

}

func main() {
	models.InitializeDataBase()
	defer models.CloseConnection()

	mux := config.Routes()
	
	fmt.Println("Server ready.")
	log.Fatal(http.ListenAndServe(port, mux))
}