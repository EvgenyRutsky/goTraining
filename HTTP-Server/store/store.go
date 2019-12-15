package store

import (
	"errors"
)

type Store struct {
	config *Config
	repository Repository
}

func New(config *Config) (*Store, error) {

	if config.DB == "mongodb" {
		return &Store{
			config: config,
			repository: NewMongoRepository(config),
		}, nil
	}

	return nil, errors.New("error during the storage creation")
}

func (s *Store) Book() Repository {
	return s.repository
}