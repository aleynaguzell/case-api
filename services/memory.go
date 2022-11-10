package services

import (
	"case-api/storage/cache"
)

type MemoryService struct {
}

func NewMemoryService() *MemoryService {
	return &MemoryService{
	}
}

func (s *MemoryService) Set(key, value string) error {
	err := cache.New().Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *MemoryService) Get(key string) (string,error) {
	value, err := cache.New().Get(key)
	if err != nil {
		return "", err
	}
	return value, nil
}
