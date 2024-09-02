package test

import (
	"awesomeProject/dto/label_dto"
	"awesomeProject/models"
	"awesomeProject/service"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLabels(t *testing.T) {
	serviceMock := new(service.MockLabelService)
	controller := LabelController{Service: serviceMock}

	expectedLabels := []models.Label{
		{ID: 1, Name: "Bug", Color: "Red"},
		{ID: 2, Name: "Feature", Color: "Blue"},
	}
	expectedDTOs := []label_dto.LabelListingDTO{
		{ID: 1, Name: "Bug", Color: "Red"},
		{ID: 2, Name: "Feature", Color: "Blue"},
	}

	serviceMock.On("GetAllLabels").Return(expectedDTOs, nil)

	req, err := http.NewRequest("GET", "/labels", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetLabels)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `[
		{"id":1,"name":"Bug","color":"Red"},
		{"id":2,"name":"Feature","color":"Blue"}
	]`, rr.Body.String())
}

func TestGetLabelByID(t *testing.T) {
	serviceMock := new(service.MockLabelService)
	controller := LabelController{Service: serviceMock}

	expectedDTO := label_dto.LabelListingDTO{ID: 1, Name: "Bug", Color: "Red"}
	serviceMock.On("GetLabelByID", uint(1)).Return(expectedDTO, nil)

	req, err := http.NewRequest("GET", "/labels/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetLabelByID)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"id":1,"name":"Bug","color":"Red"}`, rr.Body.String())
}

func TestCreateLabel(t *testing.T) {
	serviceMock := new(service.MockLabelService)
	controller := LabelController{Service: serviceMock}

	label := models.Label{Name: "Bug", Color: "Red"}
	labelJSON, err := json.Marshal(label)
	assert.NoError(t, err)

	serviceMock.On("CreateLabel", label).Return(nil)

	req, err := http.NewRequest("POST", "/labels", bytes.NewBuffer(labelJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateLabel)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateLabel(t *testing.T) {
	serviceMock := new(service.MockLabelService)
	controller := LabelController{Service: serviceMock}

	label := models.Label{ID: 1, Name: "Bug", Color: "Red"}
	labelJSON, err := json.Marshal(label)
	assert.NoError(t, err)

	serviceMock.On("UpdateLabel", uint(1), label).Return(nil)

	req, err := http.NewRequest("PUT", "/labels/1", bytes.NewBuffer(labelJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateLabel)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteLabel(t *testing.T) {
	serviceMock := new(service.MockLabelService)
	controller := LabelController{Service: serviceMock}

	serviceMock.On("DeleteLabel", uint(1)).Return(nil)

	req, err := http.NewRequest("DELETE", "/labels/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.DeleteLabel)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
