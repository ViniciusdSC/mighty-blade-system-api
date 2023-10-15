package services

import (
	"gorm.io/gorm"
)

func PaginationScope(page *int, perPage *int) func(conn *gorm.DB) *gorm.DB {
	if page == nil {
		*page = 1
	}

	if perPage == nil {
		*perPage = 10
	}

	return func(conn *gorm.DB) *gorm.DB {
		return conn.Limit(*perPage).Offset(*perPage * (*page - 1))
	}
}
