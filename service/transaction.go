package service

import (
	"OneLab2/dto"
	"OneLab2/models"
	"OneLab2/storage"
)

type Transaction struct {
	Repo *storage.Storage
}

func (s *Transaction) GetList() ([]dto.BookTransaction, error) {
	res, err := s.Repo.Transaction.GetList()
	ret := make([]dto.BookTransaction, 0)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	curBookTransaction := dto.BookTransaction{
		Book:        res[0].BookName,
		TotalAmount: 0,
	}
	for _, bookTransaction := range res {
		if curBookTransaction.Book == bookTransaction.BookName {
			curBookTransaction.TotalAmount += bookTransaction.Amount
		} else {
			ret = append(ret, curBookTransaction)
			curBookTransaction = dto.BookTransaction{
				Book:        bookTransaction.BookName,
				TotalAmount: bookTransaction.Amount,
			}
		}
	}
	ret = append(ret, curBookTransaction)
	return ret, nil
}

func (s *Transaction) Get(i int64) (*models.Transaction, error) {
	return s.Get(i)
}

func (s *Transaction) Create(t *models.Transaction) (*models.Transaction, error) {
	return s.Repo.Transaction.Create(t)
}

func NewTransaction(repo *storage.Storage) *Transaction {
	return &Transaction{Repo: repo}
}
