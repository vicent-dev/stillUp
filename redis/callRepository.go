package redis

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"os"
	"sync"
)

var callRepository *CallRepository

type CallRepository struct {
	client *redis.Client
}

func GetCallRepository() *CallRepository {
	if callRepository == nil {
		redisHost := os.Getenv("REDIS_HOST")
		redisPort := os.Getenv("REDIS_PORT")
		redisPassword := os.Getenv("REDIS_PASSWORD")
		callRepository, _ = Connect(redisHost, redisPort, redisPassword)
	}
	return callRepository
}

func Connect(host string, port string, pwd string) (*CallRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pwd,
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &CallRepository{client: client}, nil
}

func (c *CallRepository) Find(key string, wg *sync.WaitGroup) (*Response, error) {
	defer wg.Done()
	val, err := c.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	call := &Call{}
	err = json.Unmarshal([]byte(val), call)

	return call.Response, err
}

func (c *CallRepository) Save(call *Call) error {
	cJson, err := json.Marshal(call)
	if err != nil {
		return err
	}
	return c.client.Set(call.Key(), cJson, 0).Err()
}
