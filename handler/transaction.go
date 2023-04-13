package handler

import (
	"OneLab2/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Transaction struct {
	service *service.Service
}

func NewTransaction(service *service.Service) *Transaction {
	return &Transaction{service: service}
}

func (h *Transaction) GetList(ctx echo.Context) error {
	transactions, err := h.service.Transaction.GetList()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, transactions)
}

func (h *Transaction) Create(ctx echo.Context) error {
	panic("implement me")
}

func (h *Transaction) Get(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
