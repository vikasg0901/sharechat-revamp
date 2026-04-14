package main

import (
	"log"
	"net/http"

	"github.com/sharechat/revamp/backend/router"
)

func main() {
	addr := ":8080"
	log.Printf("starting server on %s", addr)
	if err := http.ListenAndServe(addr, router.New()); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
