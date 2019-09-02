package main

import (
	"api/v1"
	"base"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	// I created an api out of the system to launch the flattening process
	router := mux.NewRouter()
	router.HandleFunc("/flatten_standard", apiv1.Flatten_standard).Methods("Post")
	router.HandleFunc("/flatten_standard", apiv1.Flatten_3d).Methods("Post")

	http.ListenAndServe(":"+base.GetPort(), nil)

}
