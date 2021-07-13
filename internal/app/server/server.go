package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nazandr/ozonTest/internal/app/store"
)

type Server struct {
	Config *Config
	Echo   *echo.Echo
	Store  *store.Store
}

func New(config *Config) *Server {
	s := &Server{
		Config: config,
		Echo:   echo.New(),
	}
	s.configureRouter()
	return s
}

func (s *Server) Start() error {
	if err := s.configureStore(); err != nil {
		return err
	}
	return s.Echo.Start(s.Config.IP_addr)
}

func (s *Server) configureStore() error {
	store := store.New(s.Config.Store)
	if err := store.Open(); err != nil {
		return err
	}
	s.Store = store
	return nil
}

func (s *Server) configureRouter() {
	s.Echo.Use(middleware.RequestID())
	s.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "id=${id} method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	s.Echo.POST("/short", s.handleShort)
	s.Echo.POST("/long", s.handleLong)
}

func (s *Server) handleShort(c echo.Context) error {
	u := new(struct {
		Url string `json:"url"`
	})
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

func (s *Server) handleLong(c echo.Context) error {
	u := new(struct {
		Url string `json:"url"`
	})
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}
