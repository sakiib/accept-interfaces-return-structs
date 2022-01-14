package db

import "errors"

type Store struct {
	store map[string]string
}

func NewStore(st map[string]string) (*Store, error) {
	if st == nil {
		return nil, errors.New("store not initialzed")
	}

	return &Store{
		store: st,
	}, nil
}

func (s *Store) Insert(key, value string) error {
	if len(key) == 0 || len(value) == 0 {
		return errors.New("key/value empty")
	}

	s.store[key] = value
	return nil
}

func (s *Store) Get(key string) (string, error) {
	if val, ok := s.store[key]; ok {
		return val, nil
	}
	return "", errors.New("user not found")
}
