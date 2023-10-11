package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemTypeEnum string

const (
	ItemTypeGeneral ItemTypeEnum = "general"
	ItemTypeWeapon  ItemTypeEnum = "weapon"
)

var itemType map[ItemTypeEnum]bool = map[ItemTypeEnum]bool{
	ItemTypeGeneral: true,
	ItemTypeWeapon:  true,
}

type ItemModel struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string       `json:"name" validate:"required"`
	Price       int          `json:"price" validate:"required"`
	Weight      float64      `json:"weight" validate:"required"`
	Description string       `json:"description" validate:"required"`
	Type        ItemTypeEnum `json:"type" validate:"required"`
}

func (im ItemModel) TableName() string {
	return "item"
}

func (im ItemModel) ItemTypeIsValid() bool {
	return itemType[im.Type]
}
