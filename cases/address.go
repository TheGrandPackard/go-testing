package cases

import "github.com/thegrandpackard/go-testing/models"

type AddressCases interface {
	GetAddress(*models.GetAddressRequest) (*models.GetAddressResponse, error)
	GetUserAddresses(*models.GetUserAddressesRequest) (*models.GetUserAddressesResponse, error)
	SetAddress(*models.SetAddressRequest) (*models.SetAddressResponse, error)
	DeleteAddress(*models.DeleteAddressRequest) (*models.DeleteAddressResponse, error)
}

func (c *Cases) GetAddress(req *models.GetAddressRequest) (resp *models.GetAddressResponse, err error) {

	resp = &models.GetAddressResponse{
		Address: &models.Address{ID: req.ID},
	}
	err = c.storage.GetAddress(resp.Address)

	return
}

func (c *Cases) GetUserAddresses(req *models.GetUserAddressesRequest) (resp *models.GetUserAddressesResponse, err error) {

	var addresses []*models.Address
	resp = &models.GetUserAddressesResponse{}
	addresses, err = c.storage.GetUserAddresses(&models.User{ID: req.UserID})
	if err != nil {
		resp.Addresses = addresses
	}

	return
}

func (c *Cases) SetAddress(req *models.SetAddressRequest) (resp *models.SetAddressResponse, err error) {

	resp = &models.SetAddressResponse{}
	err = c.storage.SetAddress(req.Address)

	return
}

func (c *Cases) DeleteAddress(req *models.DeleteAddressRequest) (resp *models.DeleteAddressResponse, err error) {

	resp = &models.DeleteAddressResponse{}
	err = c.storage.DeleteAddress(&models.Address{ID: req.ID})

	return
}
