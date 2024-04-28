package main

import "fmt"

func main() {

	canal := escrever("capivara")

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(texto string) <-chan string {

	canal := make(chan string)
	go func() {

		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
		}

	}()

	return canal
}
