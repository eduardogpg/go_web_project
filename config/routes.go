package config

import(
    "github.com/gorilla/mux"
    "net/http"
    "../controllers"
)

func Routes()(http.Handler){
   mux := mux.NewRouter()
   
   mux.HandleFunc("/", controllers.Index).Methods("GET")
   
   mux.HandleFunc("/register", controllers.HandlerUserNew).Methods("GET")
   mux.HandleFunc("/register", controllers.HandlerUserCreate).Methods("POST")

   return mux
}