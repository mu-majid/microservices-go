package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye World!")

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
	fmt.Fprintf(rw, "Goodbye %s", d)
}
