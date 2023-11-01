package mock_test

import (
	"github.com/stretchr/testify/mock"
)

type AppValidateMock struct {
	mock.Mock
}

func (avm *AppValidateMock) StructExceptRules(s interface{}, rules ...string) error {
	args := avm.Called(s, rules)

	return args.Error(1)
}
