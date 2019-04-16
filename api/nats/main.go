package main

import (
	"log"
	"runtime"

	"github.com/thegrandpackard/go-testing/api/nats/routes"
	"github.com/thegrandpackard/go-testing/cases"
	n "github.com/thegrandpackard/go-testing/nats"
	"github.com/thegrandpackard/go-testing/storage"
)

func main() {
	NATSUrl := "nats://localhost:4222"

	log.Println("NATS Endpoint")

	// Init storage
	s, err := storage.Init("test:test@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	log.Printf("Storage Initialized")

	// Init cases with storage
	c, err := cases.Init(s)
	if err != nil {
		panic(err)
	}
	log.Printf("Cases Initialized")

	// Init NATS
	nc, err := n.Init(NATSUrl, c)
	if err != nil {
		panic(err)
	}
	log.Printf("NATS Connected to: %s", NATSUrl)

	// Register Routes
	err = routes.RegisterRoutes(nc)
	if err != nil {
		panic(err)
	}
	log.Printf("Routes Registered")

	runtime.Goexit()
}
