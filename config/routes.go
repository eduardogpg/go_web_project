package config

import(
    "github.com/gorilla/mux"
    "net/http"
    "../controllers"
)

func Routes()(http.Handler){
   mux := mux.NewRouter()
   
   mux.HandleFunc("/", controllers.Index).Methods("GET")
   mux.HandleFunc("/register", controllers.New).Methods("GET")

   return mux
}