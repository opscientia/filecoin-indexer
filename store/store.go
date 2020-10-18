package store

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Store handles all database operations
type Store struct {
	Db *gorm.DB
}

// New returns a new store from the connection string
func New(connStr string, logMode logger.LogLevel) (*Store, error) {
	config := gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	}

	db, err := gorm.Open(postgres.Open(connStr), &config)
	if err != nil {
		return nil, err
	}

	return &Store{Db: db}, nil
}

// Conn returns an underlying database connection
func (s *Store) Conn() *sql.DB {
	db, _ := s.Db.DB()
	return db
}

// Close closes the database connection
func (s *Store) Close() error {
	db, err := s.Db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
