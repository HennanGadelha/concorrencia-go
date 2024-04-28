package main

import (
	"fmt"
	"time"
)

func main() {

	// Concorrencia duas tarefas revezando o tempo de execuçao != pararelismo -> duas ou mais tarefas executadas ao mesmo tempo
	/*
		utilizando "go" antes da função escrever1() signgificará que a função será executada e o programa seguirá executando
		sem que haja a necessidade da funçao escrever1() ter finalizado, ou seja, as funçoes escrever 1 e 2 estarão rodando ao mesmo tempo
	*/
	go escrever1("Capivara na função 1")
	escrever2("Capivara na função 2")

}

func escrever1(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func escrever2(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
