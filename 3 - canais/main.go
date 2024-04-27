package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Um canal em Go é uma estrutura de dados que permite a comunicação e sincronização entre goroutines.
		É utilizado para a troca de dados entre goroutines de forma segura e coordenada
	*/

	canal := make(chan string)
	go escrever("Capivara", canal)

	/*
		A função abaixo irá forçar um erro de deadlock. O deadlock acontece quando nao existe nenhum lugar enviando dados para o canal,
		mas o canal continuar esperando um dado chegar. O deadlock não é pego em momento de compilação, mas sim em execução
	*/

	// exemploErroDeadlock(canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim da aplicaçao")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Envio de dado para o canal
		time.Sleep(time.Second)
	}

	close(canal)
}

func exemploErroDeadlock(canal chan string) {

	for {
		mensagem := <-canal
		fmt.Println(mensagem)
	}

}
