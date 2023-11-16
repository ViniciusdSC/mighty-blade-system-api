package services

import (
	"net/http"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	dbconnection "github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/db-connection"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/validation"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/presenters"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var itemType map[models.ItemTypeEnum]bool = map[models.ItemTypeEnum]bool{
	models.ItemTypeGeneral: true,
	models.ItemTypeWeapon:  true,
}

type (
	ItemFilter struct {
		presenters.PresenterPaginationInQuery
		Name *string              `form:"name"`
		Type *models.ItemTypeEnum `form:"type"`
	}

	ItemService interface {
		Validate(model *models.ItemModel, ignoreRules ...string) error
		Find(filter ItemFilter) ([]models.ItemModel, error)
		Count(filter ItemFilter) (*int64, error)
		FindOne(id string) (*models.ItemModel, error)
		Create(model *models.ItemModel) error
		Update(model *models.ItemModel, values *models.ItemModel) error
		Delete(model *models.ItemModel) error
	}

	itemService struct {
		validator validation.AppValidate
		conn      dbconnection.GormDB
	}
)

func NewItemService(validator validation.AppValidate, conn dbconnection.GormDB) ItemService {
	return itemService{
		validator,
		conn,
	}
}

func (is itemService) Validate(model *models.ItemModel, ignoreRules ...string) error {
	err := is.validator.StructExceptRules(model, ignoreRules...)

	if err != nil {
		response := presenters.ValidationError{
			Message: "Validation Errors",
			Status:  http.StatusUnprocessableEntity,
			Errors:  []presenters.PresenterValidationErr{},
		}

		for _, vErr := range err.(validator.ValidationErrors) {
			response.Errors = append(response.Errors, presenters.PresenterValidationErr{
				Tag: vErr.Tag(),
				Key: vErr.Field(),
			})
		}

		return response
	}

	if !itemType[model.Type] {
		return presenters.ValidationError{
			Message: "Validation Errors",
			Status:  http.StatusUnprocessableEntity,
			Errors: []presenters.PresenterValidationErr{
				{
					Tag: "invalid_type",
					Key: "type",
				},
			},
		}
	}

	return nil
}

func (is itemService) Create(model *models.ItemModel) error {
	result := is.conn.Create(model)

	return result.Error
}

func (is itemService) FindOne(id string) (*models.ItemModel, error) {
	var model models.ItemModel

	result := is.conn.Where("id = ?", id).First(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &model, nil
}

func (is itemService) Find(filter ItemFilter) ([]models.ItemModel, error) {
	var items []models.ItemModel

	result := is.conn.
		Scopes(is.filterScope(filter), PaginationScope(filter.Page, filter.PerPage)).
		Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (is itemService) Count(filter ItemFilter) (*int64, error) {
	var count int64

	result := is.conn.Model(&models.ItemModel{}).Scopes(is.filterScope(filter)).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	return &count, nil
}

func (is itemService) Update(model *models.ItemModel, values *models.ItemModel) error {
	result := is.conn.Model(&model).Updates(values)

	return result.Error
}

func (is itemService) Delete(model *models.ItemModel) error {
	result := is.conn.Delete(&model)

	return result.Error
}

func (is itemService) filterScope(filter ItemFilter) func(conn *gorm.DB) *gorm.DB {
	return func(conn *gorm.DB) *gorm.DB {
		tx := conn.Session(&gorm.Session{})
		if filter.Name != nil && len(*filter.Name) >= 3 {
			tx = conn.Where("name LIKE ?", "%"+*filter.Name+"%")
		}

		if filter.Type != nil {
			tx = conn.Where("type = ?", *filter.Type)
		}

		return tx
	}
}
