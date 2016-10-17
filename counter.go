package counter

import (
	"gopkg.in/redis.v4"
	"github.com/najeira/conv"
)

type Counter struct {
	client *redis.Client
	key    string
}

func New(client *redis.Client, key string) *Counter {
	return &Counter{
		client: client,
		key: key,
	}
}

func (s *Counter) Set(value int64) {
	s.client.Set(s.key, value, 0).Result()
}

func (s *Counter) Get() (int64, error) {
	text, error := s.client.Get(s.key).Result()
	return conv.Int(text), error
}

func (s *Counter) Incr() (int64, error) {
	return s.client.Incr(s.key).Result()
}

func (s *Counter) Decr() (int64, error) {
	return s.client.Decr(s.key).Result()
}
