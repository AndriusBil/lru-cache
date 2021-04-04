# lru-cache

## Built with
Project built with go1.16, using only go builtin.

## Getting started

### Prerequisites

Install [golang](https://golang.org/doc/install)

## Usage

1. Create new cache object
```go
c := lru_cache.New(1000)
```

2. Put new value to cache
```go
c.Put("any unique key", "any value")
```

3. Get value from cache
```go
value := c.Get("any unique key")
```