package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) FindByID(id int) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func TestGetUserName_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	service := NewService(mockRepo)
	mockUser := &User{ID: 1, Name: "Asif"}
	mockRepo.On("FindByID", 1).Return(mockUser, nil)

	name, err := service.GetUserName(1)

	assert.Nil(t, err)
	assert.Equal(t, "Asif", name)
	mockRepo.AssertExpectations(t)
}

func TestGetUserName_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	service := NewService(mockRepo)
	errNotFound := errors.New("user not found")
	mockRepo.On("FindByID", 2).Return(nil, errNotFound)

	name, err := service.GetUserName(2)

	assert.Equal(t, "", name)
	assert.EqualError(t, err, "user not found")
	mockRepo.AssertExpectations(t)
}
