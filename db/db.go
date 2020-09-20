package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database wraps a backend db
type Database struct {
	db *gorm.DB
}

// New returns a new sqlite3 database wrapper
func New(dbPath string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Deposit{}, &Mint{})
	return &Database{db}, nil

}

func (d *Database) Close() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
