package handlers

import (
	"fmt"
	"net/http"
)

func SayHello(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}
