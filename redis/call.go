package redis

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

type Call struct {
	Url    string            `json:"url"`
	Header map[string]string `json:"header"`
}

func (c *Call) Key() string {
	hasher := sha512.New()
	hasher.Write([]byte(fmt.Sprintf("%v", c)))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
