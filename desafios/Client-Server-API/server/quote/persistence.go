package quote

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type persistence struct {
	db *gorm.DB
}

func newPersistence() (*persistence, error) {
	dsn := "file:./quote.db?cache=shared&mode=rwc"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&CoinEntity{}); err != nil {
		return nil, err
	}
	return &persistence{db: db}, nil
}

func (p *persistence) create(ctx context.Context, coin *CoinEntity) error {
	return p.db.WithContext(ctx).Create(coin).Error
}

func (p *persistence) close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get db from gorm: %w", err)
	}
	return sqlDB.Close()
}
