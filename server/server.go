package server

import (
	"echo_ifElse/pkg/handler"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"log"
	"time"
)

type Server struct {
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{handler: handler}
}

func (s *Server) Run() error {
	e := echo.New()
	s.handler.InitRoutes(e)
	c := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	log.Print("server started")
	return e.StartH2CServer(":8080", c)
}
