package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config        *Config
	db            *sql.DB
	urlRepository *UrlRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DBUrl)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Url() *UrlRepository {
	if s.urlRepository != nil {
		return s.urlRepository
	}

	s.urlRepository = &UrlRepository{
		store: s,
	}
	return s.urlRepository
}
