package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("que facemos")
	})

	http.HandleFunc("/quefacemos", func(w http.ResponseWriter, r *http.Request) {
		log.Println("non sei")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "ooooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "que facemos? %s", d)
	})

	http.ListenAndServe(":9090", nil)
}
