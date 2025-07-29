package locality

import (
	"ProyectoFinal/mocks/locality"
	"ProyectoFinal/pkg/models"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLocalityService_Create(t *testing.T) {
	localityToCreate := models.Locality{
		LocalityName: "test",
		Province: models.Province{
			ProvinceName: "Lombardy",
			Country: models.Country{
				CountryName: "Italy",
			},
		},
	}
	expectedLocalityCreated := models.Locality{
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
	tests := []struct {
		name       string
		mockReturn models.Locality
		mockError  error
		assertFunc func(t *testing.T, locality models.Locality, err error)
	}{
		{
			name:       "should return err when repository return error",
			mockReturn: models.Locality{},
			mockError:  errors.New("error repository"),
			assertFunc: func(t *testing.T, locality models.Locality, err error) {
				expectedError := errors.New("error repository")
				require.Equal(t, models.Locality{}, locality)
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name:       "should create an locality",
			mockReturn: expectedLocalityCreated,
			mockError:  nil,
			assertFunc: func(t *testing.T, locality models.Locality, err error) {
				require.Equal(t, expectedLocalityCreated, locality)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(locality.MockLocalityRepository)
			mockRepo.On("Create", localityToCreate).Return(test.mockReturn, test.mockError)
			srv := NewLocalityService(mockRepo)

			//Act
			result, err := srv.Create(localityToCreate)
			//Assert
			test.assertFunc(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
func TestLocalityService_GetSellersByIdLocality(t *testing.T) {
	idLocality := 1
	expectedSellersByLocalities := []models.SellersByLocalityReport{
		{
			LocalityId:   1,
			LocalityName: "test",
			SellersCount: 3,
		},
		{
			LocalityId:   2,
			LocalityName: "test 2",
			SellersCount: 12,
		},
	}
	tests := []struct {
		name       string
		mockReturn models.SellersByLocalityReport
		mockError  error
		assertFunc func(t *testing.T, sellersByLocality models.SellersByLocalityReport, err error)
	}{
		{
			name:       "should return err when repository return error",
			mockReturn: models.SellersByLocalityReport{},
			mockError:  errors.New("error repository"),
			assertFunc: func(t *testing.T, sellersByLocality models.SellersByLocalityReport, err error) {
				expectedError := errors.New("error repository")
				require.Equal(t, models.SellersByLocalityReport{}, sellersByLocality)
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name:       "should get seller by locality id",
			mockReturn: expectedSellersByLocalities[0],
			mockError:  nil,
			assertFunc: func(t *testing.T, sellersByLocality models.SellersByLocalityReport, err error) {
				require.Equal(t, expectedSellersByLocalities[0], sellersByLocality)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(locality.MockLocalityRepository)
			mockRepo.On("GetSellersByIdLocality", idLocality).Return(test.mockReturn, test.mockError)
			srv := NewLocalityService(mockRepo)

			//Act
			result, err := srv.GetSellersByIdLocality(idLocality)
			//Assert
			test.assertFunc(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
func TestLocalityService_GetSellersByLocalities(t *testing.T) {
	expectedSellersByLocalities := []models.SellersByLocalityReport{
		{
			LocalityId:   1,
			LocalityName: "test",
			SellersCount: 3,
		},
		{
			LocalityId:   2,
			LocalityName: "test 2",
			SellersCount: 12,
		},
	}
	tests := []struct {
		name       string
		mockReturn []models.SellersByLocalityReport
		mockError  error
		assertFunc func(t *testing.T, sellersByLocality []models.SellersByLocalityReport, err error)
	}{
		{
			name:       "should return err when repository return error",
			mockReturn: nil,
			mockError:  errors.New("error repository"),
			assertFunc: func(t *testing.T, sellersByLocality []models.SellersByLocalityReport, err error) {
				expectedError := errors.New("error repository")
				require.Nil(t, sellersByLocality)
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name:       "should get all sellers by locality",
			mockReturn: expectedSellersByLocalities,
			mockError:  nil,
			assertFunc: func(t *testing.T, sellersByLocality []models.SellersByLocalityReport, err error) {
				require.Equal(t, expectedSellersByLocalities, sellersByLocality)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(locality.MockLocalityRepository)
			mockRepo.On("GetSellersByLocalities").Return(test.mockReturn, test.mockError)
			srv := NewLocalityService(mockRepo)

			//Act
			result, err := srv.GetSellersByLocalities()
			//Assert
			test.assertFunc(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestLocalityService_ReportCarriersByLocality(t *testing.T) {
	expectedCarriersByLocality := []models.CarrierReport{
		{
			LocalityId:    1,
			LocalityName:  "test",
			CarriersCount: 3,
		},
		{
			LocalityId:    2,
			LocalityName:  "test 2",
			CarriersCount: 12,
		},
	}
	tests := []struct {
		name       string
		idLocality *int
		mockReturn []models.CarrierReport
		mockError  error
		assertFunc func(t *testing.T, carriersByLocality []models.CarrierReport, err error)
	}{
		{
			name:       "should return err when repository return error",
			mockReturn: nil,
			mockError:  errors.New("error repository"),
			assertFunc: func(t *testing.T, carriersByLocality []models.CarrierReport, err error) {
				expectedError := errors.New("error repository")
				require.Nil(t, carriersByLocality)
				require.Error(t, err)
				require.Equal(t, expectedError, err)
			},
		},
		{
			name:       "should get carrier by locality id",
			mockReturn: []models.CarrierReport{expectedCarriersByLocality[0]},
			mockError:  nil,
			idLocality: &[]int{1}[0],
			assertFunc: func(t *testing.T, carriersByLocality []models.CarrierReport, err error) {
				require.Equal(t, expectedCarriersByLocality[0], carriersByLocality[0])
				require.NoError(t, err)
			},
		},
		{
			name:       "should get all carriers by locality",
			mockReturn: expectedCarriersByLocality,
			mockError:  nil,
			assertFunc: func(t *testing.T, carriersByLocality []models.CarrierReport, err error) {
				require.Equal(t, expectedCarriersByLocality, carriersByLocality)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockRepo := new(locality.MockLocalityRepository)
			mockRepo.On("ReportCarriersByLocality", test.idLocality).Return(test.mockReturn, test.mockError)
			srv := NewLocalityService(mockRepo)

			//Act
			result, err := srv.ReportCarriersByLocality(test.idLocality)
			//Assert
			test.assertFunc(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
