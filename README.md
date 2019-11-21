# Go Cache
A simple, fast and easy to use cache storage solution for caching of data on the same machine. Cache is a small package that you can use in your Go application to store data for fast access and provides convenient methods to do so.

## Install
```commandline
go get github.com/mrbenosborne/go-cache
```

### Creating a Cache

```go
package main

import "github.com/mrbenosborne/go-cache"

func main() {

    // Create a new cache with a default
    // size of 0.
    _ = cache.New(0)

}
```