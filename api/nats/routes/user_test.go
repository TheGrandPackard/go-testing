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

func TestGetUser(t *testing.T) {

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
	s, err := storage.Init()
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
	req := &models.GetUserRequest{ID: 1}
	reqBytes, err := json.Marshal(req)
	assert.Nil(t, err, "Get user request should marshal")

	respBytes, err := nc.Request("user.get", reqBytes)
	assert.Nil(t, err, "Get user request should not error")

	resp := &models.GetUserResponse{}
	err = json.Unmarshal(respBytes, &resp)
	assert.Nil(t, err, "Get user request should unmarshal")

	assert.NotNil(t, resp.User, "User should not be nil")
	assert.Equal(t, u1.ID, resp.User.ID, "User ID should be equal")
	assert.Equal(t, u1.Name, resp.User.Name, "User name should be equal")
	assert.Equal(t, u1.Age, resp.User.Age, "User age should be equal")

	assert.NotNil(t, resp.User.Addresses[0], "Address should not be nil")
	assert.Equal(t, a1.ID, resp.User.Addresses[0].ID, "Address ID should be equal")
	assert.Equal(t, a1.Name, resp.User.Addresses[0].Name, "Address Name should be equal")
	assert.Equal(t, a1.Street, resp.User.Addresses[0].Street, "Address Street should be equal")
	assert.Equal(t, a1.City, resp.User.Addresses[0].City, "Address City should be equal")
	assert.Equal(t, a1.State, resp.User.Addresses[0].State, "Address State should be equal")
	assert.Equal(t, a1.PostalCode, resp.User.Addresses[0].PostalCode, "Address PostalCode should be equal")
}
