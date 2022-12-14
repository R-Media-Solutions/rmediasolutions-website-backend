package main

import (
	"log"
	"net/http"

	"github.com/R-Media-Solutions/rmediasolutions-website-backend/controllers/aboutuscontroller"
	"github.com/R-Media-Solutions/rmediasolutions-website-backend/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/aboutus", aboutuscontroller.Index).Methods("GET")
	//r.HandleFunc("/aboutus/{id}", aboutuscontroller.Show).Methods("GET")
	//r.HandleFunc("/aboutus", aboutuscontroller.Create).Methods("POST")
	//r.HandleFunc("/aboutus/{id}", aboutuscontroller.Update).Methods("PUT")
	//r.HandleFunc("/aboutus", aboutuscontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
