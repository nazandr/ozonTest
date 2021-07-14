package store_test

import (
	"testing"

	"github.com/nazandr/ozonTest/internal/app/models"
	"github.com/nazandr/ozonTest/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUrlRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, "host=localhost dbname=short_url_test sslmode=disable")
	defer teardown("urls")

	u, err := s.Url().Create(&models.URL{
		Long:  "longurl.com/example",
		Short: "shorturl",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUrlRepository_FindByLong(t *testing.T) {
	s, teardown := store.TestStore(t, "host=localhost dbname=short_url_test sslmode=disable")
	defer teardown("urls")

	url := &models.URL{
		Long:  "longurl.com/example",
		Short: "shorturl",
	}

	_, err := s.Url().FindByLong(url.Long)
	assert.Error(t, err)

	_, err = s.Url().Create(url)
	assert.NoError(t, err)
	u, err := s.Url().FindByLong(url.Long)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUrlRepository_FindByShort(t *testing.T) {
	s, teardown := store.TestStore(t, "host=localhost dbname=short_url_test sslmode=disable")
	defer teardown("urls")

	url := &models.URL{
		Long:  "longurl.com/example",
		Short: "shorturl",
	}

	_, err := s.Url().FindByShort(url.Short)
	assert.Error(t, err)

	_, err = s.Url().Create(url)
	assert.NoError(t, err)
	u, err := s.Url().FindByShort(url.Short)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
