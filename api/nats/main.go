package main

import (
	"log"
	"runtime"

	"github.com/thegrandpackard/go-testing/api/nats/routes"
	"github.com/thegrandpackard/go-testing/cases"
	n "github.com/thegrandpackard/go-testing/nats"
	"github.com/thegrandpackard/go-testing/storage"
)

var NATSUrl = "nats://localhost:4222"

func main() {
	log.Println("NATS Endpoint")

	// Init storage
	s, err := storage.Init()
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
