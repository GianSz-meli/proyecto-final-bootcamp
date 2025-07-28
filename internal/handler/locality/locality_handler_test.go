package locality

import (
	"ProyectoFinal/mocks/locality"
	pkgError "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	pkgRequest "github.com/bootcamp-go/web/request"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLocalityHandler_Create_ValidateRequest_Errors(t *testing.T) {
	tests := []struct {
		name                 string
		reqBody              io.Reader
		contentType          string
		expectedRequestError error
	}{
		{
			name:                 "should return error when json request body is invalid",
			reqBody:              strings.NewReader(`{"id":15,"locality_name":"test","province_name":"Lombardy","country_name":"Italy"`),
			contentType:          "application/json",
			expectedRequestError: pkgRequest.ErrRequestJSONInvalid,
		},
		{
			name:                 "should return error when request body is text/plain",
			reqBody:              strings.NewReader(`{"id":15,"locality_name":"test","province_name":"Lombardy","country_name":"Italy"}`),
			contentType:          "text/plain",
			expectedRequestError: pkgRequest.ErrRequestContentTypeNotJSON,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := new(locality.MockLocalityService)
			hd := NewLocalityHandler(srv)
			hdFunc := hd.Create()
			expectedCode := http.StatusBadRequest

			//Act
			request := httptest.NewRequest("POST", "/", test.reqBody)
			request.Header.Set("Content-Type", test.contentType)
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			require.Equal(t, expectedCode, response.Code)
			require.Contains(t, response.Body.String(), test.expectedRequestError.Error())
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			srv.AssertNotCalled(t, "Create")
		})
	}
}
func TestLocalityHandler_Create_ValidateRequestData_Errors(t *testing.T) {
	tests := []struct {
		name                 string
		createRequest        models.LocalityCreateRequest
		expectedMissingField string
	}{
		{
			name: "should return error when LocalityName is not present",
			createRequest: models.LocalityCreateRequest{
				Id:           1,
				ProvinceName: "Lombardy",
				CountryName:  "Italy",
			},
			expectedMissingField: "LocalityName",
		},
		{
			name: "should return error when ProvinceName is not present",
			createRequest: models.LocalityCreateRequest{
				Id:           1,
				LocalityName: "test",
				CountryName:  "Italy",
			},
			expectedMissingField: "ProvinceName",
		},
		{
			name: "should return error when CountryName is not present",
			createRequest: models.LocalityCreateRequest{
				Id:           1,
				ProvinceName: "Lombardy",
				LocalityName: "test",
			},
			expectedMissingField: "CountryName",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &locality.MockLocalityService{}
			hd := NewLocalityHandler(srv)
			hdFunc := hd.Create()
			expectedCode := http.StatusUnprocessableEntity

			//Act
			reqBody, _ := json.Marshal(test.createRequest)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert

			require.Contains(t, response.Body.String(), pkgError.ErrUnprocessableEntity.Error())
			require.Equal(t, expectedCode, response.Code)
			require.Contains(t, response.Body.String(), test.expectedMissingField)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			srv.AssertNotCalled(t, "Create")
		})
	}
}
func TestLocalityHandler_Create_Errors(t *testing.T) {
	createRequestSuccess := models.LocalityCreateRequest{
		LocalityName: "test",
		ProvinceName: "Lombardy",
		CountryName:  "Italy",
	}
	tests := []struct {
		name          string
		contentType   string
		expectedCode  int
		expectedError error
	}{
		{
			name:          "should return 409 conflict error when service returns conflict error",
			contentType:   "application/json",
			expectedCode:  http.StatusConflict,
			expectedError: pkgError.ErrConflict,
		},
		{
			name:          "should return 400 bad request error when service returns bad request error",
			contentType:   "application/json",
			expectedCode:  http.StatusBadRequest,
			expectedError: pkgError.ErrBadRequest,
		},
		{
			name:          "should return 404 not found error when service returns not found error",
			contentType:   "application/json",
			expectedCode:  http.StatusNotFound,
			expectedError: pkgError.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &locality.MockLocalityService{}
			srv.On("Create", createRequestSuccess.DocToModel()).Return(models.Locality{}, test.expectedError)
			hd := NewLocalityHandler(srv)
			hdFunc := hd.Create()
			//Act
			reqBody, _ := json.Marshal(createRequestSuccess)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			request.Header.Set("Content-Type", test.contentType)
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			require.Contains(t, response.Body.String(), test.expectedError.Error())
			require.Equal(t, test.expectedCode, response.Code)
			srv.AssertExpectations(t)
		})
	}
}
func TestLocalityHandler_Create_Success(t *testing.T) {
	createRequestSuccess := models.LocalityCreateRequest{
		LocalityName: "test",
		ProvinceName: "Lombardy",
		CountryName:  "Italy",
	}

	newLocality := models.Locality{
		Id:           1,
		LocalityName: "test",
		Province: models.Province{
			Id:           1,
			ProvinceName: "Lombardy",
			Country: models.Country{
				Id:          1,
				CountryName: "Italy",
			},
		},
	}

	//Arrange
	srv := &locality.MockLocalityService{}
	hd := NewLocalityHandler(srv)
	srv.On("Create", createRequestSuccess.DocToModel()).Return(newLocality, nil)
	hdFunc := hd.Create()
	expectedSellerDocJson, _ := json.Marshal(newLocality.ModelToDoc())
	expectedBody := fmt.Sprintf(`{"data":[%s]}`, string(expectedSellerDocJson))
	expectedCode := http.StatusCreated
	//Act
	reqBody, _ := json.Marshal(createRequestSuccess)
	request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	hdFunc(response, request)

	//Assert
	require.Equal(t, "application/json", response.Header().Get("Content-Type"))
	require.Equal(t, expectedCode, response.Code)
	require.JSONEq(t, expectedBody, response.Body.String())
	srv.AssertExpectations(t)

}

func TestLocalityHandler_GetSellersByLocality_All_Error(t *testing.T) {
	tests := []struct {
		name          string
		expectedCode  int
		expectedError error
	}{
		{
			name:          "should return 400 bad request error when service returns bad request error",
			expectedCode:  http.StatusBadRequest,
			expectedError: pkgError.ErrBadRequest,
		},
		{
			name:          "should return 404 not found error when service returns not found error",
			expectedCode:  http.StatusNotFound,
			expectedError: pkgError.ErrNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &locality.MockLocalityService{}
			hd := NewLocalityHandler(srv)
			hdFunc := hd.GetSellersByLocality()
			srv.On("GetSellersByLocalities").Return(nil, test.expectedError)
			expectedBody := fmt.Sprintf(`{"status":"%v", "message":"%s"}`, http.StatusText(test.expectedCode), test.expectedError.Error())

			//Act
			request := httptest.NewRequest("GET", "/reportSellers", nil)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)
			//Assert
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.Equal(t, test.expectedCode, response.Code)
			require.JSONEq(t, expectedBody, response.Body.String())
			srv.AssertExpectations(t)
		})
	}
}
func TestLocalityHandler_GetSellersByLocality_All_Success(t *testing.T) {
	expectedSellerByLocalities := []models.SellersByLocalityReport{
		{
			LocalityId:   1,
			LocalityName: "La Plata",
			SellersCount: 3,
		},
		{
			LocalityId:   2,
			LocalityName: "Rosario",
			SellersCount: 2,
		},
		{
			LocalityId:   3,
			LocalityName: "Cordoba Capital",
			SellersCount: 1,
		},
	}
	//Arrange
	srv := &locality.MockLocalityService{}
	hd := NewLocalityHandler(srv)
	hdFunc := hd.GetSellersByLocality()
	srv.On("GetSellersByLocalities").Return(expectedSellerByLocalities, nil)
	expectedJson, err := json.Marshal(expectedSellerByLocalities)
	expectedBody := fmt.Sprintf(`{"data":%s}`, string(expectedJson))
	expectedCode := http.StatusOK
	//Act
	request := httptest.NewRequest("GET", "/reportSellers", nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	hdFunc(response, request)
	//Assert
	require.NoError(t, err)
	require.Equal(t, "application/json", response.Header().Get("Content-Type"))
	require.Equal(t, expectedCode, response.Code)
	require.JSONEq(t, expectedBody, response.Body.String())
	srv.AssertExpectations(t)
}
func TestLocalityHandler_GetSellersByLocality_Bad_QueryParam(t *testing.T) {

	tests := []struct {
		name       string
		queryParam string
	}{
		{
			name:       "should return 400 bad request error when path param is not a number",
			queryParam: "id=2",
		},
		{
			name:       "should return 400 bad request error when path param is less to zero",
			queryParam: "-2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &locality.MockLocalityService{}
			hd := NewLocalityHandler(srv)
			hdFunc := hd.GetSellersByLocality()
			expectedCode := http.StatusBadRequest

			//Act
			target := fmt.Sprintf("/reportSellers?id=%s", test.queryParam)
			request := httptest.NewRequest("GET", target, nil)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			require.Equal(t, expectedCode, response.Code)
			require.Contains(t, response.Body.String(), pkgError.ErrBadRequest.Error())
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			srv.AssertNotCalled(t, "GetSellersByLocality")
		})
	}
}
func TestLocalityHandler_GetSellersByLocality_Errors(t *testing.T) {

	localityId := 1
	tests := []struct {
		name          string
		expectedCode  int
		expectedError error
	}{
		{
			name:          "should return 400 bad request error when service returns bad request error",
			expectedCode:  http.StatusBadRequest,
			expectedError: pkgError.ErrBadRequest,
		},
		{
			name:          "should return 404 not found error when service returns not found error",
			expectedCode:  http.StatusNotFound,
			expectedError: pkgError.WrapErrNotFound("Seller", "id", localityId),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &locality.MockLocalityService{}
			srv.On("GetSellersByIdLocality", localityId).Return(models.SellersByLocalityReport{}, test.expectedError)
			hd := NewLocalityHandler(srv)
			hdFunc := hd.GetSellersByLocality()
			expectedBody := fmt.Sprintf(`{"status":"%s", "message":"%s"}`, http.StatusText(test.expectedCode), test.expectedError.Error())
			//Act
			request := httptest.NewRequest("GET", fmt.Sprintf("/reportSellers?id=%v", localityId), nil)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			require.JSONEq(t, expectedBody, response.Body.String())
			require.Equal(t, test.expectedCode, response.Code)
			srv.AssertExpectations(t)
		})
	}
}
func TestLocalityHandler_GetSellersByLocality_Success(t *testing.T) {
	expectedSellerByIdLocality := models.SellersByLocalityReport{
		LocalityId:   1,
		LocalityName: "La Plata",
		SellersCount: 3,
	}
	localityId := 1

	//Arrange
	srv := &locality.MockLocalityService{}
	srv.On("GetSellersByIdLocality", localityId).Return(expectedSellerByIdLocality, nil)
	hd := NewLocalityHandler(srv)
	hdFunc := hd.GetSellersByLocality()
	//Act
	request := httptest.NewRequest("GET", fmt.Sprintf("/reportSellers?id=%v", localityId), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	hdFunc(response, request)

	//Assert
	expectedSellerByIdLocalityJSON, err := json.Marshal(expectedSellerByIdLocality)
	require.NoError(t, err)
	expectedBody := fmt.Sprintf(`{"data": %s}`, string(expectedSellerByIdLocalityJSON))
	require.JSONEq(t, expectedBody, response.Body.String())
	require.Equal(t, http.StatusOK, response.Code)
	srv.AssertExpectations(t)
}

func TestLocalityHandler_ReportCarriersByLocality_Bad_QueryParam(t *testing.T) {

	tests := []struct {
		name       string
		queryParam string
	}{
		{
			name:       "should return 400 bad request error when path param is not a number",
			queryParam: "id=2",
		},
		{
			name:       "should return 400 bad request error when path param is less to zero",
			queryParam: "-2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &locality.MockLocalityService{}
			hd := NewLocalityHandler(srv)
			hdFunc := hd.ReportCarriersByLocality()
			expectedCode := http.StatusBadRequest

			//Act
			target := fmt.Sprintf("/reportCarriers?id=%s", test.queryParam)
			request := httptest.NewRequest("GET", target, nil)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			require.Equal(t, expectedCode, response.Code)
			require.Contains(t, response.Body.String(), pkgError.ErrBadRequest.Error())
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			srv.AssertNotCalled(t, "ReportCarriersByLocality")
		})
	}
}
func TestLocalityHandler_ReportCarriersByLocality_Errors(t *testing.T) {

	localityId := 1
	tests := []struct {
		name          string
		expectedCode  int
		expectedError error
	}{
		{
			name:          "should return 400 bad request error when service returns bad request error",
			expectedCode:  http.StatusBadRequest,
			expectedError: pkgError.ErrBadRequest,
		},
		{
			name:          "should return 404 not found error when service returns not found error",
			expectedCode:  http.StatusNotFound,
			expectedError: pkgError.WrapErrNotFound("Locality", "id", localityId),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &locality.MockLocalityService{}
			srv.On("ReportCarriersByLocality", &localityId).Return(nil, test.expectedError)
			hd := NewLocalityHandler(srv)
			hdFunc := hd.ReportCarriersByLocality()
			expectedBody := fmt.Sprintf(`{"status":"%s", "message":"%s"}`, http.StatusText(test.expectedCode), test.expectedError.Error())
			//Act
			request := httptest.NewRequest("GET", fmt.Sprintf("/reportCarriers?id=%v", localityId), nil)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			require.JSONEq(t, expectedBody, response.Body.String())
			require.Equal(t, test.expectedCode, response.Code)
			srv.AssertExpectations(t)
		})
	}
}
func TestLocalityHandler_ReportCarriersByLocality_Success(t *testing.T) {
	expectedCarrierByIdLocality := []models.CarrierReport{
		{
			LocalityId:    1,
			LocalityName:  "La Plata",
			CarriersCount: 3,
		},
	}
	localityId := 1

	//Arrange
	srv := &locality.MockLocalityService{}
	srv.On("ReportCarriersByLocality", &localityId).Return(expectedCarrierByIdLocality, nil)
	hd := NewLocalityHandler(srv)
	hdFunc := hd.ReportCarriersByLocality()
	//Act
	request := httptest.NewRequest("GET", fmt.Sprintf("/reportSellers?id=%v", localityId), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	hdFunc(response, request)

	//Assert
	expectedCarrierByIdLocalityJSON, err := json.Marshal(expectedCarrierByIdLocality)
	require.NoError(t, err)
	expectedBody := fmt.Sprintf(`{"data": %s}`, string(expectedCarrierByIdLocalityJSON))
	require.JSONEq(t, expectedBody, response.Body.String())
	require.Equal(t, http.StatusOK, response.Code)
	srv.AssertExpectations(t)
}
