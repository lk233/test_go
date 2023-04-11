package main

import (
	"github.com/lk233/test_go/service1"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", service1.Handler1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
