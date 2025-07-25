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
