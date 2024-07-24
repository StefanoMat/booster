package main

import (
	"fmt"
	"sync"
	"time"
)

func numeroAoQuadrado(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)
	time.Sleep(1 * time.Second)
	wg.Done()
}
func main() {
	wg := sync.WaitGroup{}
	start := time.Now()
	//bloco de codigo
	for x := 0; x < 100; x++ {
		wg.Add(1)
		go numeroAoQuadrado(x, &wg)
	}
	wg.Wait() //aguardando até todas as go routines finalizarem
	end := time.Now()
	fmt.Printf("duration: %s\n", end.Sub(start))
}
