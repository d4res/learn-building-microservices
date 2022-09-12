package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// a HTTP handler is an interface
type Hello struct {
    l *log.Logger
}

// 建立l， 用于之后的配置注入
func NewHello(l *log.Logger) *Hello {
    return &Hello{l: l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.l.Println("Hello World")

    data, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Oops", http.StatusBadRequest)
        // w.WriteHeader(http.StatusBadRequest)
    }

    fmt.Fprintf(w, "Hello %s", data)
}
