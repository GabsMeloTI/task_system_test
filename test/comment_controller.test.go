package test

import (
	"awesomeProject/dto/comment_dto"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Mock do CommentService
type MockCommentService struct {
	mock.Mock
}

func (m *MockCommentService) GetAllComments() ([]comment_dto.CommentListingDTO, error) {
	args := m.Called()
	return args.Get(0).([]comment_dto.CommentListingDTO), args.Error(1)
}

func (m *MockCommentService) GetCommentByID(id uint) (comment_dto.CommentListingDTO, error) {
	args := m.Called(id)
	return args.Get(0).(comment_dto.CommentListingDTO), args.Error(1)
}

func (m *MockCommentService) CreateComment(content string, imageURL string) error {
	args := m.Called(content, imageURL)
	return args.Error(0)
}

func (m *MockCommentService) UpdateComment(id uint, content string, imageURL string) error {
	args := m.Called(id, content, imageURL)
	return args.Error(0)
}

func (m *MockCommentService) DeleteComment(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCommentService) UploadToS3(file multipart.File, filename string) (string, error) {
	args := m.Called(file, filename)
	return args.Get(0).(string), args.Error(1)
}

func TestGetCommentsController(t *testing.T) {
	mockService := new(MockCommentService)
	controller := &CommentController{Service: mockService}

	expectedComments := []comment_dto.CommentListingDTO{
		{ID: 1, Content: "Comment 1", PublishedAt: time.Now()},
		{ID: 2, Content: "Comment 2", PublishedAt: time.Now()},
	}

	mockService.On("GetAllComments").Return(expectedComments, nil)

	req := httptest.NewRequest("GET", "/comments", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetComments)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var comments []comment_dto.CommentListingDTO
	json.NewDecoder(rr.Body).Decode(&comments)

	assert.Len(t, comments, 2)
	assert.Equal(t, "Comment 1", comments[0].Content)
}

func TestCreateCommentController(t *testing.T) {
	mockService := new(MockCommentService)
	controller := &CommentController{Service: mockService}

	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	_ = writer.WriteField("content", "New Comment")
	_ = writer.Close()

	req := httptest.NewRequest("POST", "/comments", form)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateComment)

	mockService.On("CreateComment", "New Comment", "").Return(nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateCommentController(t *testing.T) {
	mockService := new(MockCommentService)
	controller := &CommentController{Service: mockService}

	commentDTO := struct {
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}{
		Content:  "Updated Comment",
		ImageURL: "http://example.com/new-image.jpg",
	}

	body, _ := json.Marshal(commentDTO)
	req := httptest.NewRequest("PUT", "/comments/1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateComment)

	mockService.On("UpdateComment", uint(1), commentDTO.Content, commentDTO.ImageURL).Return(nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteCommentController(t *testing.T) {
	mockService := new(MockCommentService)
	controller := &CommentController{Service: mockService}

	req := httptest.NewRequest("DELETE", "/comments/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteComment)

	mockService.On("DeleteComment", uint(1)).Return(nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
