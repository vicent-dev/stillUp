package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stillUp/redis"
	"sync"
)

func Get(c *redis.Call, wg *sync.WaitGroup) (*redis.Response, error) {
	defer wg.Done()
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
	r.Header = make(map[string]string, len(response.Header))

	for k, v := range response.Header {
		r.Header[k] = v[0]
	}

	return r, err
}
