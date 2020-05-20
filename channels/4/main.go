package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	// channel <- 10 - isso sem estar dentro de uma função causa Deadlock!
	go func() {
		channel <- 10 // é obrigado a estar dentro de uma go função...
	}()
	fmt.Println(<-channel)
}
