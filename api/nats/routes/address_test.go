package routes

import (
	"log"
	"testing"

	"github.com/thegrandpackard/go-testing/cases"
	n "github.com/thegrandpackard/go-testing/nats"
	"github.com/thegrandpackard/go-testing/storage"
)

func TestUserSubscription(t *testing.T) {

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

	// Init NATS
	nc, err := n.Init("nats://localhost:4222", c)
	if err != nil {
		panic(err)
	}

	var resp []byte
	resp, err = nc.Request("address", []byte(""))
	if err != nil {
		panic(err)
	}
	log.Printf("Received response: " + string(resp))
}
