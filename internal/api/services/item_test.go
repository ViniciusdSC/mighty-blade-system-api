package services

import (
	"testing"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	mock_test "github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/validation/mock"
	gorm_test "github.com/ViniciusdSC/mighty-blade-system-api/internal/tests/gorm"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ItemServiceTestSuite struct {
	suite.Suite
	conn *gorm.DB
	v    *mock_test.AppValidateMock
}

func (suite *ItemServiceTestSuite) SetupTest() {
	suite.conn = gorm_test.CreateGormMock(suite.T())

	// suite.v = mock_test.NewAppValidateMock()
	suite.v = new(mock_test.AppValidateMock)
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
	model := new(models.ItemModel)
	suite.v.On("StructExceptRules", model).Return(nil)

}

func (suite *ItemServiceTestSuite) TestValidateWhenItemTypeIsNotValid() {

}

func (suite *ItemServiceTestSuite) TestValidateWhenValidatorRetunsAnError() {

}

func TestItemServiceSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceTestSuite))
}
