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
	tx *gorm.DB

	Epoch       epochStore
	Miner       minerStore
	Transaction transactionStore
	Event       eventStore
	Job         jobStore
}

const _createBatchSize = 1000

// NewStore creates a database store
func NewStore(dsn string, logMode logger.LogLevel) (*Store, error) {
	logger := logger.Default.LogMode(logMode)

	config := gorm.Config{
		CreateBatchSize: _createBatchSize,
		Logger:          logger,
	}

	db, err := gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		return nil, err
	}

	store := &Store{db: db}

	return store.setTransaction(db), nil
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
	config := gorm.Session{
		SkipDefaultTransaction: true,
	}

	tx := s.db.Session(&config).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	s.setTransaction(tx)

	return nil
}

// Commit commits the database transaction
func (s *Store) Commit() error {
	defer s.clearTransaction()

	err := s.tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

// Rollback rolls back the database transaction
func (s *Store) Rollback() error {
	defer s.clearTransaction()

	err := s.tx.Rollback().Error
	if err != nil {
		return err
	}
	return nil
}

// DatabaseSize returns the size of the database
func (s *Store) DatabaseSize() (int64, error) {
	var result int64

	err := s.tx.
		Raw("SELECT pg_database_size(current_database())").
		Scan(&result).
		Error

	if err != nil {
		return 0, err
	}

	return result, nil
}

// Close closes the database connection
func (s *Store) Close() error {
	conn, err := s.Conn()
	if err != nil {
		return err
	}

	return conn.Close()
}

func (s *Store) setTransaction(tx *gorm.DB) *Store {
	s.tx = tx

	s.Epoch = epochStore{db: tx}
	s.Miner = minerStore{db: tx}
	s.Transaction = transactionStore{db: tx}
	s.Event = eventStore{db: tx}

	s.Job = jobStore{db: s.db} // Use the default session

	return s
}

func (s *Store) clearTransaction() *Store {
	return s.setTransaction(s.db)
}
