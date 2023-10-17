package services

import (
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/validation"
	"gorm.io/gorm"
)

type AppServices struct {
	IS ItemService
}

func NewAppServices(v *validation.AppValidate, conn *gorm.DB) *AppServices {
	itemService := NewItemService(v, conn)

	return &AppServices{
		IS: itemService,
	}
}

func PaginationScope(page *int, perPage *int) func(conn *gorm.DB) *gorm.DB {
	return func(conn *gorm.DB) *gorm.DB {
		return conn.Limit(*perPage).Offset(*perPage * (*page - 1))
	}
}
