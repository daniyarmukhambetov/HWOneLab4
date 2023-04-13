package postgres

import (
	"OneLab2/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PgURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
