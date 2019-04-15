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

type GetAddressRequest struct {
	ID int
}

type GetAddressResponse struct {
	Address *Address
}

type GetUserAddressesRequest struct {
	UserID int
}

type GetUserAddressesResponse struct {
	Addresses []*Address
}

type SetAddressRequest struct {
	Address *Address
}

type SetAddressResponse struct {
}

type DeleteAddressRequest struct {
	ID int
}

type DeleteAddressResponse struct {
}
