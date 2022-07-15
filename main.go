package main

import (
	"accountanApp/controllers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//router.HandleFunc("/endpoint", controllers.Function).Methods("GET/POST/PUT/DELETE)

	router.HandleFunc("/users", controllers.GetUsersData).Methods("GET")
	router.HandleFunc("/savings", controllers.InsertNewSaving).Methods("POST")
	log.Println("Starting web on port")
	err := http.ListenAndServe(":8000", router)

	log.Fatal(err)
}
