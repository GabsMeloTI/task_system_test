package test

import (
	"awesomeProject/controllers"
	"awesomeProject/dto/task_dto"
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockTaskService simula o serviço de tarefas para os testes
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetAllTasks() ([]task_dto.TaskListingDTO, error) {
	args := m.Called()
	return args.Get(0).([]task_dto.TaskListingDTO), args.Error(1)
}

func (m *MockTaskService) GetTaskByID(id uint) (task_dto.TaskListingDTO, error) {
	args := m.Called(id)
	return args.Get(0).(task_dto.TaskListingDTO), args.Error(1)
}

func (m *MockTaskService) CreateTask(task models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskService) UpdateTask(id uint, task models.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskService) DeleteTask(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskService) AssignLabelsToTask(taskID uint, labels []models.Label) error {
	args := m.Called(taskID, labels)
	return args.Error(0)
}

func TestGetTasksController(t *testing.T) {
	mockService := new(MockTaskService)
	controller := &controllers.TaskController{Service: mockService}

	tasksDTO := []task_dto.TaskListingDTO{
		{ID: 1, Title: "Task1", Description: "Description1", Priority: "High", Status: "Open"},
		{ID: 2, Title: "Task2", Description: "Description2", Priority: "Medium", Status: "InProgress"},
	}

	mockService.On("GetAllTasks").Return(tasksDTO, nil)

	req := httptest.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetTasks)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response []task_dto.TaskListingDTO
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.ElementsMatch(t, tasksDTO, response)
}

func TestGetTaskByIDController(t *testing.T) {
	mockService := new(MockTaskService)
	controller := &controllers.TaskController{Service: mockService}

	taskDTO := task_dto.TaskListingDTO{
		ID:          1,
		Title:       "Task1",
		Description: "Description1",
		Priority:    "High",
		Status:      "Open",
	}

	mockService.On("GetTaskByID", uint(1)).Return(taskDTO, nil)

	req := httptest.NewRequest("GET", "/tasks/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetTaskByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response task_dto.TaskListingDTO
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, taskDTO, response)
}

func TestCreateTaskController(t *testing.T) {
	mockService := new(MockTaskService)
	controller := &controllers.TaskController{Service: mockService}
	task := models.Task{Title: "New Task", Description: "New Description", UserID: 1, SectionID: 1}

	body, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateTask)

	mockService.On("CreateTask", task).Return(nil)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateTaskController(t *testing.T) {
	mockService := new(MockTaskService)
	controller := &controllers.TaskController{Service: mockService}
	task := models.Task{Title: "Updated Task", Description: "Updated Description"}

	body, _ := json.Marshal(task)
	req := httptest.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateTask)

	mockService.On("UpdateTask", uint(1), task).Return(nil)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteTaskController(t *testing.T) {
	mockService := new(MockTaskService)
	controller := &controllers.TaskController{Service: mockService}

	req := httptest.NewRequest("DELETE", "/tasks/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteTask)

	mockService.On("DeleteTask", uint(1)).Return(nil)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestAssignLabelsToTaskController(t *testing.T) {
	mockService := new(MockTaskService)
	controller := &controllers.TaskController{Service: mockService}

	labels := []models.Label{{Name: "Label1", Color: "Red"}}
	body, _ := json.Marshal(labels)
	req := httptest.NewRequest("POST", "/tasks/1/labels", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.AssignLabelsToTask)

	mockService.On("AssignLabelsToTask", uint(1), labels).Return(nil)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
