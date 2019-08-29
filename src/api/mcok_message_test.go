package api_test

import (
	"greeting/model"

	"github.com/stretchr/testify/mock"
)

type mockGreeting struct {
	mock.Mock
}

func (greeting *mockGreeting) GetMessage() (model.Greeting, error) {
	argument := greeting.Called()
	return argument.Get(0).(model.Greeting), argument.Error(1)
}
