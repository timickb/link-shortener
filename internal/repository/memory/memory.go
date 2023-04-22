package memory

import (
	"errors"
	"sync"
)

// There must be only one instance of repo for all API
var (
	singleton Repository
	once      sync.Once
)

type Repository struct {
	sync.RWMutex
	data map[string]string
}

func New() *Repository {
	once.Do(func() {
		singleton = Repository{data: map[string]string{}}
	})
	return &singleton
}

func (r *Repository) CreateShortening(shortened, original string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.data[shortened]; ok {
		return nil
	}
	r.data[shortened] = original
	return nil
}

func (r *Repository) GetOriginal(shortened string) (string, error) {
	r.RLock()
	defer r.RUnlock()

	original, ok := r.data[shortened]
	if !ok {
		return "", errors.New("err not found")
	}

	return original, nil
}
