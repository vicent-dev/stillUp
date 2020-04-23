package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"stillUp/redis"
	"stillUp/request"
)

func writErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func CallHandler(w http.ResponseWriter, r *http.Request) {
	var c redis.Call
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil || !c.Validate() {
		writErrorResponse(w, errors.New("Error, wrong request"))
		return
	}

	response, err := request.Get(&c)
	if err != nil {
		writErrorResponse(w, err)
		return
	}
	c.Response = response

	go redis.GetCallRepository().Save(&c)

	if err != nil {
		log.Println("Can't store on redis")
	}

	json.NewEncoder(w).Encode(c.Response.Body)
	return
}
