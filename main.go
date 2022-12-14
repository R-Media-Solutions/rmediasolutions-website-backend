package main

import (
	"log"
	"net/http"

	"github.com/R-Media-Solutions/rmediasolutions-website-backend/controllers/aboutuscontroller"
	"github.com/R-Media-Solutions/rmediasolutions-website-backend/controllers/productcontroller"
	"github.com/R-Media-Solutions/rmediasolutions-website-backend/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/aboutus", aboutuscontroller.Index).Methods("GET")
	r.HandleFunc("/aboutus/{id}", aboutuscontroller.Show).Methods("GET")
	r.HandleFunc("/aboutus/create", aboutuscontroller.Create).Methods("POST")
	r.HandleFunc("/aboutus/update/{id}", aboutuscontroller.Update).Methods("PUT")
	r.HandleFunc("/aboutus/delete", aboutuscontroller.Delete).Methods("DELETE")

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/products/create", productcontroller.Create).Methods("POST")
	r.HandleFunc("/products/update/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/products/delete", productcontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
