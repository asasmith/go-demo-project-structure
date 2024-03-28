package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asasmith/go-demo-project-structure/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 1111
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

