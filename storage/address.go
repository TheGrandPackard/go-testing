package storage

import (
	"database/sql"
	"errors"

	"github.com/thegrandpackard/go-testing/models"
)

type AddressStorer interface {
	CreateAddressTable() error
	DropAddressTable() error

	GetAddress(a *models.Address) error
	GetUserAddresses(u *models.User) error
	SetAddress(a *models.Address) error
	DeleteAddress(a *models.Address) error
}

func (s *Storage) CreateAddressTable() (err error) {
	_, err = s.db.Exec(`
	CREATE TABLE address (
		id INT UNSIGNED NOT NULL,
		user_id INT UNSIGNED NOT NULL,
		name VARCHAR(128) DEFAULT ' ',
		street VARCHAR(128) DEFAULT ' ',
		city VARCHAR(64) DEFAULT ' ',
		state VARCHAR(64) DEFAULT ' ',
		postal_code VARCHAR(64) DEFAULT ' ',
		PRIMARY KEY (id)
	  );`)
	return
}

func (s *Storage) GetAddress(a *models.Address) (err error) {
	if a.User == nil {
		a.User = &models.User{}
	}
	err = s.db.QueryRow("SELECT user_id, name, street, city, state, postal_code FROM address WHERE id = ?;", a.ID).
		Scan(&a.User.ID, &a.Name, &a.Street, &a.City, &a.State, &a.PostalCode)
	return
}

func (s *Storage) GetUserAddresses(u *models.User) (addresses []*models.Address, err error) {
	var rows *sql.Rows
	rows, err = s.db.Query("SELECT id, name, street, city, state, postal_code FROM address WHERE user_id = ?;", u.ID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		a := &models.Address{}
		err = rows.Scan(&a.ID, &a.Name, &a.Street, &a.City, &a.State, &a.PostalCode)
		if err != nil {
			return
		}
		addresses = append(addresses, a)
	}
	return
}

func (s *Storage) SetAddress(a *models.Address) (err error) {
	if a.User == nil {
		return errors.New("Cannot set address for user with invalid user")
	}
	_, err = s.db.Exec("REPLACE INTO address (id, user_id, name, street, city, state, postal_code) VALUES (?, ?, ?, ?, ?, ?, ?);",
		a.ID, a.User.ID, a.Name, a.Street, a.City, a.State, a.PostalCode)
	return
}

func (s *Storage) DeleteAddress(a *models.Address) (err error) {
	_, err = s.db.Exec("DELETE FROM address WHERE id = ?;", a.ID)
	return
}

func (s *Storage) DropAddressTable() (err error) {
	_, err = s.db.Exec(`DROP TABLE address;`)
	return
}
