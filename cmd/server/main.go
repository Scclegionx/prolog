package main

import (
	"github.com/Scclegionx/prolog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8081")
	log.Println("Server is starting on port 8081")
	log.Fatal(srv.ListenAndServe())
}
