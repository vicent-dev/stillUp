package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stillUp/redis"
)

func Get(c *redis.Call) (*redis.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range c.Header {
		req.Header.Add(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := &redis.Response{}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &r.Body)
	return r, err
}
