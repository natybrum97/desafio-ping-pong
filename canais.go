//canal: modo de duas goroutines se comunicarem e sincronizarem sua execução.
//time imprime uma mensagem a cada intervalo de tempo.

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) { // palavra reservada para Canal: chan
	for i := 0; ; i++ {
		c <- "ping" //usado para enviar e receber mensagem pelo canal
		//Neste caso, está enviando a string "ping" para o canal c.

	}
}

func pongar(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

	}
}
func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)
	go pongar(c)

	var entrada string
	fmt.Scanln(&entrada)
}
