package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	runProcess("P1", 20)
	runProcess("P2", 20)

	go runProcess("P1", 20)  // O fato de colocar o go na frente significa que irão rodar em paralelo
	go runProcess("P2", 20)  // O fato de colocar o go na frente significa que irão rodar em paralelo

	var s string
	// Isso abaixo serve para pegar os resultados em tempo de execução
	fmt.Scanln(&s)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}

// Executar o comando: go run -race main.go // para aguardar a verificação de race Conditions