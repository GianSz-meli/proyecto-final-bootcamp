package handler_test

import (
	"ProyectoFinal/internal/handler"
	"ProyectoFinal/mocks"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"context"
	"encoding/json"

	// "errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	const validProductJSON = `{
    "product_code": "A1000",
    "description": "Coca Cola 2L",
    "width": 5,
    "height": 20,
    "length": 12,
    "net_weight": 1.9,
    "expiration_rate": 15,
    "recommended_freezing_temperature": 10,
    "freezing_rate": 10,
    "product_type_id": 83,
    "seller_id": 7
}`
	const invalidProductJSON = `{
    "description": "Coca Cola 2L",
    "width": 5,
    "height": 20,
    "length": 12,
    "net_weight": 1.9,
    "expiration_rate": 15,
    "recommended_freezing_temperature": 10,
    "freezing_rate": 10,
    "product_type_id": 83,
    "seller_id": 7
}`
	sellerid := 7
	productCode := "A1000"
	description := "Coca Cola 2L"
	width := 5.00
	height := 20.00
	length := 12.00
	netWeight := 1.9
	expirationRate := 15.00
	temperature := 10.00
	freezingRate := 10.00
	productTypeId := 83
	sellerID := sellerid

	inputProduct := models.Product{
		ProductCode:    productCode,
		Description:    description,
		Width:          width,
		Height:         height,
		Length:         length,
		NetWeight:      netWeight,
		ExpirationRate: expirationRate,
		Temperature:    float32(temperature),
		FreezingRate:   freezingRate,
		ProductTypeID:  productTypeId,
		SellerID:       &sellerID,
	}

	returnedProduct := models.Product{
		ID:             1,
		ProductCode:    productCode,
		Description:    description,
		Width:          width,
		Height:         height,
		Length:         length,
		NetWeight:      netWeight,
		ExpirationRate: expirationRate,
		Temperature:    float32(temperature),
		FreezingRate:   freezingRate,
		ProductTypeID:  productTypeId,
		SellerID:       &sellerID,
	}

	t.Run("create_ok", func(t *testing.T) {
		mockService := &mocks.MockProductService{}
		mockService.On("CreateProduct", inputProduct).Return(returnedProduct, nil)
		hd := handler.NewProductHandler(mockService)

		request := httptest.NewRequest("POST", "/products", strings.NewReader(validProductJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		//act
		hd.CreateProduct(response, request)

		//assert
		require.Equal(t, http.StatusCreated, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var resp struct {
			Data models.ProductDoc `json:"data"`
		}

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		require.NoError(t, err)
		require.Equal(t, returnedProduct, resp.Data.DocToModel())
		mockService.AssertExpectations(t)

	})

	t.Run("create_fail", func(t *testing.T) {

		mockServiceFail := &mocks.MockProductService{}
		hd := handler.NewProductHandler(mockServiceFail)

		request := httptest.NewRequest("POST", "/products", strings.NewReader(invalidProductJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProduct(response, request)

		require.Equal(t, http.StatusUnprocessableEntity, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		mockServiceFail.AssertNotCalled(t, "CreateProduct", mock.Anything)
	})

	t.Run("conflict", func(t *testing.T) {

		mockServiceFail := &mocks.MockProductService{}
		mockServiceFail.On("CreateProduct", mock.Anything).Return(models.Product{}, pkgErrors.ErrConflict)
		hd := handler.NewProductHandler(mockServiceFail)

		request := httptest.NewRequest("POST", "/products", strings.NewReader(validProductJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProduct(response, request)

		require.Equal(t, http.StatusConflict, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		mockServiceFail.AssertExpectations(t)

	})
}

func TestFindProducts(t *testing.T) {
	prod1 := models.Product{
		ID:             1,
		ProductCode:    "A123",
		Description:    "Caja de manzanas",
		Width:          40.0,
		Height:         25.0,
		Length:         60.0,
		NetWeight:      15.0,
		ExpirationRate: 0.05,
		Temperature:    4.0,
		FreezingRate:   0.02,
		ProductTypeID:  2,
		SellerID:       nil,
	}
	sellerID := 17
	prod2 := models.Product{
		ID:             2,
		ProductCode:    "B456",
		Description:    "Botella de jugo",
		Width:          8.0,
		Height:         30.0,
		Length:         8.0,
		NetWeight:      1.2,
		ExpirationRate: 0.01,
		Temperature:    6.5,
		FreezingRate:   0.0,
		ProductTypeID:  3,
		SellerID:       &sellerID,
	}
	prods := map[int]models.Product{
		1: prod1,
		2: prod2,
	}

	t.Run("find_all", func(t *testing.T) {
		mockService := &mocks.MockProductService{}
		mockService.On("FindAllProducts").Return(prods, nil)
		hd := handler.NewProductHandler(mockService)

		request := httptest.NewRequest("GET", "/products", nil)
		response := httptest.NewRecorder()

		hd.FindAllProducts(response, request)

		require.Equal(t, http.StatusOK, response.Code)

	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		mockService := &mocks.MockProductService{}
		mockService.On("FindProductsById", 3).Return(models.Product{}, pkgErrors.ErrNotFound)
		hd := handler.NewProductHandler(mockService)

		request := httptest.NewRequest("GET", "/products/3", nil)

		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "3")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, routeCtx))

		response := httptest.NewRecorder()

		hd.FindProductsById(response, request)

		require.Equal(t, http.StatusNotFound, response.Code)
		require.Contains(t, response.Body.String(), pkgErrors.ErrNotFound.Error())
		mockService.AssertExpectations(t)
	})

	t.Run("find_by_id_success", func(t *testing.T) {

		mockService := &mocks.MockProductService{}
		mockService.On("FindProductsById", 1).Return(prod1, nil)
		hd := handler.NewProductHandler(mockService)

		req := httptest.NewRequest("GET", "/products/1", nil)

		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

		res := httptest.NewRecorder()
		hd.FindProductsById(res, req)

		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		type Response struct {
			Data models.ProductDoc `json:"data"`
		}
		var got Response
		err := json.Unmarshal(res.Body.Bytes(), &got)
		require.NoError(t, err)
		require.Equal(t, prod1.ModelToDoc(), got.Data)

		mockService.AssertExpectations(t)
	})

}

func TestUpdateProduct(t *testing.T) {
	const validProductJSON = `{
		"product_code": "A1000",
		"description": "Coca Cola 2L",
		"width": 5,
		"height": 20,
		"length": 12,
		"net_weight": 1.9,
		"expiration_rate": 15,
		"recommended_freezing_temperature": 10,
		"freezing_rate": 10,
		"product_type_id": 83,
		"seller_id": 7
	}`
	prodinput := models.Product{
		ProductCode:    "A123",
		Description:    "Caja de manzanas",
		Width:          40.0,
		Height:         25.0,
		Length:         60.0,
		NetWeight:      15.0,
		ExpirationRate: 0.05,
		Temperature:    4.0,
		FreezingRate:   0.02,
		ProductTypeID:  2,
		SellerID:       nil,
	}
	prodoutput := models.Product{
		ID:             1,
		ProductCode:    "A123",
		Description:    "Caja de manzanas",
		Width:          40.0,
		Height:         25.0,
		Length:         60.0,
		NetWeight:      15.0,
		ExpirationRate: 0.05,
		Temperature:    4.0,
		FreezingRate:   0.02,
		ProductTypeID:  2,
		SellerID:       nil,
	}

	t.Run("update_ok", func(t *testing.T) {
		mockService := &mocks.MockProductService{}
		mockService.On("UpdateProduct", 1, prodinput).Return(prodoutput, nil)
		hd := handler.NewProductHandler(mockService)

		request := httptest.NewRequest("PUT", "/products/1", strings.NewReader(validProductJSON))
		request.Header.Set("Content-Type", "application/json")

		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, routeCtx))

		response := httptest.NewRecorder()

		hd.UpdateProduct(response, request)

		require.Equal(t, http.StatusCreated, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))

	})
	t.Run("update_non_existent", func(t *testing.T) {

		mockService := &mocks.MockProductService{}
		mockService.On("UpdateProduct", 1, mock.AnythingOfType("models.ProductDocUpdate")).Return(models.Product{}, pkgErrors.ErrNotFound)
		hd := handler.NewProductHandler(mockService)

		request := httptest.NewRequest("PUT", "/products/1", strings.NewReader(validProductJSON))
		request.Header.Set("Content-Type", "application/json")

		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, routeCtx))

		response := httptest.NewRecorder()

		hd.UpdateProduct(response, request)

		require.Equal(t, http.StatusNotFound, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		mockService.AssertExpectations(t)
	})

}
