package postgres

import (
	"OneLab2/models"
	"gorm.io/gorm"
)

type Transaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{db: db}
}

func (r *Transaction) GetList() ([]*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Transaction) Get(i int64) (*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Transaction) Create(t *models.Transaction) (*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
