package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start(port string) {
	r := mux.NewRouter()
	r.Use(jsonMiddleware)

	r.HandleFunc("/", CallHandler).Methods("POST").Schemes("http")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
