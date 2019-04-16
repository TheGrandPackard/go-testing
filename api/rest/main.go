package main

import (
	"log"
	"runtime"

	"github.com/thegrandpackard/go-testing/api/rest/routes"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/rest"
	"github.com/thegrandpackard/go-testing/storage"
)

func main() {
	RESTAddress := "0.0.0.0:8080"

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

	// Init REST
	r, err := rest.Init(c)
	if err != nil {
		panic(err)
	}
	log.Printf("REST Server Initialized")

	// Register Routes
	err = routes.RegisterRoutes(r)
	if err != nil {
		panic(err)
	}
	log.Printf("Routes Registered")

	// Start REST Service
	err = r.StartServer(RESTAddress)
	if err != nil {
		panic(err)
	}

	runtime.Goexit()
}
