package main

import (
	"context"
	"learn-ms/ep1/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
    l := log.New(os.Stdout, "product-api", log.LstdFlags)
    helloHandler := handlers.NewHello(l)

    // NewServeMux creates a new ServerMux which implements the Handler interface
    // which implements the func ServeHTTP
    sm := http.NewServeMux()

    // register router to the sm
    sm.Handle("/", helloHandler)


    s := &http.Server{
        Addr: ":9090",
        Handler: sm,
        IdleTimeout: 120 * time.Second,
        ReadTimeout: 1 * time.Second,
        WriteTimeout: 1 * time.Second,
    }
    
    go func() {
        err := s.ListenAndServe()
        if err != nil {
            l.Fatal(err)
        }
    }()
    
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)
    sig := <-sigChan
    l.Println("received terminate, shutdown", sig)

    // ListenAndServe will create a http.Server and using some default config
    // if the second parameter(the handler) is nil, ListenAndServe uses the defaultservermux

    //http.ListenAndServe(":9090", sm)

    tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
    s.Shutdown(tc)
}
