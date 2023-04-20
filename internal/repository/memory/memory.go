package memory

import "errors"

type Repository struct {
	data map[string]string
}

func New() *Repository {
	return &Repository{data: map[string]string{}}
}

func (r *Repository) CreateShortening(shortened, original string) error {
	if _, ok := r.data[shortened]; ok {
		return nil
	}
	r.data[shortened] = original
	return nil
}

func (r *Repository) GetOriginal(shortened string) (string, error) {
	original, ok := r.data[shortened]
	if !ok {
		return "", errors.New("err not found")
	}

	return original, nil
}
