package store

import (
	"database/sql"

	"github.com/nazandr/ozonTest/internal/app/models"
)

type UrlRepository struct {
	store *Store
}

func (r *UrlRepository) Create(url *models.URL) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO urls(long) VALUES ($1) RETURNING id",
		url.Long,
	).Scan(&url.ID); err != nil {
		return err
	}
	return nil
}

func (r *UrlRepository) UpdateShort(url *models.URL) error {
	if err := r.store.db.QueryRow("UPDATE urls SET short = $1 WHERE id = $2", url.Short, url.ID).Err(); err != nil {
		return err
	}
	return nil
}

func (r *UrlRepository) FindByLong(LongUrl string) (*models.URL, error) {
	url := models.NewURL()
	sh := sql.NullString{}
	if err := r.store.db.QueryRow(
		"SELECT id, long, short FROM urls WHERE long = $1",
		LongUrl,
	).Scan(
		&url.ID,
		&url.Long,
		&sh,
	); err != nil {
		return nil, err
	}
	if sh.Valid {
		url.Short = sh.String
	} else {
		url.Short = ""
	}

	return url, nil
}

func (r *UrlRepository) FindByShort(ShortUrl string) (*models.URL, error) {
	url := models.NewURL()
	if err := r.store.db.QueryRow(
		"SELECT id, long, short FROM urls WHERE short = $1",
		ShortUrl,
	).Scan(
		&url.ID,
		&url.Long,
		&url.Short,
	); err != nil {
		return nil, err
	}

	return url, nil
}
