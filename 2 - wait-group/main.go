package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
		Utiliza-se de wait-group para garantir a sincronia de goroutines

	*/

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // quantas goroutines serão gerenciadas

	go func() {
		escrever1("Capivara na função 1")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever2("Capivara na função 2")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // garante q as duas funçoes serão finalizadas

}

func escrever1(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func escrever2(texto string) {
	for i := 0; i < 7; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
