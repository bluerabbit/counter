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

var counter = counter.NewCounter(client, "comment_count")

value, ok := counter.Set(1)
counter.Incr() // 2
counter.Incr() // 3
counter.Decr() // 2
counter.Get()  // 2
```

## License

MIT
