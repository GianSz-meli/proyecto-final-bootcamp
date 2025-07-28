package seller

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/mocks/seller"
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

func TestSellerHandler_Create_ValidateRequest_Errors(t *testing.T) {
	tests := []struct {
		name                 string
		reqBody              io.Reader
		contentType          string
		expectedRequestError error
	}{
		{
			name:                 "should return error when json request body is invalid",
			reqBody:              strings.NewReader(`{"cid":"GDJ2SJ3","company_name":"Farm to Table Produce Hub","address":"812 Cypress Way, Denver, CO 80201","telephone":"+1-555-1901","locality_id":1`),
			contentType:          "application/json",
			expectedRequestError: pkgRequest.ErrRequestJSONInvalid,
		},
		{
			name:                 "should return error when request body is text/plain",
			reqBody:              strings.NewReader(`{"cid":"GDJ2SJ3","company_name":"Farm to Table Produce Hub","address":"812 Cypress Way, Denver, CO 80201","telephone":"+1-555-1901","locality_id":1}`),
			contentType:          "text/plain",
			expectedRequestError: pkgRequest.ErrRequestContentTypeNotJSON,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &seller.MockSellerService{}
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

func TestSellerHandler_Create_ValidateRequestData_Errors(t *testing.T) {

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
			srv := &seller.MockSellerService{}
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

func TestSellerHandler_Create_Errors(t *testing.T) {
	createRequestSuccess := models.CreateSellerRequest{
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	tests := []struct {
		name          string
		contentType   string
		createFunc    func(seller models.Seller) (models.Seller, error)
		expectedCode  int
		expectedError error
	}{
		{
			name:        "should return 409 conflict error when service returns conflict error",
			contentType: "application/json",
			createFunc: func(seller models.Seller) (models.Seller, error) {
				return models.Seller{}, pkgError.ErrConflict
			},
			expectedCode:  http.StatusConflict,
			expectedError: pkgError.ErrConflict,
		},
		{
			name:        "should return 400 bad request error when service returns bad request error",
			contentType: "application/json",
			createFunc: func(seller models.Seller) (models.Seller, error) {
				return models.Seller{}, pkgError.ErrBadRequest
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: pkgError.ErrBadRequest,
		},
		{
			name:        "should return 404 not found error when service returns not found error",
			contentType: "application/json",
			createFunc: func(seller models.Seller) (models.Seller, error) {
				return models.Seller{}, pkgError.ErrNotFound
			},
			expectedCode:  http.StatusNotFound,
			expectedError: pkgError.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &seller.MockSellerService{
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
			require.Contains(t, response.Body.String(), test.expectedError.Error())
			require.Equal(t, test.expectedCode, response.Code)
			require.Equal(t, 1, srv.Spy.CountCreateFunc)
		})
	}
}

func TestSellerHandler_Create_Success(t *testing.T) {
	createRequestSuccess := models.CreateSellerRequest{
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}

	newSeller := models.Seller{
		Id:          1,
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}

	//Arrange
	srv := &seller.MockSellerService{
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
			name: "should return 500 internal server  error when service returns a mysql error not mapped",
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
			srv := &seller.MockSellerService{
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
			srv := &seller.MockSellerService{}
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
	s := models.Seller{
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
				return s, nil
			},
			assertFunc: func(t *testing.T, response *httptest.ResponseRecorder) {
				expectedSellerJson, _ := json.Marshal(s)
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
			srv := &seller.MockSellerService{
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
			srv := &seller.MockSellerService{}
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

func TestSellerHandler_Delete_NotFoundError(t *testing.T) {

	//Arrange
	sellerId := 1
	expectedCode := http.StatusNotFound
	expectedBody := fmt.Sprintf(`{"status":"%s", "message": "%s"}`, http.StatusText(http.StatusNotFound), pkgError.WrapErrNotFound("seller", "id", sellerId))
	//Act
	srv := &seller.MockSellerService{
		DeleteFunc: func(id int) error {
			return pkgError.WrapErrNotFound("seller", "id", sellerId)
		},
	}
	hd := NewSellerHandler(srv)
	hdFunc := hd.Delete()
	request := httptest.NewRequest("PATCH", fmt.Sprintf("/%v", sellerId), nil)
	request = utils.AddPathParamToRequest(request, "id", fmt.Sprintf("%v", sellerId))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	hdFunc(response, request)

	//Assert
	require.Equal(t, expectedCode, response.Code)
	require.JSONEq(t, expectedBody, response.Body.String())
	require.Equal(t, srv.Spy.CountDeleteFunc, 1)
}

func TestSellerHandler_Delete_Success(t *testing.T) {

	//Arrange
	srv := &seller.MockSellerService{
		DeleteFunc: func(id int) error {
			return nil
		},
	}
	hd := NewSellerHandler(srv)
	hdFunc := hd.Delete()
	//Act
	request := httptest.NewRequest("DELETE", "/1", nil)
	request = utils.AddPathParamToRequest(request, "id", "1")
	response := httptest.NewRecorder()
	hdFunc(response, request)

	//Assert
	expectedCode := http.StatusNoContent
	require.Equal(t, expectedCode, response.Code)
	require.Equal(t, 1, srv.Spy.CountDeleteFunc)

}

func TestSellerHandler_Update_Bad_PathParam_Errors(t *testing.T) {

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
			srv := &seller.MockSellerService{}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Update()
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
			require.Equal(t, 0, srv.Spy.CountUpdateFunc)

		})
	}
}

func TestSellerHandler_Update_ValidateRequest_Errors(t *testing.T) {
	tests := []struct {
		name                 string
		reqBody              io.Reader
		contentType          string
		expectedRequestError error
	}{
		{
			name:                 "should return error when json request body is invalid",
			reqBody:              strings.NewReader(`{"cid":"GDJ2SJ3","company_name":"Farm to Table Produce Hub","address":"812 Cypress Way, Denver, CO 80201","telephone":"+1-555-1901","locality_id":1`),
			contentType:          "application/json",
			expectedRequestError: pkgRequest.ErrRequestJSONInvalid,
		},
		{
			name:                 "should return error when request body is text/plain",
			reqBody:              strings.NewReader(`{"cid":"GDJ2SJ3","company_name":"Farm to Table Produce Hub","address":"812 Cypress Way, Denver, CO 80201","telephone":"+1-555-1901","locality_id":1}`),
			contentType:          "text/plain",
			expectedRequestError: pkgRequest.ErrRequestContentTypeNotJSON,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &seller.MockSellerService{}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Update()
			//Act
			target := fmt.Sprintf("/%s", "1")
			request := httptest.NewRequest("DELETE", target, nil)
			request = utils.AddPathParamToRequest(request, "id", "1")
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			expectedCode := http.StatusBadRequest
			require.Equal(t, expectedCode, response.Code)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.Equal(t, 0, srv.Spy.CountUpdateFunc)

		})
	}
}

func TestSellerHandler_Update_ValidateRequestData_Errors(t *testing.T) {

	tests := []struct {
		name             string
		updateRequest    models.UpdateSellerRequest
		expectedBadField string
	}{
		{
			name: "should return error when Cid is present but empty",
			updateRequest: models.UpdateSellerRequest{
				Cid:         &[]string{""}[0],
				CompanyName: &[]string{"Farm to Table Produce Hub"}[0],
				Address:     &[]string{"812 Cypress Way, Denver, CO 80201"}[0],
				Telephone:   &[]string{"+1-555-1901"}[0],
				LocalityId:  &[]int{1}[0],
			},
			expectedBadField: "Cid",
		},
		{
			name: "should return error when CompanyName is present but empty",
			updateRequest: models.UpdateSellerRequest{
				Cid:         &[]string{"1"}[0],
				CompanyName: &[]string{""}[0],
				Address:     &[]string{"812 Cypress Way, Denver, CO 80201"}[0],
				Telephone:   &[]string{"+1-555-1901"}[0],
				LocalityId:  &[]int{1}[0],
			},
			expectedBadField: "CompanyName",
		},
		{
			name: "should return error when Address is present but empty",
			updateRequest: models.UpdateSellerRequest{
				Cid:         &[]string{"1"}[0],
				CompanyName: &[]string{"Farm to Table Produce Hub"}[0],
				Address:     &[]string{""}[0],
				Telephone:   &[]string{"+1-555-1901"}[0],
				LocalityId:  &[]int{1}[0],
			},
			expectedBadField: "Address",
		},
		{
			name: "should return error when Telephone is present but empty",
			updateRequest: models.UpdateSellerRequest{
				Cid:         &[]string{"1"}[0],
				CompanyName: &[]string{"Farm to Table Produce Hub"}[0],
				Address:     &[]string{"812 Cypress Way, Denver, CO 80201"}[0],
				Telephone:   &[]string{""}[0],
				LocalityId:  &[]int{1}[0],
			},
			expectedBadField: "Telephone",
		},
		{
			name: "should return error when LocalityId is present but less to zero",
			updateRequest: models.UpdateSellerRequest{
				Cid:         &[]string{"1"}[0],
				CompanyName: &[]string{"Farm to Table Produce Hub"}[0],
				Address:     &[]string{"812 Cypress Way, Denver, CO 80201"}[0],
				Telephone:   &[]string{"+1-555-1901"}[0],
				LocalityId:  &[]int{-1}[0],
			},
			expectedBadField: "LocalityId",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			srv := &seller.MockSellerService{}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Update()
			//Act
			target := fmt.Sprintf("/%s", "1")
			reqBody, _ := json.Marshal(test.updateRequest)
			request := httptest.NewRequest("PATCH", target, bytes.NewReader(reqBody))
			request = utils.AddPathParamToRequest(request, "id", "1")
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			expectedCode := http.StatusUnprocessableEntity
			require.Equal(t, expectedCode, response.Code)
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.Contains(t, response.Body.String(), test.expectedBadField)
			require.Equal(t, 0, srv.Spy.CountUpdateFunc)
		})
	}
}

func TestSellerHandler_Update_Errors(t *testing.T) {
	updateRequestSuccess := models.UpdateSellerRequest{
		Cid:         &[]string{"1"}[0],
		CompanyName: &[]string{"Farm to Table Produce Hub"}[0],
		Address:     &[]string{"812 Cypress Way, Denver, CO 80201"}[0],
		Telephone:   &[]string{"+1-555-1901"}[0],
		LocalityId:  &[]int{1}[0],
	}
	tests := []struct {
		name          string
		expectedCode  int
		expectedError error
	}{
		{
			name:          "should return 409 conflict error when service returns conflict error",
			expectedCode:  http.StatusConflict,
			expectedError: pkgError.ErrConflict,
		},
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
			srv := &seller.MockSellerService{
				UpdateFunc: func(id int, reqBody *models.UpdateSellerRequest) (models.Seller, error) {
					return models.Seller{}, test.expectedError
				},
			}
			hd := NewSellerHandler(srv)
			hdFunc := hd.Update()
			//Act
			reqBody, _ := json.Marshal(updateRequestSuccess)
			request := httptest.NewRequest("POST", "/1", bytes.NewReader(reqBody))
			request = utils.AddPathParamToRequest(request, "id", "1")
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hdFunc(response, request)

			//Assert
			require.Contains(t, response.Body.String(), test.expectedError.Error())
			require.Equal(t, test.expectedCode, response.Code)
			require.Equal(t, 1, srv.Spy.CountUpdateFunc)
		})
	}
}

func TestSellerHandler_Update_NotFoundError(t *testing.T) {

	//Arrange
	updateRequestSuccess := models.UpdateSellerRequest{
		Cid:         &[]string{"1235F"}[0],
		CompanyName: &[]string{"New Farm to Table Produce Hub"}[0],
	}
	sellerId := 1
	expectedCode := http.StatusNotFound
	expectedBody := fmt.Sprintf(`{"status":"%s", "message": "%s"}`, http.StatusText(http.StatusNotFound), pkgError.WrapErrNotFound("seller", "id", sellerId))
	//Act
	srv := &seller.MockSellerService{
		UpdateFunc: func(id int, reqBody *models.UpdateSellerRequest) (models.Seller, error) {
			return models.Seller{}, pkgError.WrapErrNotFound("seller", "id", id)
		},
	}
	hd := NewSellerHandler(srv)
	hdFunc := hd.Update()
	reqBody, _ := json.Marshal(updateRequestSuccess)
	request := httptest.NewRequest("PATCH", fmt.Sprintf("/%v", sellerId), bytes.NewReader(reqBody))
	request = utils.AddPathParamToRequest(request, "id", fmt.Sprintf("%v", sellerId))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	hdFunc(response, request)

	//Assert
	require.Equal(t, expectedCode, response.Code)
	require.JSONEq(t, expectedBody, response.Body.String())
	require.Equal(t, srv.Spy.CountUpdateFunc, 1)
}

func TestSellerHandler_Update_Success(t *testing.T) {
	currentSeller := models.Seller{
		Id:          1,
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	//Arrange
	updateRequestSuccess := models.UpdateSellerRequest{
		Cid:         &[]string{"1235F"}[0],
		CompanyName: &[]string{"New Farm to Table Produce Hub"}[0],
	}
	sellerId := 1
	expectedCode := http.StatusOK
	expectedSellerUpdated := models.Seller{
		Id:          1,
		Cid:         "1235F",
		CompanyName: "New Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	expectedSellerDocUpdated := expectedSellerUpdated.ModelToDoc()
	expectedJsonSellerDoc, _ := json.Marshal(expectedSellerDocUpdated)
	expectedBody := fmt.Sprintf(`{"data":[%s]}`, string(expectedJsonSellerDoc))
	//Act
	srv := &seller.MockSellerService{
		UpdateFunc: func(id int, reqBody *models.UpdateSellerRequest) (models.Seller, error) {
			updatedSeller := currentSeller
			updatedSeller.Cid = *reqBody.Cid
			updatedSeller.CompanyName = *reqBody.CompanyName
			return updatedSeller, nil
		},
	}
	hd := NewSellerHandler(srv)
	hdFunc := hd.Update()
	reqBody, _ := json.Marshal(updateRequestSuccess)
	request := httptest.NewRequest("PATCH", fmt.Sprintf("/%v", sellerId), bytes.NewReader(reqBody))
	request = utils.AddPathParamToRequest(request, "id", fmt.Sprintf("%v", sellerId))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	hdFunc(response, request)

	//Assert
	require.Equal(t, expectedCode, response.Code)
	require.JSONEq(t, expectedBody, response.Body.String())
	require.Equal(t, srv.Spy.CountUpdateFunc, 1)
}
