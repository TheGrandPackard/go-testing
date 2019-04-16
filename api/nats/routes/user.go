package routes

import (
	"github.com/nats-io/nats"
	"github.com/thegrandpackard/go-testing/cases"
)

func userHandler(msg *nats.Msg, c *cases.Cases) (response []byte) {
	response = []byte("User: Josh")
	return
}
