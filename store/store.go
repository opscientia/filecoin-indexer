package store

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Store handles database operations
type Store struct {
	Db *gorm.DB
}

// New creates a store from the connection string
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

// Conn returns the underlying database connection
func (s *Store) Conn() (*sql.DB, error) {
	return s.Db.DB()
}

// Close closes the database connection
func (s *Store) Close() error {
	conn, err := s.Conn()
	if err != nil {
		return err
	}

	return conn.Close()
}
