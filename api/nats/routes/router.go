package routes

import (
	n "github.com/thegrandpackard/go-testing/nats"
)

func RegisterRoutes(n *n.NATSConnection) (err error) {

	_, err = n.QueueSubscribe("user", "user_queue", userHandler)
	if err != nil {
		return
	}

	_, err = n.QueueSubscribe("address", "address_queue", addressHandler)
	if err != nil {
		return
	}

	return
}
