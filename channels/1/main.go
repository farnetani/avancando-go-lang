package main

import (
	"fmt"
)

func main() {

	msg := make(chan string)

	go func() {
		msg <- "Hello World"  // Manda para o meu channel = msg esse valor = Hello World
	}()

	result := <-msg  // esvazia o canal e joga pra result
	fmt.Println(result)

}
