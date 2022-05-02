package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("non sei")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "ooooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "que facemos? %s", d)
}
