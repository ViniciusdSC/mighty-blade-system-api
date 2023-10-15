package services

import (
	"fmt"
	"net/http"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/validation"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/presenters"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type (
	ItemFilter struct {
		presenters.PresenterPaginationInQuery
		Name *string              `form:"name"`
		Type *models.ItemTypeEnum `form:"type"`
	}

	ItemService interface {
		Validate(model *models.ItemModel, ignoreRules ...string) *presenters.PresenterValidationErrors
		Find(filter ItemFilter) ([]models.ItemModel, error)
		Count(filter ItemFilter) (*int64, error)
		FindOne(id string) (*models.ItemModel, error)
		Create(model *models.ItemModel) error
		// Update(id uuid.UUID, model *models.ItemModel) error
		// Delete(id uuid.UUID) error
	}

	itemService struct {
		validator *validation.AppValidate
		conn      *gorm.DB
	}
)

func NewItemService(validator *validation.AppValidate, conn *gorm.DB) ItemService {
	return itemService{
		validator,
		conn,
	}
}

func (is itemService) Validate(model *models.ItemModel, ignoreRules ...string) *presenters.PresenterValidationErrors {
	err := is.validator.StructExceptRules(model, ignoreRules...)

	if err != nil {
		response := presenters.PresenterValidationErrors{
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

		return &response
	}

	if !model.ItemTypeIsValid() {
		return &presenters.PresenterValidationErrors{
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

	result := is.conn.First(&model).Where("id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &model, nil
}

func (is itemService) Find(filter ItemFilter) ([]models.ItemModel, error) {
	var items []models.ItemModel

	result := is.conn.
		Scopes(is.filterScope(filter), PaginationScope(filter.Page, filter.PerPage)).
		Find(items)

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

func (is itemService) filterScope(filter ItemFilter) func(conn *gorm.DB) *gorm.DB {
	return func(conn *gorm.DB) *gorm.DB {
		if filter.Name != nil && len(*filter.Name) > 3 {
			conn = conn.Where("name LIKE ?", filter.Name)
		}

		fmt.Println("type", filter.Type)

		if filter.Type != nil {
			conn = conn.Where("type = ?", filter.Type)
		}

		return conn
	}
}
