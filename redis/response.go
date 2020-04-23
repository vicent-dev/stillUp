package redis

type Response struct {
	Body   map[string]interface{}
	Header map[string][]string `json:"header"`
	Code   int
}
