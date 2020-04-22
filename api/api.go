package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/", CallHandler).Methods("GET").Schemes("http")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
