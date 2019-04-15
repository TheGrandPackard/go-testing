package storage

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thegrandpackard/go-testing/models"
)

func TestAddress(t *testing.T) {

	// Create storage interface
	storage := &Storage{}
	err := storage.Open()
	assert.Nil(t, err, "Database connection should be opened")

	// Create address table
	err = storage.CreateAddressTable()
	assert.Nil(t, err, "Address table should be created")

	// Set address
	u1 := &models.User{
		ID: 1,
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
	err = storage.SetAddress(a1)
	assert.Nil(t, err, "Address should be set")

	// Verify set address worked
	a2 := &models.Address{
		ID: 1,
	}
	err = storage.GetAddress(a2)
	assert.Nil(t, err, "Address should be get")

	// Delete address
	err = storage.DeleteAddress(a2)
	assert.Nil(t, err, "Address should be deleted")

	// Verify addresss was deleted
	a3 := &models.Address{
		ID: 1,
	}
	err = storage.GetAddress(a3)
	assert.Equal(t, sql.ErrNoRows, err, "Address should not be get")

	// Drop address table
	err = storage.DropAddressTable()
	assert.Nil(t, err, "Address table should be dropped")

	// Get user address without valid table
	err = storage.GetUserAddresses(&models.User{})
	assert.NotNil(t, err, "User Addresses should not be get")
}
