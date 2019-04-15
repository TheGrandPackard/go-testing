package cases

import "github.com/thegrandpackard/go-testing/models"

type UserCases interface {
	GetUser(*models.GetUserRequest) (*models.GetUserResponse, error)
	SetUser(*models.SetUserRequest) (*models.SetUserResponse, error)
	DeleteUser(*models.DeleteUserRequest) (*models.DeleteUserResponse, error)
}

func (c *Cases) GetUser(req *models.GetUserRequest) (resp *models.GetUserResponse, err error) {

	resp = &models.GetUserResponse{
		User: &models.User{ID: req.ID},
	}
	err = c.storage.GetUser(resp.User)

	return
}

func (c *Cases) SetUser(req *models.SetUserRequest) (resp *models.SetUserResponse, err error) {

	resp = &models.SetUserResponse{}
	err = c.storage.SetUser(req.User)

	return
}

func (c *Cases) DeleteUser(req *models.DeleteUserRequest) (resp *models.DeleteUserResponse, err error) {

	resp = &models.DeleteUserResponse{}
	err = c.storage.DeleteUser(&models.User{ID: req.ID})

	return
}
