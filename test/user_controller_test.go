package test

import (
	"awesomeProject/controllers"
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the UserService
type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) GetUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserController) GetUserByID(id string) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserController) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserController) UpdateUser(id string, user *models.User) error {
	args := m.Called(id, user)
	return args.Error(0)
}

func (m *MockUserController) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserController) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserController) RegisterUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUserController(t *testing.T) {
	mockController := new(MockUserController)
	user := &models.User{Name: "New User", Email: "newuser@example.com", Password: "password"}
	mockController.On("CreateUser", user).Return(nil)

	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetUsersController(t *testing.T) {
	mockController := new(MockUserController)
	users := []models.User{
		{Name: "User1", Email: "user1@example.com"},
		{Name: "User2", Email: "user2@example.com"},
	}
	mockController.On("GetUsers").Return(users, nil)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
