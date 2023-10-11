package services

import (
	"net/http"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/presenters.go"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	ItemService interface {
		Validate(model *models.ItemModel) *presenters.PresenterValidationErrors
		// Find() []*models.ItemModel
		FindOne(id uuid.UUID) (*models.ItemModel, error)
		Create(model *models.ItemModel) error
		// Update(id uuid.UUID, model *models.ItemModel) error
		// Delete(id uuid.UUID) error
	}

	itemService struct {
		validator *validator.Validate
		conn      *gorm.DB
	}
)

func NewItemService(validator *validator.Validate, conn *gorm.DB) ItemService {
	return itemService{
		validator,
		conn,
	}
}

func (is itemService) Validate(model *models.ItemModel) *presenters.PresenterValidationErrors {
	err := is.validator.Struct(model)

	if err != nil {
		var response presenters.PresenterValidationErrors
		response.Message = "Validation Errors"
		response.Status = http.StatusUnprocessableEntity
		response.Errors = []presenters.PresenterValidationErr{}

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

func (is itemService) FindOne(id uuid.UUID) (*models.ItemModel, error) {
	model := *&models.ItemModel{
		ID: id,
	}

	// is.conn.First(model)

	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	return &model, nil
}
