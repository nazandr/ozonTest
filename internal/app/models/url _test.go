package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl_Validation(t *testing.T) {
	testCases := []struct {
		name    string
		field   string
		isValid bool
	}{
		{
			name:    "valid",
			field:   "example.com/long",
			isValid: true,
		},
		{
			name:    "empty",
			field:   "",
			isValid: false,
		},
		{
			name:    "invalid",
			field:   "invalid",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := NewURL()
			url.Long = tc.field
			err := url.Validation()
			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
