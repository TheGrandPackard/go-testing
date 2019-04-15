package models

type User struct {
	ID        int
	Name      string
	Age       int
	Addresses []*Address
}

type GetUserRequest struct {
	ID   int
	Name string
}

type GetUserResponse struct {
	User *User
}

type SetUserRequest struct {
	User *User
}

type SetUserResponse struct {
}

type DeleteUserRequest struct {
	ID int
}

type DeleteUserResponse struct {
}
