package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	msg = "Hello, world!"

	messages := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	for _, message := range messages {
		wg.Add(1)
		go func() {
			updateMessage(message)
			wg.Done()
		}()
		wg.Wait()
		printMessage()
	}
}
