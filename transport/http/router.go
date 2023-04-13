package http

import (
	"OneLab2/handler"
	"github.com/labstack/echo/v4"
)

func NewRouter(h *handler.Handler) *echo.Echo {
	r := echo.New()
	r.GET("/transactions", h.Transaction.GetList)
	r.GET("/transactions/{id}", h.Transaction.Get)
	r.POST("/transactions", h.Transaction.Create)
	return r
}
