package service1

import (
	"fmt"
	"net/http"
)

func Handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world222!")
}
