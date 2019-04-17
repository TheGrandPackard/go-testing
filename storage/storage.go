package storage

import (
	"database/sql"

	// Import sqlite3 database driver
	_ "github.com/mattn/go-sqlite3"
	// Import mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

type Storer interface {
	Close() (err error)
}

func Init(connectionString string) (storage *Storage, err error) {

	storage = &Storage{}
	storage.db, err = sql.Open("mysql", connectionString)
	if err != nil {
		return
	}

	err = storage.db.Ping()
	if err != nil {
		return
	}

	return
}

func InitMemoryDatabase() (storage *Storage, err error) {

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

func (s *Storage) Close() (err error) {
	return s.db.Close()
}
