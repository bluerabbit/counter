# store

## Import

```go
import "github.com/bluerabbit/counter"
```

## Usage

```go
var client *redis.Client

client = redis.NewClient(&redis.Options{
  Addr:       "localhost:6379",
  Password:   "",
  PoolSize:    10,
  PoolTimeout: 10,
})

var c = counter.New(client, "comment_count")

value, ok := c.Set(1)
c.Incr() // 2
c.Incr() // 3
c.Decr() // 2
c.Get()  // 2
```

## License

MIT
