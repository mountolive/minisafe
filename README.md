# MINISAFE
Safe net for panics inside goroutines.
This is a very raw port of [cookoo's](https://github.com/Masterminds/cookoo) safely.

# Installation
`go get github.com/mountolive/minisafe`

# Usage

```
package main

import (
    "errors"
    "fmt"

    "github.com/mountolive/minisafe"
)
func main() {
    // no params
    minisafe.Go(routine)
    // with params
    msg := "I will be passed on"
    minisafe.Go(func() {
        routine(msg)
    })
}

func routine(args...interface{}) {
  fmt.Println("I'll run as a goroutine")
  panicker()
}

func panicker() {
  panic(errors.New("I will panick"))
}
```
