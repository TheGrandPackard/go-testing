package routes

import (
	"encoding/json"
	"testing"

	"github.com/nats-io/gnatsd/server"
	"github.com/stretchr/testify/assert"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
	n "github.com/thegrandpackard/go-testing/nats"
	"github.com/thegrandpackard/go-testing/storage"
)

func TestGetAddress(t *testing.T) {

	NATSUrl := "nats://localhost:4333"

	// Init embedded NATS server
	options := &server.Options{
		Host:           "127.0.0.1",
		Port:           4333,
		NoLog:          true,
		NoSigs:         true,
		MaxControlLine: 2048,
	}
	_, err := n.InitEmbeddedServer(options)
	assert.Nil(t, err, "NATS server should initialize")

	// Init storage
	s, err := storage.InitMemoryDatabase()
	assert.Nil(t, err, "Storage should initialize")

	// Init cases with storage
	c, err := cases.Init(s)
	assert.Nil(t, err, "Cases should initialize")

	// Init NATS
	nc, err := n.Init(NATSUrl, c)
	assert.Nil(t, err, "NATS should initialize")

	// Test Data
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
	assert.Nil(t, err, "User should be set")

	// Register Routes
	err = RegisterRoutes(nc)
	assert.Nil(t, err, "Routes should register")

	// Test request
	req := &models.GetAddressRequest{ID: 1}
	reqBytes, err := json.Marshal(req)
	assert.Nil(t, err, "Get Address request should marshal")

	respBytes, err := nc.Request("address.get", reqBytes)
	assert.Nil(t, err, "Get Address request should not error")

	resp := &models.GetAddressResponse{}
	err = json.Unmarshal(respBytes, &resp)
	assert.Nil(t, err, "Get Address request should unmarshal")

	assert.NotNil(t, resp.Address, "Address should not be nil")
	assert.Equal(t, a1.ID, resp.Address.ID, "Address ID should be equal")
	assert.Equal(t, a1.Name, resp.Address.Name, "Address Name should be equal")
	assert.Equal(t, a1.Street, resp.Address.Street, "Address Street should be equal")
	assert.Equal(t, a1.City, resp.Address.City, "Address City should be equal")
	assert.Equal(t, a1.State, resp.Address.State, "Address State should be equal")
	assert.Equal(t, a1.PostalCode, resp.Address.PostalCode, "Address PostalCode should be equal")
}
