package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mu-majid/microservices-go/product-api/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// every IP is bound to 9090, nil means use the defaultServeMUX, but now pass our sm
	http.ListenAndServe(":9090", sm)
}
