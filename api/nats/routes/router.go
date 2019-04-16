package routes

import (
	n "github.com/thegrandpackard/go-testing/nats"
)

func RegisterRoutes(n *n.NATSConnection) (err error) {

	_, err = n.QueueSubscribe("user.get", "queue_user.get", userGetHandler)
	if err != nil {
		return
	}

	_, err = n.QueueSubscribe("address.get", "queue_address.get", addressGetHandler)
	if err != nil {
		return
	}

	return
}
