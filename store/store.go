package store

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Store handles database operations
type Store struct {
	db *gorm.DB

	Epoch       epochStore
	Miner       minerStore
	Transaction transactionStore
	Event       eventStore
}

// New creates a store from the connection string
func New(connStr string, logMode logger.LogLevel) (*Store, error) {
	store := &Store{}

	config := gorm.Config{
		SkipDefaultTransaction: true,

		Logger: logger.Default.LogMode(logMode),
	}

	db, err := gorm.Open(postgres.Open(connStr), &config)
	if err != nil {
		return nil, err
	}

	return store.setSession(db), nil
}

// Conn returns the underlying database connection
func (s *Store) Conn() (*sql.DB, error) {
	return s.db.DB()
}

// Test checks the database connection
func (s *Store) Test() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

// Begin starts a database transaction
func (s *Store) Begin() error {
	db := s.db.Begin()
	if db.Error != nil {
		return db.Error
	}

	s.setSession(db)

	return nil
}

// Commit commits the database transaction
func (s *Store) Commit() error {
	err := s.db.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

// Rollback rolls back the database transaction
func (s *Store) Rollback() error {
	err := s.db.Rollback().Error
	if err != nil {
		return err
	}
	return nil
}

// Close closes the database connection
func (s *Store) Close() error {
	conn, err := s.Conn()
	if err != nil {
		return err
	}

	return conn.Close()
}

func (s *Store) setSession(db *gorm.DB) *Store {
	s.db = db

	s.Epoch = epochStore{db: db}
	s.Miner = minerStore{db: db}
	s.Transaction = transactionStore{db: db}
	s.Event = eventStore{db: db}

	return s
}
