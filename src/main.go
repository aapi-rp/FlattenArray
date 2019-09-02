package main

import (
	"api/v1"
	"base"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/flatten_standard", apiv1.Flatten_standard).Methods("Post")
	router.HandleFunc("/v1/flatten_3d", apiv1.Flatten_3d).Methods("Post")
	fmt.Println("Port is set to: ", base.GetPort())

	srv := &http.Server{
		Addr:         ":" + base.GetPort(),
		Handler:      router,
		ReadTimeout:  base.ReadTimeout,
		WriteTimeout: base.WriteTimeout,
		IdleTimeout:  base.IdleTimeout,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	srv.ListenAndServe()

}
