# Golang thread-safe in-memory cache

In-memory cache allows you to cache any data thread-safely

## Installation

``` sh
go get github.com/gibiw/cache
```

## Usage

``` go
import (
    "fmt"
    "github.com/gibiw/cache"
)

func main() {
    // Create a cache
    cache := cache.New()

    // Set the value of the key "userId" to 42 and TTL 10 seconds
    cache.Set("userId", 42, time.Second * 10)

    // Get the value of the key "userId"
    userId, err := cache.Get("userId")
    if err != nil {
        fmt.Println(err)
    }   
    fmt.Println(userId)

    // Delete the value of the key "userId"
    cache.Delete("userId")
}
```
