package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(connStr string) (*Store, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Store{Db: db}, nil
}

type Store struct {
	Db *gorm.DB
}

func (s *Store) Close() error {
	db, err := s.Db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
