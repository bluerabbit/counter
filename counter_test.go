package counter

import (
	"gopkg.in/redis.v4"
	"testing"
)

func Test(t *testing.T) {
	var client *redis.Client

	client = redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "",
		PoolSize:    10,
		PoolTimeout: 10,
	})

	c := NewCounter(client, "key")
	c.Set(1)

	v1, error := c.Get()
	if error != nil {
		t.Errorf("error")
	}
	if v1 != 1 {
		t.Errorf("got %v", v1)
	}

	v2, error := c.Incr()
	if error != nil {
		t.Errorf("error")
	}

	if v2 != 2 {
		t.Errorf("got %v", v2)
	}

	v3, error := c.Decr()
	if error != nil {
		t.Errorf("error")
	}

	if v3 != 1 {
		t.Errorf("got %v", v3)
	}
}
