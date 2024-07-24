package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02T15:04:05.999999999Z07:00"
	str := "2024-07-24T10:00:00.000Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		month := t.Month()

		// Imprime o mês
		fmt.Println("Mês:", month)
		fmt.Println(t)

	}
}
