package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	carrierMocks "ProyectoFinal/mocks/carrier"
	customErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/require"
)

func TestCarrierHandler_Create_Success(t *testing.T) {

	mockService := new(carrierMocks.MockCarrierService)
	handler := NewCarrierHandler(mockService)

	requestBody := models.CarrierCreateDTO{
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}

	expectedCarrier := &models.Carrier{
		Id:          1,
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}
	expectedResponse := `{
		"data": {
			"id": 1,
			"cid": "C001",
			"company_name": "Express Delivery",
			"address": "123 Main St",
			"telephone": "1234567890",
			"locality_id": 1
		}
	}`

	mockService.On("Create", requestBody.CreateDtoToModel()).Return(expectedCarrier, nil)

	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/carriers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Create()(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	require.JSONEq(t, expectedResponse, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestCarrierHandler_Create_MalformedJson_BadRequest(t *testing.T) {
	mockService := new(carrierMocks.MockCarrierService)
	handler := NewCarrierHandler(mockService)
	expectedResponseBody := `
	{ 
		"status":"Bad Request",
		"message":"bad request : it was not possible to decode json"
	}`

	req := httptest.NewRequest(http.MethodPost, "/carriers", bytes.NewReader([]byte(`{"invalid": json}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Create()(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.JSONEq(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestCarrierHandler_Create_InvalidRequest_UnprocessableEntity(t *testing.T) {

	mockService := new(carrierMocks.MockCarrierService)
	handler := NewCarrierHandler(mockService)

	testCases := []struct {
		name                 string
		requestBody          models.CarrierCreateDTO
		expectedResponseBody string
	}{
		{
			name: "cid is required",
			requestBody: models.CarrierCreateDTO{
				CompanyName: "Express Delivery",
				Address:     "123 Main St",
				Telephone:   "1234567890",
				LocalityId:  1,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CarrierCreateDTO.Cid' Error:Field validation for 'Cid' failed on the 'required' tag"
			}`,
		},
		{
			name: "company_name is required",
			requestBody: models.CarrierCreateDTO{
				Cid:        "C001",
				Address:    "123 Main St",
				Telephone:  "1234567890",
				LocalityId: 1,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CarrierCreateDTO.CompanyName' Error:Field validation for 'CompanyName' failed on the 'required' tag"
			}`,
		},
		{
			name: "address is required",
			requestBody: models.CarrierCreateDTO{
				Cid:         "C001",
				CompanyName: "Express Delivery",
				Telephone:   "1234567890",
				LocalityId:  1,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CarrierCreateDTO.Address' Error:Field validation for 'Address' failed on the 'required' tag"
			}`,
		},
		{
			name: "telephone is required",
			requestBody: models.CarrierCreateDTO{
				Cid:         "C001",
				CompanyName: "Express Delivery",
				Address:     "123 Main St",
				LocalityId:  1,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CarrierCreateDTO.Telephone' Error:Field validation for 'Telephone' failed on the 'required' tag"
			}`,
		},
		{
			name: "telephone must be numeric",
			requestBody: models.CarrierCreateDTO{
				Cid:         "C001",
				CompanyName: "Express Delivery",
				Address:     "123 Main St",
				Telephone:   "123abc",
				LocalityId:  1,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CarrierCreateDTO.Telephone' Error:Field validation for 'Telephone' failed on the 'numeric' tag"
			}`,
		},
		{
			name: "telephone must be at least 7 digits long",
			requestBody: models.CarrierCreateDTO{
				Cid:         "C001",
				CompanyName: "Express Delivery",
				Address:     "123 Main St",
				Telephone:   "123456",
				LocalityId:  1,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CarrierCreateDTO.Telephone' Error:Field validation for 'Telephone' failed on the 'min' tag"
			}`,
		},
		{
			name: "locality_id is required",
			requestBody: models.CarrierCreateDTO{
				Cid:         "C001",
				CompanyName: "Express Delivery",
				Address:     "123 Main St",
				Telephone:   "1234567890",
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CarrierCreateDTO.LocalityId' Error:Field validation for 'LocalityId' failed on the 'required' tag"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/carriers", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler.Create()(w, req)

			require.Equal(t, http.StatusUnprocessableEntity, w.Code)
			require.JSONEq(t, tc.expectedResponseBody, w.Body.String())
		})
	}

	mockService.AssertExpectations(t)
}

func TestCarrierHandler_Create_DuplicateCid_Conflict(t *testing.T) {
	mockService := new(carrierMocks.MockCarrierService)
	handler := NewCarrierHandler(mockService)

	requestBody := models.CarrierCreateDTO{
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}
	expectedResponseBody := `
		{
			"status":"Conflict",
			"message":"conflict : carriers with cid C001 already exists"
		}
	`

	expectedError := customErrors.WrapErrConflict("carriers", "cid", "C001")
	mockService.On("Create", requestBody.CreateDtoToModel()).Return(nil, expectedError)

	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/carriers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Create()(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	require.JSONEq(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}
