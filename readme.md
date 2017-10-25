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

## Usage

#### 1. Import package
```go
import (
  "github.com/ganglio/gocha"
)
```

### `Mux`

#### 2. Initialize the muxer
```go
m := gocha.NewMux()
```

#### 3. Add few channels

```go
c1 := make(chan interface{})
c2 := make(chan interface{})
c3 := make(chan interface{})

m.AddChannels(c1,c2,c3)
```

#### 4. Profit
```go
c1<-1
fmt.Println(<-m.Out())
```

#### Closing

When all the channels muxed into the Mux are closed the Mux is closed as well.

### `Pipe`

#### 2. Initialize the pipe
```go
p := gocha.NewPipe()
```

#### 3. Add few functions

```go
f1 := func(i interface{}) interface{} { return 3 * i.(int) + 1}
f2 := func(i interface{}) interface{} { return i.(int) / 2}
f3 := func(i interface{}) interface{} { return i.(int) * i.(int)}

p.AddProcs(f1,f2,f3)
```

#### 4. Profit
```go
p.In()<-1
fmt.Println(<-p.Out())
```