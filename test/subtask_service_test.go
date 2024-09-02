package test

import (
	"awesomeProject/db"
	"awesomeProject/dto/subtask_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

// MockDB simula um banco de dados para os testes
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Preload(query string) *gorm.DB {
	args := m.Called(query)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(value, where)
	return args.Get(0).(*gorm.DB)
}

func TestGetAllSubtasksService(t *testing.T) {
	mockDB := new(MockDB)
	service := SubtaskService{}

	subtasks := []models.Subtask{
		{ID: 1, Title: "Subtask1", Description: "Description1", Status: "Open", Task: models.Task{ID: 1, Title: "Task1", Priority: "High"}},
		{ID: 2, Title: "Subtask2", Description: "Description2", Status: "InProgress", Task: models.Task{ID: 2, Title: "Task2", Priority: "Medium"}},
	}
	subtasksDTO := []subtask_dto.SubtaskListingDTO{
		{ID: 1, Title: "Subtask1", Description: "Description1", CreatedAt: subtasks[0].CreatedAt, Status: "Open", Task: task_dto.TaskBasicDTO{ID: 1, Title: "Task1", Priority: "High"}},
		{ID: 2, Title: "Subtask2", Description: "Description2", CreatedAt: subtasks[1].CreatedAt, Status: "InProgress", Task: task_dto.TaskBasicDTO{ID: 2, Title: "Task2", Priority: "Medium"}},
	}

	mockDB.On("Preload", "Task").Return(mockDB)
	mockDB.On("Find", &subtasks).Return(mockDB)

	db.DB = mockDB
	result, err := service.GetAllSubtasks()
	assert.Nil(t, err)
	assert.ElementsMatch(t, subtasksDTO, result)
}

func TestGetSubtaskByIDService(t *testing.T) {
	mockDB := new(MockDB)
	service := SubtaskService{}

	subtask := models.Subtask{ID: 1, Title: "Subtask1", Description: "Description1", Status: "Open", Task: models.Task{ID: 1, Title: "Task1", Priority: "High"}}
	subtaskDTO := subtask_dto.SubtaskListingDTO{ID: 1, Title: "Subtask1", Description: "Description1", CreatedAt: subtask.CreatedAt, Status: "Open", Task: task_dto.TaskBasicDTO{ID: 1, Title: "Task1", Priority: "High"}}

	mockDB.On("Preload", "Task").Return(mockDB)
	mockDB.On("First", &subtask, uint(1)).Return(mockDB)

	db.DB = mockDB
	result, err := service.GetSubtaskByID(1)
	assert.Nil(t, err)
	assert.Equal(t, subtaskDTO, result)
}

func TestCreateSubtaskService(t *testing.T) {
	mockDB := new(MockDB)
	service := SubtaskService{}
	subtask := models.Subtask{Title: "New Subtask", Description: "New Description", TaskID: 1}

	mockDB.On("Create", &subtask).Return(mockDB)

	db.DB = mockDB
	err := service.CreateSubtask(subtask)
	assert.Nil(t, err)
}

func TestUpdateSubtaskService(t *testing.T) {
	mockDB := new(MockDB)
	service := SubtaskService{}
	subtask := models.Subtask{ID: 1, Title: "Updated Subtask", Description: "Updated Description"}

	mockDB.On("First", &models.Subtask{}, uint(1)).Return(mockDB)
	mockDB.On("Save", &subtask).Return(mockDB)

	db.DB = mockDB
	err := service.UpdateSubtask(1, subtask)
	assert.Nil(t, err)
}

func TestDeleteSubtaskService(t *testing.T) {
	mockDB := new(MockDB)
	service := SubtaskService{}

	mockDB.On("Delete", &models.Subtask{}, uint(1)).Return(mockDB)

	db.DB = mockDB
	err := service.DeleteSubtask(1)
	assert.Nil(t, err)
}
