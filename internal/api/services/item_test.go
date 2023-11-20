package services

import (
	"net/http"
	"testing"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	dbcMocks "github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/db-connection/mocks"
	vMocks "github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/validation/mocks"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/presenters"
	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/suite"
)

type ItemServiceTestSuite struct {
	suite.Suite
	conn *dbcMocks.MockGormDB
	v    *vMocks.MockAppValidate
}

type mockFieldError struct {
	validator.FieldError
	tag   string
	field string
}

func (mfe mockFieldError) Tag() string {
	return mfe.tag
}

func (mfe mockFieldError) Field() string {
	return mfe.field
}

func (suite *ItemServiceTestSuite) SetupTest() {
	suite.conn = new(dbcMocks.MockGormDB)

	suite.v = new(vMocks.MockAppValidate)
}

func (suite *ItemServiceTestSuite) TestNewItemService() {
	is := NewItemService(suite.v, suite.conn)

	suite.Assert().Equal(is, itemService{
		validator: suite.v,
		conn:      suite.conn,
	})
}

// Tests for Validate func
func (suite *ItemServiceTestSuite) TestValidateReturnNil() {
	var model *models.ItemModel
	model = &models.ItemModel{
		Type: models.ItemTypeGeneral,
	}

	suite.v.On("StructExceptRules", model).Return(nil)

	is := NewItemService(suite.v, suite.conn)

	err := is.Validate(model)

	suite.Assert().Nil(err)
}

func (suite *ItemServiceTestSuite) TestValidateWhenItemTypeIsNotValid() {
	var model *models.ItemModel
	model = &models.ItemModel{
		Type: "test",
	}

	suite.v.On("StructExceptRules", model).Return(nil)

	is := NewItemService(suite.v, suite.conn)

	err := is.Validate(model)

	expectedErr := &presenters.ValidationError{
		Message: "Validation Errors",
		Status:  http.StatusUnprocessableEntity,
		Errors: []presenters.PresenterValidationErr{
			{
				Tag: "invalid_type",
				Key: "type",
			},
		},
	}

	suite.Require().Error(err, "Should not be nil")
	suite.Require().ErrorAs(err, expectedErr, "Should be validation error")

	vErr := err.(presenters.ValidationError)

	suite.Assert().Equal(expectedErr.Message, vErr.Message)
	suite.Assert().Equal(expectedErr.Status, vErr.Status)
	for i := range expectedErr.Errors {
		suite.Assert().Equal(expectedErr.Errors[i].Key, vErr.Errors[i].Key)
		suite.Assert().Equal(expectedErr.Errors[i].Tag, vErr.Errors[i].Tag)
	}
}

func (suite *ItemServiceTestSuite) TestValidateWhenValidatorRetunsAnError() {
	var model *models.ItemModel
	model = &models.ItemModel{
		Type: models.ItemTypeGeneral,
	}

	suite.v.On("StructExceptRules", model).Return(validator.ValidationErrors{
		&mockFieldError{tag: "required", field: "name"},
	})

	is := NewItemService(suite.v, suite.conn)

	err := is.Validate(model)

	expectedErr := &presenters.ValidationError{
		Message: "Validation Errors",
		Status:  http.StatusUnprocessableEntity,
		Errors: []presenters.PresenterValidationErr{
			{
				Tag: "required",
				Key: "name",
			},
		},
	}

	suite.Require().Error(err, "Should not be nil")
	suite.Require().ErrorAs(err, expectedErr, "Should be validation error")

	vErr := err.(presenters.ValidationError)

	suite.Assert().Equal(expectedErr.Message, vErr.Message)
	suite.Assert().Equal(expectedErr.Status, vErr.Status)
	for i := range expectedErr.Errors {
		suite.Assert().Equal(expectedErr.Errors[i].Key, vErr.Errors[i].Key)
		suite.Assert().Equal(expectedErr.Errors[i].Tag, vErr.Errors[i].Tag)
	}
}

func TestItemServiceSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceTestSuite))
}
