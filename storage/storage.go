package storage

import (
	"database/sql"

	// Import sqlite3 database driver
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) Open() (err error) {
	s.db, err = sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	return
}
