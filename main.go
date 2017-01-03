package main

import(
	"net/http"
	"log"
	"fmt"
	"./config"
)

const port = ":3000"

func init(){

}

func main() {
	mux := config.Routes()
	fmt.Println("Server ready.")
	log.Fatal(http.ListenAndServe(port, mux))
}