package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var visitNumber uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&visitNumber, 1) //visitNumber++
		w.Write([]byte(fmt.Sprintf("Visitante numero:%d", visitNumber)))
		fmt.Println(visitNumber)
	})

	http.ListenAndServe(":3000", nil)
}
