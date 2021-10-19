package main

import (
	"fmt"
	"log"
	"net/http"
)

type greeter struct {
}

func (g greeter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Greetings!")
}

func main() {
	if err := http.ListenAndServe("127.0.0.1:8080", &greeter{}); err != nil {
		log.Fatalf("Server failed to start %s", err)
	}
}
