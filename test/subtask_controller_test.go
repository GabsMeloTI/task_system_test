package test

import (
	"awesomeProject/dto/subtask_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// MockSubtaskService simula o serviço de subtarefas para os testes
type MockSubtaskService struct {
	mock.Mock
}

func (m *MockSubtaskService) GetAllSubtasks() ([]subtask_dto.SubtaskListingDTO, error) {
	args := m.Called()
	return args.Get(0).([]subtask_dto.SubtaskListingDTO), args.Error(1)
}

func (m *MockSubtaskService) GetSubtaskByID(id uint) (subtask_dto.SubtaskListingDTO, error) {
	args := m.Called(id)
	return args.Get(0).(subtask_dto.SubtaskListingDTO), args.Error(1)
}

func (m *MockSubtaskService) CreateSubtask(subtask models.Subtask) error {
	args := m.Called(subtask)
	return args.Error(0)
}

func (m *MockSubtaskService) UpdateSubtask(id uint, subtask models.Subtask) error {
	args := m.Called(id, subtask)
	return args.Error(0)
}

func (m *MockSubtaskService) DeleteSubtask(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetSubtasksController(t *testing.T) {
	mockService := new(MockSubtaskService)
	controller := &SubtaskController{Service: mockService}

	subtasksDTO := []subtask_dto.SubtaskListingDTO{
		{ID: 1, Title: "Subtask1", Description: "Description1", CreatedAt: time.Now(), Status: "Open", Task: task_dto.TaskBasicDTO{ID: 1, Title: "Task1", Priority: "High"}},
		{ID: 2, Title: "Subtask2", Description: "Description2", CreatedAt: time.Now(), Status: "InProgress", Task: task_dto.TaskBasicDTO{ID: 2, Title: "Task2", Priority: "Medium"}},
	}

	mockService.On("GetAllSubtasks").Return(subtasksDTO, nil)

	req := httptest.NewRequest("GET", "/subtasks", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetSubtasks)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response []subtask_dto.SubtaskListingDTO
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.ElementsMatch(t, subtasksDTO, response)
}

func TestGetSubtaskByIDController(t *testing.T) {
	mockService := new(MockSubtaskService)
	controller := &SubtaskController{Service: mockService}

	subtaskDTO := subtask_dto.SubtaskListingDTO{
		ID:          1,
		Title:       "Subtask1",
		Description: "Description1",
		CreatedAt:   time.Now(),
		Status:      "Open",
		Task: task_dto.TaskBasicDTO{
			ID:       1,
			Title:    "Task1",
			Priority: "High",
		},
	}

	mockService.On("GetSubtaskByID", uint(1)).Return(subtaskDTO, nil)

	req := httptest.NewRequest("GET", "/subtasks/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetSubtaskByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response subtask_dto.SubtaskListingDTO
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, subtaskDTO, response)
}

func TestCreateSubtaskController(t *testing.T) {
	mockService := new(MockSubtaskService)
	controller := &SubtaskController{Service: mockService}

	subtask := models.Subtask{Title: "New Subtask", Description: "New Description"}

	body, _ := json.Marshal(subtask)
	req := httptest.NewRequest("POST", "/subtasks", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateSubtask)

	mockService.On("CreateSubtask", subtask).Return(nil)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateSubtaskController(t *testing.T) {
	mockService := new(MockSubtaskService)
	controller := &SubtaskController{Service: mockService}
	subtask := models.Subtask{Title: "Updated Subtask", Description: "Updated Description"}

	body, _ := json.Marshal(subtask)
	req := httptest.NewRequest("PUT", "/subtasks/1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateSubtask)

	mockService.On("UpdateSubtask", uint(1), subtask).Return(nil)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteSubtaskController(t *testing.T) {
	mockService := new(MockSubtaskService)
	controller := &SubtaskController{Service: mockService}

	req := httptest.NewRequest("DELETE", "/subtasks/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteSubtask)

	mockService.On("DeleteSubtask", uint(1)).Return(nil)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
