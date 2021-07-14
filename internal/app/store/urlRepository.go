package store

import "github.com/nazandr/ozonTest/internal/app/models"

type UrlRepository struct {
	store *Store
}

func (r *UrlRepository) Create(url *models.URL) (*models.URL, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO urls(long, short) VALUES ($1, $2) RETURNING id",
		url.Long,
		url.Short,
	).Scan(&url.ID); err != nil {
		return nil, err
	}
	return url, nil
}

func (r *UrlRepository) FindByLong(LongUrl string) (*models.URL, error) {
	url := models.NewURL()
	if err := r.store.db.QueryRow(
		"SELECT id, long, short FROM urls WHERE long = $1",
		LongUrl,
	).Scan(
		&url.ID,
		&url.Long,
		&url.Short,
	); err != nil {
		return nil, err
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
