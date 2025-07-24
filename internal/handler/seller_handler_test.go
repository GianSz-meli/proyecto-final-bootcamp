package handler

import (
	pkgError "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/bootcamp-go/web/request"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockSellerService struct {
	CreateFunc  func(seller models.Seller) (models.Seller, error)
	GetAllFunc  func() ([]models.Seller, error)
	GetByIdFunc func(id int) (models.Seller, error)
	DeleteFunc  func(id int) error
	UpdateFunc  func(seller models.Seller) (models.Seller, error)
}

func (m *MockSellerService) Create(seller models.Seller) (models.Seller, error) {
	return m.CreateFunc(seller)
}

func (m *MockSellerService) GetAll() ([]models.Seller, error) {
	return m.GetAllFunc()
}

func (m *MockSellerService) GetById(id int) (models.Seller, error) {
	return m.GetByIdFunc(id)
}

func (m *MockSellerService) Delete(id int) error {
	return m.DeleteFunc(id)
}

func (m *MockSellerService) Update(seller models.Seller) (models.Seller, error) {
	return m.UpdateFunc(seller)
}

var reqBodySuccesfull = "{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n} "

func TestSellerHandler_Create_ValidateRequest(t *testing.T) {
	tests := []struct {
		name        string
		reqBody     io.Reader
		contentType string
		service     *MockSellerService
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return error when json request body is invalid",
			reqBody:     strings.NewReader("{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n "),
			contentType: "application/json",
			service:     &MockSellerService{},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusBadRequest
				require.Equal(t, expectedCode, response.Code)
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
				require.Contains(t, response.Body.String(), request.ErrRequestJSONInvalid.Error())

			},
		},
		{
			name:        "should return error when request body is text/plain",
			reqBody:     strings.NewReader(reqBodySuccesfull),
			contentType: "text/plain",
			service:     &MockSellerService{},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusBadRequest
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), request.ErrRequestContentTypeNotJSON.Error())
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))

			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			hd := NewSellerHandler(test.service)
			hdFunc := hd.Create()
			//Act
			request := httptest.NewRequest("POST", "/", test.reqBody)
			request.Header.Set("Content-Type", test.contentType)
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
		})
	}
}

func TestSellerHandler_Create_ValidateRequestData(t *testing.T) {
	tests := []struct {
		name        string
		reqBody     io.Reader
		contentType string
		service     *MockSellerService
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return error when Cid is not present",
			reqBody:     strings.NewReader("{ \n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n} "),
			contentType: "application/json",
			service:     &MockSellerService{},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusUnprocessableEntity
				expectedMissingField := "Cid"
				require.Contains(t, response.Body.String(), pkgError.ErrUnprocessableEntity.Error())
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), expectedMissingField)
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			},
		},
		{
			name:        "should return error when CompanyName is not present",
			reqBody:     strings.NewReader("{   \n    \"cid\": \"GDJ2SJ3\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n} "),
			contentType: "application/json",
			service:     &MockSellerService{},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusUnprocessableEntity
				expectedMissingField := "CompanyName"
				require.Contains(t, response.Body.String(), pkgError.ErrUnprocessableEntity.Error())
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), expectedMissingField)
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			},
		},
		{
			name:        "should return error when Address is not present",
			reqBody:     strings.NewReader("{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n} "),
			contentType: "application/json",
			service:     &MockSellerService{},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusUnprocessableEntity
				expectedMissingField := "Address"
				require.Contains(t, response.Body.String(), pkgError.ErrUnprocessableEntity.Error())
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), expectedMissingField)
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			},
		},
		{
			name:        "should return error when Telephone is not present",
			reqBody:     strings.NewReader("{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"locality_id\": 1\n} "),
			contentType: "application/json",
			service:     &MockSellerService{},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusUnprocessableEntity
				expectedMissingField := "Telephone"
				require.Contains(t, response.Body.String(), pkgError.ErrUnprocessableEntity.Error())
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), expectedMissingField)
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			},
		},
		{
			name:        "should return error when LocalityId is not present",
			reqBody:     strings.NewReader("{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\"} "),
			contentType: "application/json",
			service:     &MockSellerService{},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusUnprocessableEntity
				expectedMissingField := "LocalityId"
				require.Contains(t, response.Body.String(), pkgError.ErrUnprocessableEntity.Error())
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), expectedMissingField)
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			hd := NewSellerHandler(test.service)
			hdFunc := hd.Create()
			//Act
			request := httptest.NewRequest("POST", "/", test.reqBody)
			request.Header.Set("Content-Type", test.contentType)
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
		})
	}
}

func TestSellerHandler_Create_MySQL_Errors(t *testing.T) {

	tests := []struct {
		name        string
		reqBody     io.Reader
		contentType string
		service     *MockSellerService
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return 409 conflict error when service returns duplicate entry error",
			reqBody:     strings.NewReader(reqBodySuccesfull),
			contentType: "application/json",
			service: &MockSellerService{
				CreateFunc: func(seller models.Seller) (models.Seller, error) {
					mysqlError := &mysql.MySQLError{
						Number:  1062,
						Message: "Duplicate entry",
					}
					return models.Seller{}, mysqlError
				},
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusConflict
				require.Contains(t, response.Body.String(), pkgError.ErrConflict.Error())
				require.Equal(t, expectedCode, response.Code)
			},
		},
		{
			name:        "should return 409 conflict error when service returns foreign key constraint error",
			reqBody:     strings.NewReader(reqBodySuccesfull),
			contentType: "application/json",
			service: &MockSellerService{
				CreateFunc: func(seller models.Seller) (models.Seller, error) {
					mysqlError := &mysql.MySQLError{
						Number:  1452,
						Message: "Cannot add or update a child row: a foreign key constraint fails",
					}
					return models.Seller{}, mysqlError
				},
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusConflict
				require.Contains(t, response.Body.String(), pkgError.ErrConflict.Error())
				require.Equal(t, expectedCode, response.Code)
			},
		},
		{
			name:        "should return 400 bad request error when service returns column cannot be null",
			reqBody:     strings.NewReader(reqBodySuccesfull),
			contentType: "application/json",
			service: &MockSellerService{
				CreateFunc: func(seller models.Seller) (models.Seller, error) {
					mysqlError := &mysql.MySQLError{
						Number:  1048,
						Message: "Column 'xxx' cannot be null",
					}
					return models.Seller{}, mysqlError
				},
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusBadRequest
				require.Contains(t, response.Body.String(), pkgError.ErrBadRequest.Error())
				require.Equal(t, expectedCode, response.Code)
			},
		},
		{
			name:        "should return 409 conflict error when service returns cannot delete or update",
			reqBody:     strings.NewReader(reqBodySuccesfull),
			contentType: "application/json",
			service: &MockSellerService{
				CreateFunc: func(seller models.Seller) (models.Seller, error) {
					mysqlError := &mysql.MySQLError{
						Number:  1451,
						Message: "Cannot delete or update a parent row: a foreign key constraint fails",
					}
					return models.Seller{}, mysqlError
				},
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusConflict
				require.Contains(t, response.Body.String(), pkgError.ErrConflict.Error())
				require.Equal(t, expectedCode, response.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			hd := NewSellerHandler(test.service)
			hdFunc := hd.Create()
			//Act
			request := httptest.NewRequest("POST", "/", test.reqBody)
			request.Header.Set("Content-Type", test.contentType)
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
		})
	}
}
