package models

type Address struct {
	ID         int
	User       *User
	Name       string
	Street     string
	City       string
	State      string
	PostalCode string
}
