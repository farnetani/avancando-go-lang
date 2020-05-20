package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

// Seta a qtde de proxies (cpus) máxima para funcionar de forma paralela
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Forçado para utilizar o numero máximo de cpus
}

func main() {
	fmt.Println(runtime.NumCPU(), " CPUs")
	waitGroup.Add(3) // Starta o grupo de processos
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	go runProcess("P3", 20)
	waitGroup.Wait() // Aguarda o grupo de processos terminar
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		// t := time.Duration(rand.Intn(255))
		t := time.Duration(200) // Seta um numero exato para enxergar melhor que está sendo executado em tempo real

		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done() // Avisa que os processos acabaram
}