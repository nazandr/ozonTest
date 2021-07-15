package models

import (
	"bytes"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	l    = uint64(62)
	base = "ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

type URL struct {
	ID    uint64
	Long  string
	Short string
}

func NewURL() *URL {
	return &URL{
		ID:    0,
		Long:  "",
		Short: "",
	}
}

func (url *URL) Validation() error {
	return validation.ValidateStruct(url,
		validation.Field(&url.Long, validation.Required, is.URL),
	)
}

func encode(id uint64, buf *bytes.Buffer) {
	if id/l != 0 {
		encode(id/l, buf)
	}
	buf.WriteByte(base[id%l])
}

func (url *URL) Shortener() {
	// big.NewInt(1).Text(62)
	b := new(bytes.Buffer)
	encode(url.ID, b)
	url.Short = b.String()
}
