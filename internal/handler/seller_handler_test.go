package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/mocks"
	pkgError "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	pkgRequest "github.com/bootcamp-go/web/request"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var createRequestSuccess = models.CreateSellerRequest{
	Cid:         "1",
	CompanyName: "Farm to Table Produce Hub",
	Address:     "812 Cypress Way, Denver, CO 80201",
	Telephone:   "+1-555-1901",
	LocalityId:  1,
}

func TestSellerHandler_Create_ValidateRequest(t *testing.T) {
	reqBodySuccess := "{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n} "
	tests := []struct {
		name                 string
		reqBody              io.Reader
		contentType          string
		expectedRequestError error
	}{
		{
			name:                 "should return error when json request body is invalid",
			reqBody:              strings.NewReader("{   \n    \"cid\": \"GDJ2SJ3\",\n    \"company_name\": \"Farm to Table Produce Hub\",\n    \"address\": \"812 Cypress Way, Denver, CO 80201\",\n    \"telephone\": \"+1-555-1901\",\n    \"locality_id\": 1\n "),
			contentType:          "application/json",
			expectedRequestError: pkgRequest.ErrRequestJSONInvalid,
		},
		{
			name:                 "should return error when request body is text/plain",
			reqBody:              strings.NewReader(reqBodySuccess),
			contentType:          "text/plain",
			expectedRequestError: pkgRequest.ErrRequestContentTypeNotJSON,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &mocks.MockSellerService{}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Create()
			//Act
			request := httptest.NewRequest("POST", "/", test.reqBody)
			request.Header.Set("Content-Type", test.contentType)
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			expectedCode := http.StatusBadRequest
			require.Equal(t, expectedCode, response.Code)
			require.Contains(t, response.Body.String(), test.expectedRequestError.Error())
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.Equal(t, 0, srv.Spy.CountCreateFunc)
		})
	}
}

func TestSellerHandler_Create_ValidateRequestData(t *testing.T) {

	tests := []struct {
		name                 string
		createRequest        models.CreateSellerRequest
		expectedMissingField string
	}{
		{
			name: "should return error when Cid is not present",
			createRequest: models.CreateSellerRequest{
				CompanyName: "Farm to Table Produce Hub",
				Address:     "812 Cypress Way, Denver, CO 80201",
				Telephone:   "+1-555-1901",
				LocalityId:  1,
			},
			expectedMissingField: "Cid",
		},
		{
			name: "should return error when CompanyName is not present",
			createRequest: models.CreateSellerRequest{
				Cid:        "1",
				Address:    "812 Cypress Way, Denver, CO 80201",
				Telephone:  "+1-555-1901",
				LocalityId: 1,
			},
			expectedMissingField: "CompanyName",
		},
		{
			name: "should return error when Address is not present",
			createRequest: models.CreateSellerRequest{
				Cid:         "1",
				CompanyName: "Farm to Table Produce Hub",
				Telephone:   "+1-555-1901",
				LocalityId:  1,
			},
			expectedMissingField: "Address",
		},
		{
			name: "should return error when Telephone is not present",
			createRequest: models.CreateSellerRequest{
				Cid:         "1",
				CompanyName: "Farm to Table Produce Hub",
				Address:     "812 Cypress Way, Denver, CO 80201",
				LocalityId:  1,
			},
			expectedMissingField: "Telephone",
		},
		{
			name: "should return error when LocalityId is not present",
			createRequest: models.CreateSellerRequest{
				Cid:         "1",
				CompanyName: "Farm to Table Produce Hub",
				Address:     "812 Cypress Way, Denver, CO 80201",
				Telephone:   "+1-555-1901",
			},
			expectedMissingField: "LocalityId",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &mocks.MockSellerService{}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Create()

			//Act
			reqBody, _ := json.Marshal(test.createRequest)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert

			expectedCode := http.StatusUnprocessableEntity
			require.Contains(t, response.Body.String(), pkgError.ErrUnprocessableEntity.Error())
			require.Equal(t, expectedCode, response.Code)
			require.Contains(t, response.Body.String(), test.expectedMissingField)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.Equal(t, 0, srv.Spy.CountCreateFunc)
		})
	}
}

func TestSellerHandler_Create_MySQL_Errors(t *testing.T) {

	tests := []struct {
		name        string
		contentType string
		createFunc  func(seller models.Seller) (models.Seller, error)
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name:        "should return 409 conflict error when service returns duplicate entry error",
			contentType: "application/json",
			createFunc: func(seller models.Seller) (models.Seller, error) {
				mysqlError := &mysql.MySQLError{
					Number:  1062,
					Message: "Duplicate entry",
				}
				return models.Seller{}, mysqlError
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusConflict
				require.Contains(t, response.Body.String(), pkgError.ErrConflict.Error())
				require.Equal(t, expectedCode, response.Code)
			},
		},
		{
			name:        "should return 409 conflict error when service returns foreign key constraint error",
			contentType: "application/json",
			createFunc: func(seller models.Seller) (models.Seller, error) {
				mysqlError := &mysql.MySQLError{
					Number:  1452,
					Message: "Cannot add or update a child row: a foreign key constraint fails",
				}
				return models.Seller{}, mysqlError
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusConflict
				require.Contains(t, response.Body.String(), pkgError.ErrConflict.Error())
				require.Equal(t, expectedCode, response.Code)
			},
		},
		{
			name:        "should return 400 bad request error when service returns column cannot be null",
			contentType: "application/json",
			createFunc: func(seller models.Seller) (models.Seller, error) {
				mysqlError := &mysql.MySQLError{
					Number:  1048,
					Message: "Column 'xxx' cannot be null",
				}
				return models.Seller{}, mysqlError
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusBadRequest
				require.Contains(t, response.Body.String(), pkgError.ErrBadRequest.Error())
				require.Equal(t, expectedCode, response.Code)
			},
		},
		{
			name:        "should return 409 conflict error when service returns cannot delete or update",
			contentType: "application/json",
			createFunc: func(seller models.Seller) (models.Seller, error) {
				mysqlError := &mysql.MySQLError{
					Number:  1451,
					Message: "Cannot delete or update a parent row: a foreign key constraint fails",
				}
				return models.Seller{}, mysqlError
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
			srv := &mocks.MockSellerService{
				CreateFunc: test.createFunc,
			}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Create()
			//Act
			reqBody, _ := json.Marshal(createRequestSuccess)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			request.Header.Set("Content-Type", test.contentType)
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
			require.Equal(t, 1, srv.Spy.CountCreateFunc)
		})
	}
}

func TestSellerHandler_Create_Success(t *testing.T) {
	newSeller := models.Seller{
		Id:          1,
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}

	t.Run("should return 201 when create seller is successful", func(t *testing.T) {
		//Arrange
		srv := &mocks.MockSellerService{
			CreateFunc: func(seller models.Seller) (models.Seller, error) {
				return newSeller, nil
			},
		}
		hd := NewSellerHandler(srv)
		hdFunc := hd.Create()

		//Act
		reqBody, _ := json.Marshal(createRequestSuccess)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		//Assert
		expectedSellerDocJson, _ := json.Marshal(newSeller.ModelToDoc())
		expectedBody := fmt.Sprintf(`{"data":[%s]}`, string(expectedSellerDocJson))
		expectedCode := http.StatusCreated
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, 1, srv.Spy.CountCreateFunc)

	})
}

func TestSellerHandler_GetAll(t *testing.T) {
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
	tests := []struct {
		name       string
		getAllFunc func() ([]models.Seller, error)
		assertFunc func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name: "should return 200  when service returns all sellers",
			getAllFunc: func() ([]models.Seller, error) {
				return sellers, nil
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedSellersJson, _ := json.Marshal(sellers)
				expectedBody := fmt.Sprintf(`{"data":%s}`, string(expectedSellersJson))
				expectedCode := http.StatusOK

				require.Equal(t, expectedCode, response.Code)
				require.JSONEq(t, expectedBody, response.Body.String())
			},
		},
		{
			name: "should return 500 internal server error error when service returns a mysql error not mapped",
			getAllFunc: func() ([]models.Seller, error) {
				mysqlError := &mysql.MySQLError{
					Number:  1054,
					Message: "Unknown column 'xxx' in 'xxx'",
				}
				return nil, mysqlError
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
			srv := &mocks.MockSellerService{
				GetAllFunc: test.getAllFunc,
			}
			hd := NewSellerHandler(srv)
			hdFunc := hd.GetAll()
			//Act
			request := httptest.NewRequest("GET", "/", nil)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
			require.Equal(t, 1, srv.Spy.CountGetAllFunc)
		})
	}
}

func TestSellerHandler_GetById_Bad_PathParam(t *testing.T) {

	tests := []struct {
		name      string
		pathParam string
	}{
		{
			name:      "should return 400 bad request error when path param is not a number",
			pathParam: "id=2",
		},
		{
			name:      "should return 400 bad request error when path param is less to zero",
			pathParam: "-2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &mocks.MockSellerService{}
			hd := NewSellerHandler(srv)
			hdFunc := hd.GetById()
			//Act
			target := fmt.Sprintf("/%s", test.pathParam)
			request := httptest.NewRequest("GET", target, nil)
			request = utils.AddPathParamToRequest(request, "id", test.pathParam)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			expectedCode := http.StatusBadRequest
			require.Equal(t, expectedCode, response.Code)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.Equal(t, 0, srv.Spy.CountGetByIdFunc)

		})
	}
}

func TestSellerHandler_GetById(t *testing.T) {
	seller := models.Seller{
		Id:          1,
		Cid:         "2G2A",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	tests := []struct {
		name        string
		getByIdFunc func(id int) (models.Seller, error)
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			name: "should return 404 not found error when seller id not exist",
			getByIdFunc: func(id int) (models.Seller, error) {
				return models.Seller{}, pkgError.ErrNotFound
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusNotFound
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), pkgError.ErrNotFound.Error())
			},
		},
		{
			name: "should return 200 when seller id exist",
			getByIdFunc: func(id int) (models.Seller, error) {
				return seller, nil
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedSellerJson, _ := json.Marshal(seller)
				expectedBody := fmt.Sprintf(`{"data":[%s]}`, string(expectedSellerJson))
				expectedCode := http.StatusOK

				require.Equal(t, expectedCode, response.Code)
				require.JSONEq(t, expectedBody, response.Body.String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &mocks.MockSellerService{
				GetByIdFunc: test.getByIdFunc,
			}
			hd := NewSellerHandler(srv)
			hdFunc := hd.GetById()
			//Act
			request := httptest.NewRequest("GET", "/1", nil)
			request = utils.AddPathParamToRequest(request, "id", "1")
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
			require.Equal(t, 1, srv.Spy.CountGetByIdFunc)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		})
	}
}

func TestSellerHandler_Delete_Bad_PathParam(t *testing.T) {

	tests := []struct {
		name      string
		pathParam string
	}{
		{
			name:      "should return 400 bad request error when path param is not a number",
			pathParam: "id=2",
		},
		{
			name:      "should return 400 bad request error when path param is less to zero",
			pathParam: "-2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &mocks.MockSellerService{}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Delete()
			//Act
			target := fmt.Sprintf("/%s", test.pathParam)
			request := httptest.NewRequest("DELETE", target, nil)
			request = utils.AddPathParamToRequest(request, "id", test.pathParam)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			expectedCode := http.StatusBadRequest
			require.Equal(t, expectedCode, response.Code)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.Equal(t, 0, srv.Spy.CountDeleteFunc)

		})
	}
}

func TestSellerHandler_Delete(t *testing.T) {
	tests := []struct {
		name        string
		deleteFunc  func(id int) error
		assertFunc  func(t *testing.T, response *httptest.ResponseRecorder)
		contentType string
	}{
		{
			name: "should return 404 not found error when seller id not exist",
			deleteFunc: func(id int) error {
				return pkgError.ErrNotFound
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusNotFound
				require.Equal(t, expectedCode, response.Code)
				require.Contains(t, response.Body.String(), pkgError.ErrNotFound.Error())
				require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			},
			contentType: "application/json",
		},
		{
			name: "should return 204 when delete is success",
			deleteFunc: func(id int) error {
				return nil
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedCode := http.StatusNoContent
				require.Equal(t, expectedCode, response.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &mocks.MockSellerService{
				DeleteFunc: test.deleteFunc,
			}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Delete()
			//Act
			request := httptest.NewRequest("DELETE", "/1", nil)
			request = utils.AddPathParamToRequest(request, "id", "1")
			if test.contentType != "" {
				request.Header.Set("Content-Type", test.contentType)
			}
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			test.assertFunc(t, response)
			require.Equal(t, 1, srv.Spy.CountDeleteFunc)
		})
	}
}
