package test

import (
	"awesomeProject/models"
	_ "awesomeProject/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id string) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) UpdateUser(id string, user *models.User) error {
	args := m.Called(id, user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) RegisterUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUserService(t *testing.T) {
	mockService := new(MockUserService)
	user := &models.User{Name: "New User", Email: "newuser@example.com", Password: "password"}
	mockService.On("CreateUser", user).Return(nil)

	err := mockService.CreateUser(user)
	assert.Nil(t, err)
}

func TestGetUsersService(t *testing.T) {
	mockService := new(MockUserService)
	users := []models.User{
		{Name: "User1", Email: "user1@example.com"},
		{Name: "User2", Email: "user2@example.com"},
	}
	mockService.On("GetUsers").Return(users, nil)

	result, err := mockService.GetUsers()
	assert.Nil(t, err)
	assert.Equal(t, users, result)
}

func TestCreateUserServiceError(t *testing.T) {
	mockService := new(MockUserService)
	user := &models.User{Name: "New User", Email: "newuser@example.com", Password: "password"}
	mockService.On("CreateUser", user).Return(errors.New("error creating user"))

	err := mockService.CreateUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, "error creating user", err.Error())
}
