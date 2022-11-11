package cache

import (
	"case-api/model/errormessage"
	"case-api/pkg/logger"
	"sync"

	"github.com/pkg/errors"
)

type InMemory struct {
	m    *sync.Mutex
	data map[string]string
}

// New creates inmemory store
func New() *InMemory {
	return &InMemory{
		data: make(map[string]string),
		m:    &sync.Mutex{},
	}
}

func (s *InMemory) Set(key, value string) error {

	if len(key) == 0 || len(value) == 0 {
		err := errors.New(errormessage.KeyOrValueEmpty)
		logger.CustomError(err)
		return err
	}

	s.m.Lock()
	defer s.m.Unlock()

	s.data[key] = value
	return nil
}

func (s *InMemory) Get(key string) (string, error) {

	err := errors.New(errormessage.KeyNotFound)
	if len(key) == 0 {
		err = errors.New(errormessage.KeyEmpty)
		logger.CustomError(err)
		return "", err
	}

	s.m.Lock()
	defer s.m.Unlock()

	if val, ok := s.data[key]; ok {
		return val, nil
	}

	return "", err
}
