package handlers

import (
	"learn-ms/product-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
    return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        p.getProducts(w, r)
        return
    }

    w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
    pl := data.GetProducts()
    // encode will be faster than Marshal
    // data, err := json.Marshal(pl)
    err := pl.ToJSON(w)

    if err != nil {
        http.Error(w, "Unable to Marshal json", http.StatusInternalServerError)
    }

    // w.Write(data)
}
