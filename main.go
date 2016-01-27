package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ericchiang/simplehttp/handlers"
)

func main() {
	name := flag.String("name", "eric", "The person to say hello to!")
	addr := flag.String("addr", "127.0.01.1:8080", "The address to listen for HTTP requests on.")
	flag.Parse()

	http.Handle("/hello", handlers.SayHello(*name))
	log.Printf("Go to http://%s/hello", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
