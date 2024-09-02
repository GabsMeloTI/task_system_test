package test

import (
	"awesomeProject/controllers"
	"awesomeProject/dto/project_dto"
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mocking the ProjectService
type MockProjectService struct {
	mock.Mock
}

func (m *MockProjectService) GetProjects() ([]project_dto.ProjectListingDTO, error) {
	args := m.Called()
	return args.Get(0).([]project_dto.ProjectListingDTO), args.Error(1)
}

func (m *MockProjectService) GetProjectByID(id uint) (project_dto.ProjectListingDTO, error) {
	args := m.Called(id)
	return args.Get(0).(project_dto.ProjectListingDTO), args.Error(1)
}

func (m *MockProjectService) CreateProject(project models.Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockProjectService) UpdateProject(id uint, project models.Project) error {
	args := m.Called(id, project)
	return args.Error(0)
}

func (m *MockProjectService) DeleteProject(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetProjectsController(t *testing.T) {
	mockService := new(MockProjectService)
	projectsDTO := []project_dto.ProjectListingDTO{
		{ID: 1, Title: "Project1", Description: "Description1"},
		{ID: 2, Title: "Project2", Description: "Description2"},
	}
	mockService.On("GetProjects").Return(projectsDTO, nil)

	controller := &controllers.ProjectController{Service: mockService}
	req, err := http.NewRequest("GET", "/projects", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetProjects)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var projects []project_dto.ProjectListingDTO
	err = json.NewDecoder(rr.Body).Decode(&projects)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, projectsDTO, projects)
}

func TestCreateProjectController(t *testing.T) {
	mockService := new(MockProjectService)
	project := models.Project{Title: "New Project", Description: "Description"}
	mockService.On("CreateProject", project).Return(nil)

	controller := &controllers.ProjectController{Service: mockService}
	body, _ := json.Marshal(project)
	req, err := http.NewRequest("POST", "/projects", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateProject)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateProjectController(t *testing.T) {
	mockService := new(MockProjectService)
	project := models.Project{Title: "Updated Project", Description: "Updated Description"}
	mockService.On("UpdateProject", uint(1), project).Return(nil)

	controller := &controllers.ProjectController{Service: mockService}
	body, _ := json.Marshal(project)
	req, err := http.NewRequest("PUT", "/projects/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateProject)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteProjectController(t *testing.T) {
	mockService := new(MockProjectService)
	mockService.On("DeleteProject", uint(1)).Return(nil)

	controller := &controllers.ProjectController{Service: mockService}
	req, err := http.NewRequest("DELETE", "/projects/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteProject)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
