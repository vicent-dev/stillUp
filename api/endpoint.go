package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"stillUp/redis"
	"stillUp/request"
	"sync"
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
	rHttp, rRedis, err := getResponse(err, c)

	if err != nil {
		c.Response = rRedis
	} else {
		c.Response = rHttp
		go redis.GetCallRepository().Save(&c)
	}

	//for k, v := range c.Response.Header {
	//	w.Header().Add(k, v)
	//}

	json.NewEncoder(w).Encode(c.Response.Body)
	return
}

func getResponse(err error, c redis.Call) (*redis.Response, *redis.Response, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var rHttp, rRedis *redis.Response
	go func(c *redis.Call, wg *sync.WaitGroup) {
		rHttp, err = request.Get(c, wg)
	}(&c, &wg)

	go func(key string, wg *sync.WaitGroup) {
		rRedis, _ = redis.GetCallRepository().Find(key, wg)
	}(c.Key(), &wg)
	wg.Wait()
	return rHttp, rRedis, err
}
