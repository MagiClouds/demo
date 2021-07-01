package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

type (
	//Welcome
	Welcome struct{}
)

var _ http.Handler = (*Welcome)(nil)

func (wc *Welcome) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	c := make(chan struct{}, 10)
	n := 0
	for {
		select {
		case i := <-c:
			fmt.Println("\n------------: ", i)
			w.Write([]byte("hello world"))
			return
		default:
			fmt.Println("hhahahah")
			if n == 1 {
				c <- struct{}{}
			}
			n++
		}
	}
}

func main() {

	wc := &Welcome{}

	http.Handle("/", wc)

	http.ListenAndServe(":8080", nil)
}
