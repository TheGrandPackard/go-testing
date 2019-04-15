package storage

import (
	"github.com/thegrandpackard/go-testing/models"
)

type UserStorer interface {
	CreateUserTable() error
	DropUserTable() error

	GetUser(u *models.User) error
	SetUser(u *models.User) error
	DeleteUser(u *models.User) error
}

func (s *Storage) CreateUserTable() (err error) {
	_, err = s.db.Exec(`
	CREATE TABLE user (
		id INT UNSIGNED NOT NULL,
		name VARCHAR(64) DEFAULT ' ',
		age INT UNSIGNED NOT NULL DEFAULT 0,
		PRIMARY KEY (id)
	  );`)
	return
}

func (s *Storage) GetUser(u *models.User) (err error) {
	err = s.db.QueryRow("SELECT name, age FROM user WHERE id = ?;", u.ID).Scan(&u.Name, &u.Age)
	if err == nil {
		var addresses []*models.Address
		addresses, err = s.GetUserAddresses(u)
		if err != nil {
			return
		}
		u.Addresses = addresses
	}
	return
}

func (s *Storage) SetUser(u *models.User) (err error) {
	_, err = s.db.Exec("REPLACE INTO user (id, name, age) VALUES (?, ?, ?);", u.ID, u.Name, u.Age)
	if err == nil {
		for _, a := range u.Addresses {
			err = s.SetAddress(a)
			if err != nil {
				return
			}
		}
	}
	return
}

func (s *Storage) DeleteUser(u *models.User) (err error) {
	_, err = s.db.Exec("DELETE FROM user WHERE id = ?;", u.ID)
	return
}

func (s *Storage) DropUserTable() (err error) {
	_, err = s.db.Exec(`DROP TABLE user;`)
	return
}
