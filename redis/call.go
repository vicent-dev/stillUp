package redis

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
)

type Call struct {
	Url      string            `json:"url"`
	Header   map[string]string `json:"header"`
	Response *Response
}

func (c *Call) Key() string {
	hasher := sha512.New()
	urlHeader := make(map[string]interface{}, 2)
	urlHeader["url"] = c.Url
	urlHeader["header"] = c.Header

	jsonRequest, _ := json.Marshal(urlHeader)
	hasher.Write(jsonRequest)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func (c *Call) Validate() bool {
	return c.Url != ""
}
