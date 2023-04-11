package main

import (
	"awesomeProject1/service1"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", service1.Handler1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
