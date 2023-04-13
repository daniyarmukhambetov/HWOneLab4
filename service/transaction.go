package service

import (
	"OneLab2/models"
	"OneLab2/storage"
)

type Transaction struct {
	Repo *storage.Storage
}

func (s *Transaction) GetList() ([]*models.Transaction, error) {
	return s.Repo.Transaction.GetList()
}

func (s *Transaction) Get(i int64) (*models.Transaction, error) {
	return s.Get(i)
}

func (s *Transaction) Create(t *models.Transaction) (*models.Transaction, error) {
	return s.Create(t)
}

func NewTransaction(repo *storage.Storage) *Transaction {
	return &Transaction{Repo: repo}
}
