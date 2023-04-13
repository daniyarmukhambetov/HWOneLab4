package handler

import (
	"OneLab2/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Transaction ITransaction
}

func NewHandler(srv *service.Service) (*Handler, error) {
	return &Handler{Transaction: NewTransaction(srv)}, nil
}

type ITransaction interface {
	GetList(ctx echo.Context) error
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
}
