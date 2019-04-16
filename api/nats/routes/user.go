package routes

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
)

func userGetHandler(msg *nats.Msg, c *cases.Cases) (response []byte) {

	req := &models.GetUserRequest{}
	err := json.Unmarshal(msg.Data, &req)
	if err != nil {
		response = []byte("{ \"error\": \"Error unmarshalling request: " + err.Error() + "\" }")
		return
	}

	resp, err := c.GetUser(req)
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
		response = []byte("{ \"error\": \"Error getting user: " + err.Error() + "\" }")
		return
	}

	response, err = json.Marshal(resp)
	if err != nil {
		response = []byte("{ \"error\": \"Error marshalling response: " + err.Error() + "\" }")
		return
	}

	return
}
