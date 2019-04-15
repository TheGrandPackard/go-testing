package main

import (
	"log"

	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
	"github.com/thegrandpackard/go-testing/storage"
)

/* This is a sample main to demonstrate how to hook into cases (which then in turn invoke storage, other APIs, etc.)
   This would be an ideal starting point for integrating a REST interface, NATS endpoint, GRPC gateway, etc.
*/
func main() {

	// Init storage
	s, err := storage.Init()
	if err != nil {
		panic(err)
	}

	// Init cases with storage
	c, err := cases.Init(s)
	if err != nil {
		panic(err)
	}

	// Create user
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

	// Get user
	{
		getResp, err := c.GetUser(&models.GetUserRequest{ID: 1})
		if err != nil {
			log.Printf("Error getting user: %s", err.Error())
		} else {
			log.Printf("Got User: %+v", getResp.User)
		}
	}

	// Delete user
	{
		_, err = c.DeleteUser(&models.DeleteUserRequest{ID: 1})
		if err != nil {
			log.Printf("Error getting user: %s", err.Error())
		} else {
			log.Printf("Deleted User: %d", 1)
		}
	}

	// Get deleted user
	{
		getResp, err := c.GetUser(&models.GetUserRequest{ID: 1})
		if err != nil {
			log.Printf("Error getting user: %s", err.Error())
		} else {
			log.Printf("Got User: %+v", getResp.User)
		}
	}
}
