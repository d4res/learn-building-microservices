package main

import (
	"learn-ms/product-api/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
    l := log.New(os.Stdout, "products-api", log.LstdFlags)

    ph := handlers.NewProducts(l)

    sm := http.NewServeMux()
    sm.Handle("/", ph)

    s := &http.Server{
        Addr: ":9090",
        Handler: sm,
    }
    s.ListenAndServe()
}
