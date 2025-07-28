package carrier

import (
	"errors"
	"testing"

	carrierMocks "ProyectoFinal/mocks/carrier"
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/require"
)

func TestCarrierService_Create_Success(t *testing.T) {
	mockRepo := new(carrierMocks.MockCarrierRepository)
	service := NewCarrierService(mockRepo)

	inputCarrier := &models.Carrier{
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

	mockRepo.On("Create", inputCarrier).Return(expectedCarrier, nil)

	result, err := service.Create(inputCarrier)

	require.NoError(t, err)
	require.Equal(t, expectedCarrier, result)
	mockRepo.AssertExpectations(t)
}

func TestCarrierService_Create_RepositoryError_ReturnsError(t *testing.T) {
	mockRepo := new(carrierMocks.MockCarrierRepository)
	service := NewCarrierService(mockRepo)

	inputCarrier := &models.Carrier{
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("Create", inputCarrier).Return(nil, expectedError)

	result, err := service.Create(inputCarrier)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
