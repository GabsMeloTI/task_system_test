package test

import (
	"awesomeProject/db"
	"awesomeProject/dto/task_dto"
	"awesomeProject/models"
	service2 "awesomeProject/service"
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

func TestGetAllTasksService(t *testing.T) {
	mockDB := new(MockDB)
	service := service2.TaskService{}

	tasks := []models.Task{
		{Title: "Task1", Description: "Description1", Priority: "High", Status: "Open"},
		{Title: "Task2", Description: "Description2", Priority: "Medium", Status: "InProgress"},
	}
	tasksDTO := []task_dto.TaskListingDTO{
		{ID: 1, Title: "Task1", Description: "Description1", Priority: "High", Status: "Open"},
		{ID: 2, Title: "Task2", Description: "Description2", Priority: "Medium", Status: "InProgress"},
	}

	mockDB.On("Preload", "User").Return(mockDB)
	mockDB.On("Preload", "Section").Return(mockDB)
	mockDB.On("Preload", "Labels").Return(mockDB)
	mockDB.On("Find", &tasks).Return(mockDB)

	db.DB = mockDB
	result, err := service.GetAllTasks()
	assert.Nil(t, err)
	assert.ElementsMatch(t, tasksDTO, result)
}

func TestGetTaskByIDService(t *testing.T) {
	mockDB := new(MockDB)
	service := service2.TaskService{}

	task := models.Task{ID: 1, Title: "Task1", Description: "Description1", Priority: "High", Status: "Open"}
	taskDTO := task_dto.TaskListingDTO{ID: 1, Title: "Task1", Description: "Description1", Priority: "High", Status: "Open"}

	mockDB.On("Preload", "User").Return(mockDB)
	mockDB.On("Preload", "Section").Return(mockDB)
	mockDB.On("First", &task, uint(1)).Return(mockDB)

	db.DB = mockDB
	result, err := service.GetTaskByID(1)
	assert.Nil(t, err)
	assert.Equal(t, taskDTO, result)
}

func TestCreateTaskService(t *testing.T) {
	mockDB := new(MockDB)
	service := service2.TaskService{}
	task := models.Task{Title: "New Task", Description: "New Description", UserID: 1, SectionID: 1}

	mockDB.On("First", &models.User{}, task.UserID).Return(mockDB)
	mockDB.On("First", &models.Section{}, task.SectionID).Return(mockDB)
	mockDB.On("Create", &task).Return(mockDB)

	db.DB = mockDB
	err := service.CreateTask(task)
	assert.Nil(t, err)
}

func TestUpdateTaskService(t *testing.T) {
	mockDB := new(MockDB)
	service := service2.TaskService{}
	task := models.Task{Title: "Updated Task", Description: "Updated Description"}

	mockDB.On("First", &models.Task{}, uint(1)).Return(mockDB)
	mockDB.On("Save", &task).Return(mockDB)

	db.DB = mockDB
	err := service.UpdateTask(1, task)
	assert.Nil(t, err)
}

func TestDeleteTaskService(t *testing.T) {
	mockDB := new(MockDB)
	service := service2.TaskService{}

	mockDB.On("Delete", &models.Task{}, uint(1)).Return(mockDB)

	db.DB = mockDB
	err := service.DeleteTask(1)
	assert.Nil(t, err)
}

func TestAssignLabelsToTaskService(t *testing.T) {
	mockDB := new(MockDB)
	service := service2.TaskService{}
	taskID := uint(1)
	labels := []models.Label{{Name: "Label1", Color: "Red"}}

	task := models.Task{}
	mockDB.On("Preload", "Labels").First(&task, taskID).Return(mockDB)
	mockDB.On("Save", &task).Return(mockDB)

	db.DB = mockDB
	err := service.AssignLabelsToTask(taskID, labels)
	assert.Nil(t, err)
}
