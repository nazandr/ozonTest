package server

import (
	"testing"
)

func TestServer(t *testing.T) *Server {
	t.Helper()

	conf := NewConfig()

	s := New(conf)
	return s
}
