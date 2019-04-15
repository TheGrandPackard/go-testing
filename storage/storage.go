package storage

import (
	"database/sql"

	// Import sqlite3 database driver
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func Init() (storage *Storage, err error) {

	storage = &Storage{}
	storage.db, err = sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		return
	}

	// Init tables
	storage.CreateUserTable()
	storage.CreateAddressTable()

	return
}

func (s *Storage) Close() {
	s.db.Close()
}
