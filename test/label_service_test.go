package test

import (
	"awesomeProject/db"
	"awesomeProject/dto/label_dto"
	"awesomeProject/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(out interface{}, where ...interface{}) *db.DB {
	args := m.Called(out, where)
	return args.Get(0).(*db.DB)
}

func (m *MockDB) First(out interface{}, where ...interface{}) *db.DB {
	args := m.Called(out, where)
	return args.Get(0).(*db.DB)
}

func (m *MockDB) Create(value interface{}) *db.DB {
	args := m.Called(value)
	return args.Get(0).(*db.DB)
}

func (m *MockDB) Save(value interface{}) *db.DB {
	args := m.Called(value)
	return args.Get(0).(*db.DB)
}

func (m *MockDB) Delete(value interface{}, where ...interface{}) *db.DB {
	args := m.Called(value, where)
	return args.Get(0).(*db.DB)
}

func TestGetAllLabels(t *testing.T) {
	mockDB := new(MockDB)
	service := LabelService{}

	labels := []models.Label{
		{ID: 1, Name: "Bug", Color: "Red"},
		{ID: 2, Name: "Feature", Color: "Blue"},
	}

	mockDB.On("Find", &labels).Return(mockDB)
	mockDB.On("Error").Return(nil)

	expectedDTOs := []label_dto.LabelListingDTO{
		{ID: 1, Name: "Bug", Color: "Red"},
		{ID: 2, Name: "Feature", Color: "Blue"},
	}

	labelsDTO, err := service.GetAllLabels()
	assert.NoError(t, err)
	assert.Equal(t, expectedDTOs, labelsDTO)
}

func TestGetLabelByID(t *testing.T) {
	mockDB := new(MockDB)
	service := LabelService{}

	label := models.Label{ID: 1, Name: "Bug", Color: "Red"}

	mockDB.On("First", &label, uint(1)).Return(mockDB)
	mockDB.On("Error").Return(nil)

	expectedDTO := label_dto.LabelListingDTO{ID: 1, Name: "Bug", Color: "Red"}

	labelDTO, err := service.GetLabelByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedDTO, labelDTO)
}

func TestCreateLabel(t *testing.T) {
	mockDB := new(MockDB)
	service := LabelService{}

	label := models.Label{Name: "Bug", Color: "Red"}

	mockDB.On("Create", &label).Return(mockDB)
	mockDB.On("Error").Return(nil)

	err := service.CreateLabel(label)
	assert.NoError(t, err)
}

func TestUpdateLabel(t *testing.T) {
	mockDB := new(MockDB)
	service := LabelService{}

	label := models.Label{ID: 1, Name: "Bug", Color: "Red"}

	mockDB.On("First", &models.Label{}, uint(1)).Return(mockDB)
	mockDB.On("Save", &label).Return(mockDB)
	mockDB.On("Error").Return(nil)

	err := service.UpdateLabel(1, label)
	assert.NoError(t, err)
}

func TestDeleteLabel(t *testing.T) {
	mockDB := new(MockDB)
	service := LabelService{}

	mockDB.On("Delete", &models.Label{}, uint(1)).Return(mockDB)
	mockDB.On("Error").Return(nil)

	err := service.DeleteLabel(1)
	assert.NoError(t, err)
}
