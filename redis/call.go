package redis

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

type Call struct {
	Header map[string]string `json:"header"`
	Body   interface{} `json:"body"`
	Method string `json:"method"`
}

func (c *Call) Key() string {
	hasher := sha512.New()
	hasher.Write([]byte(fmt.Sprintf("%v", c)))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
