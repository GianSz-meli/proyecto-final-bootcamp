package seller

import (
	"ProyectoFinal/mocks/seller"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSellerService_Create(t *testing.T) {
	sellerToCreate := models.Seller{
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	tests := []struct {
		name       string
		mockReturn models.Seller
		mockError  error
		assertFunc func(t *testing.T, seller models.Seller, err error)
	}{
		{
			name:       "should return err when repository return error",
			mockReturn: models.Seller{},
			mockError:  errors.New("error repository"),
			assertFunc: func(t *testing.T, seller models.Seller, err error) {
				expectedError := errors.New("error repository")
				require.Equal(t, models.Seller{}, seller)
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name: "should create an seller",
			mockReturn: models.Seller{
				Id:          1,
				Cid:         "1",
				CompanyName: "Farm to Table Produce Hub",
				Address:     "812 Cypress Way, Denver, CO 80201",
				Telephone:   "+1-555-1901",
				LocalityId:  1,
			},
			mockError: nil,
			assertFunc: func(t *testing.T, seller models.Seller, err error) {
				expectedSellerCreated := models.Seller{
					Id:          1,
					Cid:         "1",
					CompanyName: "Farm to Table Produce Hub",
					Address:     "812 Cypress Way, Denver, CO 80201",
					Telephone:   "+1-555-1901",
					LocalityId:  1,
				}
				require.Equal(t, expectedSellerCreated, seller)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(seller.MockSellerRepository)
			mockRepo.On("Create", sellerToCreate).Return(test.mockReturn, test.mockError)
			srv := NewSellerService(mockRepo)

			//Act
			result, err := srv.Create(sellerToCreate)
			//Assert
			test.assertFunc(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
func TestSellerService_GetAll(t *testing.T) {
	expectedSellers := []models.Seller{
		{
			Id:          1,
			Cid:         "1",
			CompanyName: "Farm to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  1,
		},
		{
			Id:          2,
			Cid:         "2",
			CompanyName: "Farm Two to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  2,
		},
	}
	tests := []struct {
		name       string
		mockReturn []models.Seller
		mockError  error
		assertFunc func(t *testing.T, sellers []models.Seller, err error)
	}{
		{
			name:       "should return err when repository return error",
			mockReturn: nil,
			mockError:  errors.New("error repository"),
			assertFunc: func(t *testing.T, sellers []models.Seller, err error) {
				expectedError := errors.New("error repository")
				require.Nil(t, sellers)
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name:       "should get all users",
			mockReturn: expectedSellers,
			mockError:  nil,
			assertFunc: func(t *testing.T, sellers []models.Seller, err error) {
				require.Equal(t, expectedSellers, sellers)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(seller.MockSellerRepository)
			mockRepo.On("GetAll").Return(test.mockReturn, test.mockError)
			srv := NewSellerService(mockRepo)

			//Act
			result, err := srv.GetAll()
			//Assert
			test.assertFunc(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
func TestSellerService_GetById(t *testing.T) {
	sellerId := 1
	expectedSeller := models.Seller{
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	tests := []struct {
		name       string
		mockReturn *models.Seller
		mockError  error
		assertFunc func(t *testing.T, seller models.Seller, err error)
	}{
		{
			name:       "should return err when repository return error",
			mockReturn: nil,
			mockError:  errors.New("error repository"),
			assertFunc: func(t *testing.T, seller models.Seller, err error) {
				expectedError := errors.New("error repository")
				require.Equal(t, models.Seller{}, seller)
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name:       "should return not found err when seller by id not exist",
			mockReturn: nil,
			mockError:  pkgErrors.WrapErrNotFound("Seller", "id", sellerId),
			assertFunc: func(t *testing.T, seller models.Seller, err error) {
				require.Equal(t, models.Seller{}, seller)
				require.Error(t, err)
				require.ErrorIs(t, err, pkgErrors.ErrNotFound)
			},
		},
		{
			name:       "should get seller by id",
			mockReturn: &expectedSeller,
			mockError:  nil,
			assertFunc: func(t *testing.T, seller models.Seller, err error) {
				require.Equal(t, expectedSeller, seller)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(seller.MockSellerRepository)
			mockRepo.On("GetById", sellerId).Return(test.mockReturn, test.mockError)
			srv := NewSellerService(mockRepo)

			//Act
			result, err := srv.GetById(sellerId)
			//Assert
			test.assertFunc(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
func TestSellerService_Delete(t *testing.T) {
	sellerId := 1
	tests := []struct {
		name       string
		mockError  error
		assertFunc func(t *testing.T, err error)
	}{
		{
			name:      "should return err when repository return error",
			mockError: errors.New("error repository"),
			assertFunc: func(t *testing.T, err error) {
				expectedError := errors.New("error repository")
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name:      "should return not found err when seller by id not exist",
			mockError: pkgErrors.WrapErrNotFound("Seller", "id", sellerId),
			assertFunc: func(t *testing.T, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, pkgErrors.ErrNotFound)
			},
		},
		{
			name:      "should delete seller by id",
			mockError: nil,
			assertFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(seller.MockSellerRepository)
			mockRepo.On("Delete", sellerId).Return(test.mockError)
			srv := NewSellerService(mockRepo)

			//Act
			err := srv.Delete(sellerId)
			//Assert
			test.assertFunc(t, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestSellerService_Update_GetById_NotFound(t *testing.T) {
	//Arrange
	mockRepository := new(seller.MockSellerRepository)
	updateRequest := models.UpdateSellerRequest{
		Cid:         &[]string{"1235F"}[0],
		CompanyName: &[]string{"New Farm to Table Produce Hub"}[0],
	}
	sellerId := 1
	mockRepository.On("GetById", sellerId).Return(nil, pkgErrors.WrapErrNotFound("domain", "id", sellerId))
	//Act
	srv := NewSellerService(mockRepository)
	sellerUpdated, err := srv.Update(sellerId, &updateRequest)

	//Assert
	require.Equal(t, models.Seller{}, sellerUpdated)
	require.ErrorIs(t, err, pkgErrors.ErrNotFound)
	mockRepository.AssertExpectations(t)
}
func TestSellerService_Update_UpdateFields_Error(t *testing.T) {
	//Arrange
	currentSeller := models.Seller{
		Id:          1,
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	updateRequest := models.UpdateSellerRequest{}
	mockRepository := new(seller.MockSellerRepository)
	sellerId := 1
	mockRepository.On("GetById", sellerId).Return(&currentSeller, nil)

	//Act
	srv := NewSellerService(mockRepository)
	sellerUpdated, err := srv.Update(sellerId, &updateRequest)

	//Assert
	require.Equal(t, models.Seller{}, sellerUpdated)
	require.ErrorIs(t, err, pkgErrors.ErrUnprocessableEntity)
	mockRepository.AssertExpectations(t)

}
func TestSellerService_Update_Error(t *testing.T) {
	currentSeller := models.Seller{
		Id:          1,
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	updateRequest := models.UpdateSellerRequest{
		Cid:         &[]string{"1235F"}[0],
		CompanyName: &[]string{"New Farm to Table Produce Hub"}[0],
	}
	sellerId := 1
	expectedSellerUpdated := &models.Seller{
		Id:          currentSeller.Id,
		Cid:         *updateRequest.Cid,
		CompanyName: *updateRequest.CompanyName,
		Address:     currentSeller.Address,
		Telephone:   currentSeller.Telephone,
		LocalityId:  currentSeller.LocalityId,
	}
	mockRepository := new(seller.MockSellerRepository)
	mockRepository.On("GetById", sellerId).Return(&currentSeller, nil)
	mockRepository.On("Update", mock.MatchedBy(func(s *models.Seller) bool {
		return s.Id == expectedSellerUpdated.Id
	})).Return(models.Seller{}, errors.New("an error occurs"))

	//Act
	srv := NewSellerService(mockRepository)
	sellerUpdated, err := srv.Update(sellerId, &updateRequest)

	//Assert
	require.Equal(t, models.Seller{}, sellerUpdated)
	require.Error(t, err)
	mockRepository.AssertExpectations(t)

}
func TestSellerService_Update_Success(t *testing.T) {
	//Arrange
	currentSeller := models.Seller{
		Id:          1,
		Cid:         "1",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	updateRequest := models.UpdateSellerRequest{
		Cid:         &[]string{"1235F"}[0],
		CompanyName: &[]string{"New Farm to Table Produce Hub"}[0],
	}
	sellerId := 1
	expectedSellerUpdated := models.Seller{
		Id:          currentSeller.Id,
		Cid:         *updateRequest.Cid,
		CompanyName: *updateRequest.CompanyName,
		Address:     currentSeller.Address,
		Telephone:   currentSeller.Telephone,
		LocalityId:  currentSeller.LocalityId,
	}
	mockRepository := new(seller.MockSellerRepository)
	mockRepository.On("GetById", sellerId).Return(&currentSeller, nil)
	mockRepository.On("Update", &expectedSellerUpdated).Return(expectedSellerUpdated, nil)

	//Act
	srv := NewSellerService(mockRepository)
	sellerUpdated, err := srv.Update(sellerId, &updateRequest)

	//Assert
	require.Equal(t, expectedSellerUpdated, sellerUpdated)
	require.NoError(t, err)
	mockRepository.AssertExpectations(t)

}
