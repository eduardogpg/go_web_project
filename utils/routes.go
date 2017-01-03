package utils

import(
    "github.com/gorilla/mux"
    "net/http"
    "../controllers"
)

func Routes()(http.Handler){
   mux := mux.NewRouter()
   mux.HandleFunc("/", controllers.Index).Methods("GET")

   return mux
}