package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nazandr/ozonTest/internal/app/models"
	"github.com/nazandr/ozonTest/internal/app/store"
)

type Server struct {
	Config *Config
	Echo   *echo.Echo
	Store  *store.Store
}

type Resp struct {
	Url string `json:"url"`
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
	u := new(Resp)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	url := models.NewURL()
	url.Long = u.Url
	if err := url.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if l, _ := s.Store.Url().FindByLong(u.Url); l != nil {
		r := Resp{
			Url: l.Short,
		}
		return c.JSON(http.StatusOK, r)
	}

	if err := s.Store.Url().Create(url); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	url.Shortener()
	s.Store.Url().UpdateShort(url)

	return c.JSON(http.StatusOK, Resp{Url: url.Short})
}

func (s *Server) handleLong(c echo.Context) error {
	u := new(Resp)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, err := s.Store.Url().FindByShort(u.Url)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r := Resp{
		Url: url.Long,
	}
	return c.JSON(http.StatusOK, r)
}
