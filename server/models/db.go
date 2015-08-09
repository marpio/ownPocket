package models

import (
	"database/sql"
	// need this here
	_ "github.com/mattn/go-sqlite3"
)

// Datastore interface
type Datastore interface {
	SearchBookmarks(searchPhrase string) ([]*Bookmark, error)
	AddBookmark(bookmark *Bookmark) (*Bookmark, error)
}

// DB struct
type DB struct {
	*sql.DB
}

// NewDB creates  new DB conn
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
