package cases

import "github.com/thegrandpackard/go-testing/storage"

type Cases struct {
	storage *storage.Storage
}

func Init(s *storage.Storage) (cases *Cases, err error) {
	return &Cases{storage: s}, nil
}
