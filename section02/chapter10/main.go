package main

import (
	"fmt"
	"sync"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	for i, word := range words {
		wg.Add(1)
		go func() {
			// Use closures, functions should remain as pure as they can get
			printSomething(fmt.Sprintf("%d: %s", i, word))
			wg.Done()
		}()
	}

	wg.Wait()
}
