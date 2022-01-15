package inMemRepository

import (
	"errors"
	"fmt"
	"sync"
)

type Repository interface {
	Get(key string) (string, error)
	Set(key string, value string)
}

type repository struct {
	storage map[string]string
	mu      *sync.Mutex
}

func NewRepository() *repository {
	return &repository{
		storage: make(map[string]string),
		mu:      &sync.Mutex{},
	}
}

func (r *repository) Get(key string) (string, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	value, exist := r.storage[key]
	if !exist {

		return "", errors.New(fmt.Sprintf("%s not found", key))
	}

	return value, nil
}

func (r *repository) Set(key, value string) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[key] = value

}
