package main

import (
	"log"
	"runtime"

	"github.com/thegrandpackard/go-testing/api/nats/routes"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
	n "github.com/thegrandpackard/go-testing/nats"
	"github.com/thegrandpackard/go-testing/storage"
)

func main() {
	NATSUrl := "nats://localhost:4222"

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

	// Test Data
	{
		u1 := &models.User{
			ID:   1,
			Name: "Joshua",
			Age:  29,
		}
		a1 := &models.Address{
			ID:         1,
			User:       u1,
			Name:       "Home",
			Street:     "123 Main Street",
			City:       "Greenville",
			State:      "Indiana",
			PostalCode: "12345",
		}
		u1.Addresses = append(u1.Addresses, a1)
		_, err = c.SetUser(&models.SetUserRequest{User: u1})
		if err != nil {
			log.Printf("Error setting user: %s", err.Error())
		} else {
			log.Printf("Set User: %+v", u1)
		}
	}

	// Register Routes
	err = routes.RegisterRoutes(nc)
	if err != nil {
		panic(err)
	}
	log.Printf("Routes Registered")

	runtime.Goexit()
}
