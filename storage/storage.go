package storage

import (
	"OneLab2/models"
	"OneLab2/storage/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Transaction ITransaction
}

func NewStorage(db *gorm.DB) (*Storage, error) {
	return &Storage{Transaction: postgres.NewTransaction(db)}, nil
}

type ITransaction interface {
	GetList() ([]*models.Transaction, error)
	Get(int64) (*models.Transaction, error)
	Create(transaction *models.Transaction) (*models.Transaction, error)
}
