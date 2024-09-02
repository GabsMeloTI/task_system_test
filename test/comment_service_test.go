package test

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// Mock para o banco de dados
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Preload(association string) *MockDB {
	return m
}

func (m *MockDB) Find(out interface{}) *MockDB {
	args := m.Called(out)
	return args.Get(0).(*MockDB)
}

func (m *MockDB) First(out interface{}, where ...interface{}) *MockDB {
	args := m.Called(out, where)
	return args.Get(0).(*MockDB)
}

func (m *MockDB) Create(value interface{}) *MockDB {
	args := m.Called(value)
	return args.Get(0).(*MockDB)
}

func (m *MockDB) Save(value interface{}) *MockDB {
	args := m.Called(value)
	return args.Get(0).(*MockDB)
}

func (m *MockDB) Delete(value interface{}, where ...interface{}) *MockDB {
	args := m.Called(value, where)
	return args.Get(0).(*MockDB)
}

// Implementar os métodos do MockDB para retornar mocks adequados em cada chamada
func TestGetAllComments(t *testing.T) {
	mockDB := new(MockDB)
	db.DB = mockDB

	expectedComments := []models.Comment{
		{ID: 1, Content: "Comment 1", PublishedAt: time.Now()},
		{ID: 2, Content: "Comment 2", PublishedAt: time.Now()},
	}

	mockDB.On("Preload", "User").Return(mockDB)
	mockDB.On("Preload", "Task").Return(mockDB)
	mockDB.On("Find", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]models.Comment)
		*arg = expectedComments
	}).Return(mockDB)

	service := &CommentService{}
	comments, err := service.GetAllComments()

	assert.NoError(t, err)
	assert.Len(t, comments, 2)
	assert.Equal(t, expectedComments[0].Content, comments[0].Content)
}

func TestCreateComment(t *testing.T) {
	mockDB := new(MockDB)
	db.DB = mockDB

	service := &CommentService{}
	content := "New Comment"
	imageURL := "http://example.com/image.jpg"

	mockDB.On("Create", mock.Anything).Return(mockDB)

	err := service.CreateComment(content, imageURL)

	assert.NoError(t, err)
}

func TestUpdateComment(t *testing.T) {
	mockDB := new(MockDB)
	db.DB = mockDB

	service := &CommentService{}
	id := uint(1)
	content := "Updated Comment"
	imageURL := "http://example.com/new-image.jpg"

	mockDB.On("First", mock.Anything, id).Return(mockDB)
	mockDB.On("Save", mock.Anything).Return(mockDB)

	err := service.UpdateComment(id, content, imageURL)

	assert.NoError(t, err)
}

func TestDeleteComment(t *testing.T) {
	mockDB := new(MockDB)
	db.DB = mockDB

	service := &CommentService{}
	id := uint(1)

	mockDB.On("Delete", mock.Anything, id).Return(mockDB)

	err := service.DeleteComment(id)

	assert.NoError(t, err)
}
