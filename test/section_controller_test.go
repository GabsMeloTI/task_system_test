package test

import (
	"awesomeProject/controllers"
	"awesomeProject/dto/section_dto"
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockSectionService struct {
	mock.Mock
}

func (m *MockSectionService) GetSections() ([]section_dto.SectionListingDTO, error) {
	args := m.Called()
	return args.Get(0).([]section_dto.SectionListingDTO), args.Error(1)
}

func (m *MockSectionService) GetSectionByID(id uint) (section_dto.SectionListingDTO, error) {
	args := m.Called(id)
	return args.Get(0).(section_dto.SectionListingDTO), args.Error(1)
}

func (m *MockSectionService) CreateSection(section models.Section) error {
	args := m.Called(section)
	return args.Error(0)
}

func (m *MockSectionService) UpdateSection(id uint, section models.Section) error {
	args := m.Called(id, section)
	return args.Error(0)
}

func (m *MockSectionService) DeleteSection(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetSectionsController(t *testing.T) {
	mockService := new(MockSectionService)
	controller := &controllers.SectionController{Service: mockService}

	sectionsDTO := []section_dto.SectionListingDTO{
		{ID: 1, Title: "Section1", Description: "Description1"},
		{ID: 2, Title: "Section2", Description: "Description2"},
	}

	mockService.On("GetSections").Return(sectionsDTO, nil)

	req := httptest.NewRequest("GET", "/sections", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetSections)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response []section_dto.SectionListingDTO
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.ElementsMatch(t, sectionsDTO, response)
}

func TestGetSectionByIDController(t *testing.T) {
	mockService := new(MockSectionService)
	controller := &controllers.SectionController{Service: mockService}

	sectionDTO := section_dto.SectionListingDTO{
		ID:          1,
		Title:       "Section1",
		Description: "Description1",
	}

	mockService.On("GetSectionByID", uint(1)).Return(sectionDTO, nil)

	req := httptest.NewRequest("GET", "/sections/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetSectionByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response section_dto.SectionListingDTO
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, sectionDTO, response)
}

func TestCreateSectionController(t *testing.T) {
	mockService := new(MockSectionService)
	controller := &controllers.SectionController{Service: mockService}
	section := models.Section{Title: "New Section", Description: "Description"}

	body, _ := json.Marshal(section)
	req := httptest.NewRequest("POST", "/sections", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateSection)

	mockService.On("CreateSection", section).Return(nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateSectionController(t *testing.T) {
	mockService := new(MockSectionService)
	controller := &controllers.SectionController{Service: mockService}
	section := models.Section{Title: "Updated Section", Description: "Updated Description"}

	body, _ := json.Marshal(section)
	req := httptest.NewRequest("PUT", "/sections/1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateSection)

	mockService.On("UpdateSection", uint(1), section).Return(nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteSectionController(t *testing.T) {
	mockService := new(MockSectionService)
	controller := &controllers.SectionController{Service: mockService}

	req := httptest.NewRequest("DELETE", "/sections/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteSection)

	mockService.On("DeleteSection", uint(1)).Return(nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
