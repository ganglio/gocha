# GoCha

[![Build Status](https://travis-ci.org/ganglio/gocha.svg?branch=master)](https://travis-ci.org/ganglio/gocha)
[![codecov](https://codecov.io/gh/ganglio/gocha/branch/master/graph/badge.svg)](https://codecov.io/gh/ganglio/gocha)
[![GoDoc](https://godoc.org/github.com/ganglio/gocha?status.svg)](https://godoc.org/github.com/ganglio/gocha)

Funny tools to play with channels.

## Status

Currently under active development.

## Components

  - `Mux`, A channel multiplexer. Gets a number of channels as input and returns a single channel as output. Each message sent on any of the input channel is forwarded to the output one.
  - `Pipe`, Takes a number of functions as input and exposes an input and output channels. Upon receiving a message on the input channel it applies all the input functions in order (the order is important) and sends the result on the output channel.
  - `Tee`, Takes a number of channels as output and exposes an input channel. Upon receiving a message on the input channel it forwards it to all the output channels.

## Usage

#### Import package
```go
import (
  "github.com/ganglio/gocha"
)
```

### `Mux`


#### Initialize the muxer
```go
m := gocha.NewMux()
```

#### Add few channels
```go
c1 := make(chan interface{})
c2 := make(chan interface{})
c3 := make(chan interface{})

m.AddChannels(c1,c2,c3)
```

#### Profit
```go
c1<-1
fmt.Println(<-m.Out())
```

#### Closing

When all the channels muxed into the Mux are closed the Mux is closed as well.

### `Pipe`

#### Initialize the pipe
```go
p := gocha.NewPipe()
```

#### Add few functions
```go
f1 := func(i interface{}) interface{} { return 3 * i.(int) + 1}
f2 := func(i interface{}) interface{} { return i.(int) / 2}
f3 := func(i interface{}) interface{} { return i.(int) * i.(int)}

p.AddProcs(f1,f2,f3)
```

#### Profit
```go
p.In()<-1
fmt.Println(<-p.Out())
```

#### Closing

When the input channel is closed the out is closed automatically

### `Tee`

#### Initialize the tee
```go
t := gocha.NewTee()
```

#### Add few outs
```go
c1 := make(chan interface{})
c2 := make(chan interface{})
c3 := make(chan interface{})

t.AddOuts(c1,c2,c3)
```

#### Profit
```go
t.In()<-1
fmt.Println(<-c1)
fmt.Println(<-c2)
fmt.Println(<-c3)
```

#### Closing

When the input channel is closed all the outs are closed