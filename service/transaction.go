package service

import (
	"OneLab2/models"
	"OneLab2/storage"
)

type Transaction struct {
	Repo *storage.Storage
}

func (s *Transaction) GetList() ([]*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Transaction) Get(i int64) (*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Transaction) Create(t *models.Transaction) (*models.Transaction, error) {
	panic("implement me")
}

func NewTransaction(repo *storage.Storage) *Transaction {
	return &Transaction{Repo: repo}
}
