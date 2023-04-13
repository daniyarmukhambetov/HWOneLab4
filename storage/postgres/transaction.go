package postgres

import (
	"OneLab2/models"
	"fmt"
	"gorm.io/gorm"
)

type Transaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{db: db}
}

func (r *Transaction) GetList() ([]*models.Transaction, error) {
	sql := "SELECT * FROM transactions"
	var tr []*models.Transaction
	tx := r.db.Raw(sql).Scan(&tr)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tr, nil
}

func (r *Transaction) Get(i int64) (*models.Transaction, error) {
	sql := fmt.Sprintf("SELECT * FROM transactions WHERE id = %d", i)
	var tr *models.Transaction
	err := r.db.Raw(sql).Scan(&tr).Error
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (r *Transaction) Create(t *models.Transaction) (*models.Transaction, error) {
	sql := fmt.Sprintf(
		"INSERT INTO transactions (user_id, amount) VALUES (%d, %v)",
		t.UserID,
		t.Amount,
	)
	var tr *models.Transaction
	err := r.db.Raw(sql).Scan(&tr).Error
	if err != nil {
		return nil, err
	}
	return tr, nil
}
