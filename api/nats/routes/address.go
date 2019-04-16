package routes

import (
	"github.com/nats-io/nats"
	"github.com/thegrandpackard/go-testing/cases"
)

func addressHandler(msg *nats.Msg, c *cases.Cases) (response []byte) {
	response = []byte("Address: 1234 Main Street")
	return
}
