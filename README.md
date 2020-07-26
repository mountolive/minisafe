# MINISAFE
Safe net for panics inside goroutines
This is very raw port of [cookoo's](https://github.com/Masterminds/cookoo) safely.
Could be seen as mere copy, really... but I'm keeping it here as a personal placeholder.

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
}

func panicker() {
  panic(errors.New("I will panick"))
}
```
