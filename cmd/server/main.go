package main

import (
	"log"

	"github.com/jennwah/distributed-services-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
