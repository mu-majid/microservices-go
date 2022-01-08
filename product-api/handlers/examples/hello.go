package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World!")

	// read from request
	d, err := ioutil.ReadAll(r.Body)

	// handle err
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("OOPS"))
		// or simply
		http.Error(rw, "OOPS", http.StatusBadRequest)
		return
	}

	log.Printf("Data Is %s \n", d)

	// write back to user
	fmt.Fprintf(rw, "Hello %s", d)
}
