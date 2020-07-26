package minisafe

import "fmt"

type routine func()

// Wraps the passed function, executes it as a goroutine,
// but defers any recovering from potential
func Go(fn routine) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("The running goroutine panicked with: %s\n", err)
			}
		}()
		fn()
	}()
}
