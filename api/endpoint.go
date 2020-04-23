package api

import (
	"encoding/json"
	"net/http"
)


func CallHandler(w http.ResponseWriter, r *http.Request) {
	var c Call
	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(c)
	return
}
