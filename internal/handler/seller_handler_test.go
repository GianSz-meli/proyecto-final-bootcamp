package handler

import (
	"ProyectoFinal/mocks"
	pkgError "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"github.com/bootcamp-go/web/request"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var reqBodySuccesfull = "{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n} "

// Test for create bad request
func TestSellerHandler_Create_ValidateRequest(t *testing.T) {
	tests := []struct {
		name        string
		reqBody     io.Reader
		contentType string
		service     *mocks.MockSellerService
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return error when json request body is invalid",
			reqBody:     strings.NewReader("{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n "),
			contentType: "application/json",
			service:     &mocks.MockSellerService{},
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
			service:     &mocks.MockSellerService{},
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
			require.Equal(t, 0, test.service.Spy.CountCreateFunc)
		})
	}
}

// Test for create fails
func TestSellerHandler_Create_ValidateRequestData(t *testing.T) {
	tests := []struct {
		name        string
		reqBody     io.Reader
		contentType string
		service     *mocks.MockSellerService
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return error when Cid is not present",
			reqBody:     strings.NewReader("{ \n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n} "),
			contentType: "application/json",
			service:     &mocks.MockSellerService{},
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
			service:     &mocks.MockSellerService{},
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
			service:     &mocks.MockSellerService{},
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
			service:     &mocks.MockSellerService{},
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
			service:     &mocks.MockSellerService{},
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
			require.Equal(t, 0, test.service.Spy.CountCreateFunc)
		})
	}
}

// Test for create mysql errors
func TestSellerHandler_Create_MySQL_Errors(t *testing.T) {

	tests := []struct {
		name        string
		reqBody     io.Reader
		contentType string
		service     *mocks.MockSellerService
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return 409 conflict error when service returns duplicate entry error",
			reqBody:     strings.NewReader(reqBodySuccesfull),
			contentType: "application/json",
			service: &mocks.MockSellerService{
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
			service: &mocks.MockSellerService{
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
			service: &mocks.MockSellerService{
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
			service: &mocks.MockSellerService{
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
			require.Equal(t, 1, test.service.Spy.CountCreateFunc)
		})
	}
}

// Test for create ok
func TestSellerHandler_Create_Success(t *testing.T) {
	newSeller := models.Seller{
		Id:          1,
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	tests := []struct {
		name        string
		reqBody     io.Reader
		contentType string
		service     *mocks.MockSellerService
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return 201 when create seller is successful",
			reqBody:     strings.NewReader(reqBodySuccesfull),
			contentType: "application/json",
			service: &mocks.MockSellerService{
				CreateFunc: func(seller models.Seller) (models.Seller, error) {
					return newSeller, nil
				},
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedSellerDocJson, err := json.Marshal(newSeller.ModelToDoc())
				require.NoError(t, err)
				expectedBody := `{"data":[{"id":1,"cid":"GDJ2SJ3","company_name":"Farm to Table Produce Hub","address":"812 Cypress Way, Denver, CO 80201","telephone":"+1-555-1901","locality_id":1}]}`
				expectedCode := http.StatusCreated
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
				require.Equal(t, expectedCode, response.Code)
				require.JSONEq(t, expectedBody, response.Body.String())
				require.Contains(t, response.Body.String(), string(expectedSellerDocJson))

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
			require.Equal(t, 1, test.service.Spy.CountCreateFunc)

		})
	}
}

// Test for Get All
func TestSellerHandler_GetAll(t *testing.T) {
	tests := []struct {
		name       string
		service    *mocks.MockSellerService
		assertFunc func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name: "should return 200  when service returns all sellers",
			service: &mocks.MockSellerService{
				GetAllFunc: func() ([]models.Seller, error) {
					sellers := []models.Seller{
						{
							Id:          1,
							Cid:         "GDJ2SJ3",
							CompanyName: "Farm to Table Produce Hub",
							Address:     "812 Cypress Way, Denver, CO 80201",
							Telephone:   "+1-555-1901",
							LocalityId:  1,
						},
						{
							Id:          2,
							Cid:         "AD3HF4",
							CompanyName: "Farm to Table Produce Hub Two",
							Address:     "82 House Way, Alaska, CO 23531",
							Telephone:   "+5-165-7201",
							LocalityId:  1,
						},
					}
					return sellers, nil
				},
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedBody := `{"data":[{"id":1,"cid":"GDJ2SJ3","company_name":"Farm to Table Produce Hub","address":"812 Cypress Way, Denver, CO 80201","telephone":"+1-555-1901","locality_id":1},{"id":2,"cid":"AD3HF4","company_name":"Farm to Table Produce Hub Two","address":"82 House Way, Alaska, CO 23531","telephone":"+5-165-7201","locality_id":1}]}`
				expectedCode := http.StatusOK
				require.Equal(t, expectedCode, response.Code)
				require.JSONEq(t, expectedBody, response.Body.String())
			},
		},
		{
			name: "should return 500 internal server error error when service returns a mysql error not mapped",
			service: &mocks.MockSellerService{
				GetAllFunc: func() ([]models.Seller, error) {
					mysqlError := &mysql.MySQLError{
						Number:  1054,
						Message: "Unknown column 'xxx' in 'xxx'",
					}
					return nil, mysqlError
				},
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusInternalServerError
				require.Contains(t, response.Body.String(), http.StatusText(expectedCode))
				require.Equal(t, expectedCode, response.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			hd := NewSellerHandler(test.service)
			hdFunc := hd.GetAll()
			//Act
			request := httptest.NewRequest("GET", "/", strings.NewReader(""))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
			require.Equal(t, 1, test.service.Spy.CountGetAllFunc)
		})
	}
}
