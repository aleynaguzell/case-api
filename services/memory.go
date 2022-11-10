package services

import (
	"case-api/storage/cache"
)

type MemoryService struct {
	db *cache.InMemory
}

func NewMemoryService(db *cache.InMemory) *MemoryService {
	return &MemoryService{
		db: db,
	}
}

func (s *MemoryService) Set(key, value string) error {
	err := s.db.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *MemoryService) Get(key string) (string, error) {
	value, err := s.db.Get(key)
	if err != nil {
		return "", err
	}
	return value, nil
}
