package test

import (
	"awesomeProject/dto/section_dto"
	"awesomeProject/models"
	"awesomeProject/service"
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

func TestGetSectionsService(t *testing.T) {
	mockDB := new(MockDB)
	sectionService := service.NewSectionService(mockDB)

	sections := []models.Section{
		{Title: "Section1", Description: "Description1"},
		{Title: "Section2", Description: "Description2"},
	}
	sectionDTOs := []section_dto.SectionListingDTO{
		{Title: "Section1", Description: "Description1"},
		{Title: "Section2", Description: "Description2"},
	}

	mockDB.On("Preload", "Project").Return(mockDB)
	mockDB.On("Preload", "User").Return(mockDB)
	mockDB.On("Find", &sections).Return(mockDB)

	result, err := sectionService.GetSections()
	assert.Nil(t, err)
	assert.ElementsMatch(t, sectionDTOs, result)
}

func TestGetSectionByIDService(t *testing.T) {
	mockDB := new(MockDB)
	sectionService := service.NewSectionService(mockDB)

	section := models.Section{Title: "Section1", Description: "Description1"}
	sectionDTO := section_dto.SectionListingDTO{
		ID:          1,
		Title:       "Section1",
		Description: "Description1",
	}

	mockDB.On("Preload", "Project").Return(mockDB)
	mockDB.On("Preload", "User").Return(mockDB)
	mockDB.On("First", &section, uint(1)).Return(mockDB)

	result, err := sectionService.GetSectionByID(1)
	assert.Nil(t, err)
	assert.Equal(t, sectionDTO, result)
}

func TestCreateSectionService(t *testing.T) {
	mockDB := new(MockDB)
	sectionService := service.NewSectionService(mockDB)
	section := models.Section{Title: "New Section", Description: "Description"}

	mockDB.On("Create", &section).Return(mockDB)

	err := sectionService.CreateSection(section)
	assert.Nil(t, err)
}

func TestUpdateSectionService(t *testing.T) {
	mockDB := new(MockDB)
	sectionService := service.NewSectionService(mockDB)
	section := models.Section{Title: "Updated Section", Description: "Updated Description"}

	mockDB.On("First", &models.Section{}, uint(1)).Return(mockDB)
	mockDB.On("Model", &models.Section{}).Return(mockDB)
	mockDB.On("Updates", section).Return(mockDB)

	err := sectionService.UpdateSection(1, section)
	assert.Nil(t, err)
}

func TestDeleteSectionService(t *testing.T) {
	mockDB := new(MockDB)
	sectionService := service.NewSectionService(mockDB)

	mockDB.On("Delete", &models.Section{}, uint(1)).Return(mockDB)

	err := sectionService.DeleteSection(1)
	assert.Nil(t, err)
}
