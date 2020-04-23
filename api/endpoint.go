package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"stillUp/redis"
)

func writErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func CallHandler(w http.ResponseWriter, r *http.Request) {
	var c redis.Call
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		writErrorResponse(w, errors.New("Error decoding request"))
		return
	}
	err = redis.GetCallRepository().Save(&c)
	if err != nil {
		writErrorResponse(w, errors.New("Error saving endpoint on redis"))
		return
	}
	cRedis, err := redis.GetCallRepository().Find(c.Key())

	if err != nil {
		writErrorResponse(w, errors.New("Error getting endpoint from redis"))
		return
	}

	json.NewEncoder(w).Encode(cRedis)
	return
}
