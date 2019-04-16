package main

import (
	"log"
	"runtime"

	"github.com/thegrandpackard/go-testing/api/rest/routes"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
	"github.com/thegrandpackard/go-testing/rest"
	"github.com/thegrandpackard/go-testing/storage"
)

func main() {
	RESTAddress := "0.0.0.0:8080"

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
