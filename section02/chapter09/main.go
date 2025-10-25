package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("This is the first thing to be printed!")

	// Without this, the program runs too fast for the goroutine to be run
	time.Sleep(1 * time.Second)

	printSomething("This is the second thing to be printed!")
}
