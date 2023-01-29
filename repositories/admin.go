package repositories

import (
	"context"

	"github.com/adityarizkyramadhan/freepass-2023/domain"
	"gorm.io/gorm"
)

type (
	adminRepositories struct {
		db *gorm.DB
	}
	AdminRepositories interface {
		Register(ctx context.Context, input *domain.Admin) error
		SearchAdminByUsername(ctx context.Context, username string) (*domain.Admin, error)
	}
)

// Constructor
func NewAdmin(db *gorm.DB) AdminRepositories {
	return &adminRepositories{db: db}
}

func (thisAdmin *adminRepositories) Register(ctx context.Context, input *domain.Admin) error {
	return thisAdmin.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(input).Error
	})
}

func (thisAdmin *adminRepositories) SearchAdminByUsername(ctx context.Context, username string) (*domain.Admin, error) {
	admin := new(domain.Admin)
	err := thisAdmin.db.WithContext(ctx).Where("username = ?", username).First(admin).Error
	return admin, err
}
