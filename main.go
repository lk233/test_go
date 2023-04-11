package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world222!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
