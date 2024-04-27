package main

import "fmt"

func main() {

	/*
		Enquanto um canal padrão é sincronizado e bloqueia a goroutine remetente até que haja um receptor pronto
		para receber o valor, um canal com buffer permite que a goroutine remetente continue
		sua execução após colocar um valor no canal, desde que o buffer não esteja cheio.

	*/

	canal := make(chan string, 2)
	canal <- "capivara"

	mensagem := <-canal
	fmt.Println(mensagem)

}
