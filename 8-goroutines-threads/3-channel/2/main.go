package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) //T1
	go publicacao(ch)    //T2
	leitura(ch)          //T1

}

func leitura(ch chan int) {
	//printar
	for x := range ch {
		fmt.Printf("Publicacao recebida: %d\n", x)
	}
}
func publicacao(ch chan int) {
	//publicar
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		ch <- i

	}
	close(ch)
}

//Declarar 2 funcs: leitura e outra de publicacao.
//Publicar n dados na thread de publicacao e a funcao de leitura
//devera executar o print sob demanda.
