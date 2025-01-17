package handler

import (
	"echo_ifElse/pkg/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(e *echo.Echo) {
	e.POST("/registration", h.registration)
	e.GET("/accounts/:id", h.getAcc)
}
