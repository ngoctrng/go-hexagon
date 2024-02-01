package httpserver

import (
	"github.com/labstack/echo/v4"
	"hexagon/domain/book"
	"log"
	"net/http"
)

type Server struct {
	Router *echo.Echo
	// storage adapters
	BookStore book.Storage
}

func New() (*Server, error) {
	s := Server{Router: echo.New()}

	s.Router.POST("/api/books", s.CreateBook)
	s.Router.GET("/api/books/:isbn", s.GetBook)

	return &s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) handleError(c echo.Context, err error, status int) error {
	log.Println(err)
	return c.JSON(status, map[string]string{
		"message": http.StatusText(status),
	})
}
