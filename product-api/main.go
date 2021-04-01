package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// create http server to be able to tune some params
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// every IP is bound to 9090, nil means use the defaultServeMUX, but now pass our sm
	s.ListenAndServe()
}
