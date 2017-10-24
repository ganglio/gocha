# GoCha

[![Build Status](https://travis-ci.org/ganglio/gocha.svg?branch=master)](https://travis-ci.org/ganglio/gocha)
[![codecov](https://codecov.io/gh/ganglio/gocha/branch/master/graph/badge.svg)](https://codecov.io/gh/ganglio/gocha)
[![GoDoc](https://godoc.org/github.com/ganglio/gocha?status.svg)](https://godoc.org/github.com/ganglio/gocha)

Funny tools to play with channels.

## Status

Currently under active development.

## Usage

### 1. Import package
```go
import (
  "github.com/ganglio/gocha"
)
```

### 2. Initialize the muxer
```go
m, ch := gocha.NewMux()
```

### 3. Add few channels

Either with:

```go
c1 := make(chan interface{})
c2 := make(chan interface{})
c3 := make(chan interface{})

m.AddChannel(c1)
m.AddChannel(c2)
m.AddChannel(c3)
```

Or with:
```go
c1 := make(chan interface{})
c2 := make(chan interface{})
c3 := make(chan interface{})

m.AddChannels(c1,c2,c3)
```

### 4. Profit
```go
c1<-1
fmt.Println(<-m)
```

## Closing

When all the channels muxed into the Mux are closed the Mux is closed as well.
*TODO: Make this optional*