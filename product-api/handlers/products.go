package handlers

import (
	"log"
	"net/http"

	"github.com/mu-majid/microservices-go/product-api/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Catch All
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusNotImplemented)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle GET Products")
	// fetch the products from the datastore
	lp := data.GetProducts()
	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode Data to JSON", http.StatusInternalServerError)
		return
	}
}
