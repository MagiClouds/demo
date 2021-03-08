package start_micro

import (
	"io"
	"log"
	"net/http"
	//"net/url"
	"testing"
)

type h struct{}

func (h *h) ServeHTTP(http.ResponseWriter, *http.Request) {}

func TestHttp(t *testing.T) {

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.Handle("/id", new(h))

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Fatal(http.ListenAndServe(":80", nil))

}
