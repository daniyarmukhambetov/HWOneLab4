package service

import (
	"OneLab2/models"
	"OneLab2/storage"
)

type Service struct {
	Transaction ITransaction
}

func NewService(repo *storage.Storage) (*Service, error) {
	return &Service{Transaction: NewTransaction(repo)}, nil
}

type ITransaction interface {
	GetList() ([]*models.Transaction, error)
	Get(int64) (*models.Transaction, error)
	Create(*models.Transaction) (*models.Transaction, error)
}
