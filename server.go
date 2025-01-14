package echo_ifElse

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"log"
	"time"
)

func Run() error {
	e := echo.New()
	e.Use()
	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	log.Print("server started")
	return e.StartH2CServer(":8080", s)
}
