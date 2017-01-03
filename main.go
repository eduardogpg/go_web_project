package main

import(
	"net/http"
	"log"
	"fmt"
	"./utils"
)

const port = ":3000"

func init(){

}

func main() {
	mux := utils.Routes()
	fmt.Println("Server ready.")
	log.Fatal(http.ListenAndServe(port, mux))
}