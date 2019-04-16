package routes

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
)

func addressGetHandler(msg *nats.Msg, c *cases.Cases) (response []byte) {

	req := &models.GetAddressRequest{}
	err := json.Unmarshal(msg.Data, &req)
	if err != nil {
		response = []byte("{ \"error\": \"Error unmarshalling request: " + err.Error() + "\" }")
		return
	}

	resp, err := c.GetAddress(req)
	if err != nil {
		log.Printf("Error getting address: %s", err.Error())
		response = []byte("{ \"error\": \"Error getting address: " + err.Error() + "\" }")
		return
	}

	response, err = json.Marshal(resp)
	if err != nil {
		response = []byte("{ \"error\": \"Error marshalling response: " + err.Error() + "\" }")
		return
	}

	return
}
