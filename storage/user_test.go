package storage

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thegrandpackard/go-testing/models"
)

func TestUser(t *testing.T) {

	// Create storage interface
	storage, err := InitMemoryDatabase()
	assert.Nil(t, err, "Database connection should be opened")

	// Set person
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
	err = storage.SetUser(u1)
	assert.Nil(t, err, "User should be set")

	// Verify set person worked
	u2 := &models.User{
		ID: 1,
	}
	err = storage.GetUser(u2)
	assert.Nil(t, err, "User should be get")

	assert.Equal(t, 1, u2.ID, "ID should be 1")
	assert.Equal(t, "Joshua", u2.Name, "Name should be Joshua")
	assert.Equal(t, 29, u2.Age, "Age should be 29")

	assert.Equal(t, 29, u2.Age, "Address ID should be 29")

	// Delete person
	err = storage.DeleteUser(u2)
	assert.Nil(t, err, "User should be deleted")

	// Verify person was deleted
	u3 := &models.User{
		ID: 1,
	}
	err = storage.GetUser(u3)
	assert.Equal(t, sql.ErrNoRows, err, "User should not be get")

	// Set user with bad address
	u4 := &models.User{}
	u4.Addresses = append(u4.Addresses, &models.Address{})
	err = storage.SetUser(u4)
	assert.NotNil(t, err, "User with bad Address should return error")

	// Drop person table
	err = storage.DropUserTable()
	assert.Nil(t, err, "User table should be dropped")

	// Close the database
	storage.Close()
}
