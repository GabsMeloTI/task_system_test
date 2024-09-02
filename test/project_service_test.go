package test

import (
	"awesomeProject/dto/project_dto"
	"awesomeProject/models"
	"awesomeProject/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

// Define a MockDB para simular um banco de dados
type MockDB struct {
	mock.Mock
}

// Retorna um *gorm.DB simulado para métodos chaináveis
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

func (m *MockDB) Model(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Updates(values interface{}) *gorm.DB {
	args := m.Called(values)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(value, where)
	return args.Get(0).(*gorm.DB)
}

// Ajuste a criação de uma instância do MockDB para ser compatível com os métodos do service
func TestGetProjectsService(t *testing.T) {
	mockDB := new(MockDB)
	projectService := service.NewProjectService(mockDB)

	projects := []models.Project{
		{Title: "Project1", Description: "Description1"},
		{Title: "Project2", Description: "Description2"},
	}
	projectDTOs := []project_dto.ProjectListingDTO{
		{ID: 1, Title: "Project1", Description: "Description1"},
		{ID: 2, Title: "Project2", Description: "Description2"},
	}

	// Configurar o MockDB para simular o comportamento esperado
	mockDB.On("Preload", "User").Return(mockDB)
	mockDB.On("Find", &projects).Return(mockDB)
	mockDB.On("Preload", "User").Return(mockDB)

	result, err := projectService.GetProjects()
	assert.Nil(t, err)
	assert.ElementsMatch(t, projectDTOs, result)
}

func TestCreateProjectService(t *testing.T) {
	mockDB := new(MockDB)
	projectService := service.NewProjectService(mockDB)
	project := models.Project{Title: "New Project", Description: "Description"}

	// Configurar o MockDB para simular o comportamento esperado
	mockDB.On("Create", &project).Return(mockDB)

	err := projectService.CreateProject(project)
	assert.Nil(t, err)
}

func TestUpdateProjectService(t *testing.T) {
	mockDB := new(MockDB)
	projectService := service.NewProjectService(mockDB)
	project := models.Project{Title: "Updated Project", Description: "Updated Description"}

	// Configurar o MockDB para simular o comportamento esperado
	mockDB.On("First", &models.Project{}, uint(1)).Return(mockDB)
	mockDB.On("Model", &models.Project{}).Return(mockDB)
	mockDB.On("Updates", project).Return(mockDB)

	err := projectService.UpdateProject(1, project)
	assert.Nil(t, err)
}

func TestDeleteProjectService(t *testing.T) {
	mockDB := new(MockDB)
	projectService := service.NewProjectService(new(MockDB))

	// Configurar o MockDB para simular o comportamento esperado
	mockDB.On("Delete", &models.Project{}, uint(1)).Return(mockDB)

	err := projectService.DeleteProject(1)
	assert.Nil(t, err)
}
