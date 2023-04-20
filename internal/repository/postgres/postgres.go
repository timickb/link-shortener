package postgres

import "database/sql"

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateShortening(shortened, original string) error {
	_, err := r.db.Exec(`INSERT INTO links (shortened, original) VALUES ($1,$2)`, shortened, original)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetOriginal(shortened string) (string, error) {
	row := r.db.QueryRow(`SELECT * FROM links WHERE shortened=$1`, shortened)

	var short, original string
	if err := row.Scan(&short, &original); err != nil {
		return "", err
	}

	return original, nil
}
