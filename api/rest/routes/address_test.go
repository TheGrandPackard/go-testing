package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
	"github.com/thegrandpackard/go-testing/rest"
	"github.com/thegrandpackard/go-testing/storage"
)

func TestGetAddress(t *testing.T) {

	RESTAddress := "0.0.0.0:8181"

	// Init storage
	s, err := storage.Init()
	assert.Nil(t, err, "Storage should initialize")

	// Init cases with storage
	c, err := cases.Init(s)
	assert.Nil(t, err, "Cases should initialize")

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

	// Init REST
	r, err := rest.Init(c)
	assert.Nil(t, err, "Rest server should initialize")

	// Register Routes
	err = RegisterRoutes(r)
	assert.Nil(t, err, "Rest routes should register")

	// Start REST Service
	err = r.StartServer(RESTAddress)
	assert.Nil(t, err, "Rest service should start")

	// Test request
	response, err := http.Get("http://" + RESTAddress + "/address/1")
	assert.Nil(t, err, "HTTP request should get response")

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err, "HTTP response body should get read")

	resp := &models.GetAddressResponse{}
	err = json.Unmarshal(body, &resp)
	assert.Nil(t, err, "Get user request should unmarshal")

	assert.NotNil(t, resp.Address, "Address should not be nil")
	assert.Equal(t, a1.ID, resp.Address.ID, "Address ID should be equal")
	assert.Equal(t, a1.Name, resp.Address.Name, "Address Name should be equal")
	assert.Equal(t, a1.Street, resp.Address.Street, "Address Street should be equal")
	assert.Equal(t, a1.City, resp.Address.City, "Address City should be equal")
	assert.Equal(t, a1.State, resp.Address.State, "Address State should be equal")
	assert.Equal(t, a1.PostalCode, resp.Address.PostalCode, "Address PostalCode should be equal")
}
