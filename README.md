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

    // Set the value of the key "userId" to 42
    cache.Set("userId", 42)

    // Get the value of the key "userId"
    if userId, ok := cache.Get("userId"); ok {
        fmt.Println(userId)
    }

    // Delete the value of the key "userId"
    cache.Delete("userId")
}
```
