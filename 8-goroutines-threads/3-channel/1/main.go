package main

import (
	"fmt"
	"time"
)

// T1
func main() {
	canal := make(chan string)

	// T2
	go func() {
		canal <- "Minha primeira mensagem do chanel"
		time.Sleep(time.Second * 60)
	}()

	msg := <-canal // vai fazer a nosas T1 aguardar
	fmt.Println(msg)

}
