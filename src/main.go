package main

import (
	"api/v1"
	"base"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/flatten_standard", apiv1.Flatten_standard).Methods("Post")
	router.HandleFunc("/v1/flatten_3d", apiv1.Flatten_3d).Methods("Post")

	http.ListenAndServe(":"+base.GetPort(), nil)

}
