package buyer

import (
	"ProyectoFinal/mocks/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBuyerService_GetById(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, *models.Buyer, error)
	}{
		{
			name: "find_by_id_existent",
			id:   1,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				expectedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockRepo.On("GetById", 1).Return(expectedBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "1", buyer.CardNumberId)
				assert.Equal(t, "Pepito", buyer.FirstName)
				assert.Equal(t, "Perez", buyer.LastName)
			},
		},
		{
			name: "find_by_id_non_existent",
			id:   999,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockRepo.On("GetById", 999).Return(nil, notFoundErr)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "not found")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "999")
			},
		},
		{
			name: "find_by_id_repository_error",
			id:   1,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("GetById", 1).Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.GetById(tt.id)

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBuyerService_GetAll(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, []*models.Buyer, error)
	}{
		{
			name: "find_all",
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				expectedBuyers := []*models.Buyer{
					{
						Id:           1,
						CardNumberId: "1",
						FirstName:    "Pepito",
						LastName:     "Perez",
					},
					{
						Id:           2,
						CardNumberId: "2",
						FirstName:    "Juan",
						LastName:     "Lopez",
					},
					{
						Id:           3,
						CardNumberId: "3",
						FirstName:    "Ana",
						LastName:     "Garcia",
					},
				}
				mockRepo.On("GetAll").Return(expectedBuyers, nil)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyers)
				assert.Len(t, buyers, 3)
				assert.Equal(t, 1, buyers[0].Id)
				assert.Equal(t, "1", buyers[0].CardNumberId)
				assert.Equal(t, "Pepito", buyers[0].FirstName)
				assert.Equal(t, "Perez", buyers[0].LastName)
				assert.Equal(t, 2, buyers[1].Id)
				assert.Equal(t, "2", buyers[1].CardNumberId)
				assert.Equal(t, "Juan", buyers[1].FirstName)
				assert.Equal(t, "Lopez", buyers[1].LastName)
				assert.Equal(t, 3, buyers[2].Id)
				assert.Equal(t, "3", buyers[2].CardNumberId)
				assert.Equal(t, "Ana", buyers[2].FirstName)
				assert.Equal(t, "Garcia", buyers[2].LastName)
			},
		},
		{
			name: "get_all_buyers_repository_error",
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("GetAll").Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
		{
			name: "get_all_buyers_empty_list",
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("GetAll").Return(nil, nil)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.NoError(t, err)
				assert.Nil(t, buyers)
				assert.Len(t, buyers, 0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.GetAll()

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBuyerService_Create(t *testing.T) {
	tests := []struct {
		name           string
		buyer          *models.Buyer
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, *models.Buyer, error)
	}{
		{
			name: "create_ok",
			buyer: &models.Buyer{
				CardNumberId: "1",
				FirstName:    "Pepito",
				LastName:     "Perez",
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				expectedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockRepo.On("Create", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.CardNumberId == "1" &&
						buyer.FirstName == "Pepito" &&
						buyer.LastName == "Perez"
				})).Return(expectedBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "1", buyer.CardNumberId)
				assert.Equal(t, "Pepito", buyer.FirstName)
				assert.Equal(t, "Perez", buyer.LastName)
			},
		},
		{
			name: "create_conflict",
			buyer: &models.Buyer{
				CardNumberId: "1",
				FirstName:    "Juan",
				LastName:     "Lopez",
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				conflictErr := errors.WrapErrConflict("buyer", "card_number_id", "1")
				mockRepo.On("Create", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.CardNumberId == "1" &&
						buyer.FirstName == "Juan" &&
						buyer.LastName == "Lopez"
				})).Return(nil, conflictErr)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "conflict")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "card_number_id")
				assert.Contains(t, err.Error(), "1")
				assert.Contains(t, err.Error(), "already exists")
			},
		},
		{
			name: "create_buyer_repository_error",
			buyer: &models.Buyer{
				CardNumberId: "2",
				FirstName:    "Ana",
				LastName:     "Garcia",
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("Create", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.CardNumberId == "2" &&
						buyer.FirstName == "Ana" &&
						buyer.LastName == "Garcia"
				})).Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.Create(tt.buyer)

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBuyerService_Update(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		buyer          *models.Buyer
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, *models.Buyer, error)
	}{
		{
			name: "update_existent",
			id:   1,
			buyer: &models.Buyer{
				Id:           1,
				CardNumberId: "1",
				FirstName:    "Pepito Updated",
				LastName:     "Perez Updated",
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				updatedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito Updated",
					LastName:     "Perez Updated",
				}
				mockRepo.On("Update", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.Id == 1 &&
						buyer.CardNumberId == "1" &&
						buyer.FirstName == "Pepito Updated" &&
						buyer.LastName == "Perez Updated"
				})).Return(updatedBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "1", buyer.CardNumberId)
				assert.Equal(t, "Pepito Updated", buyer.FirstName)
				assert.Equal(t, "Perez Updated", buyer.LastName)
			},
		},
		{
			name: "update_non_existent",
			id:   999,
			buyer: &models.Buyer{
				Id:           999,
				CardNumberId: "999",
				FirstName:    "Non Existent",
				LastName:     "Buyer",
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockRepo.On("Update", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.Id == 999 &&
						buyer.CardNumberId == "999" &&
						buyer.FirstName == "Non Existent" &&
						buyer.LastName == "Buyer"
				})).Return(nil, notFoundErr)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "not found")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "999")
			},
		},
		{
			name: "update_buyer_duplicate_card_number",
			id:   1,
			buyer: &models.Buyer{
				Id:           1,
				CardNumberId: "2",
				FirstName:    "Pepito",
				LastName:     "Perez",
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				conflictErr := errors.WrapErrConflict("buyer", "card_number_id", "2")
				mockRepo.On("Update", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.Id == 1 &&
						buyer.CardNumberId == "2" &&
						buyer.FirstName == "Pepito" &&
						buyer.LastName == "Perez"
				})).Return(nil, conflictErr)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "conflict")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "card_number_id")
				assert.Contains(t, err.Error(), "2")
				assert.Contains(t, err.Error(), "already exists")
			},
		},
		{
			name: "update_buyer_repository_error",
			id:   1,
			buyer: &models.Buyer{
				Id:           1,
				CardNumberId: "1",
				FirstName:    "Pepito",
				LastName:     "Perez",
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("Update", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.Id == 1 &&
						buyer.CardNumberId == "1" &&
						buyer.FirstName == "Pepito" &&
						buyer.LastName == "Perez"
				})).Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.Update(tt.id, tt.buyer)

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBuyerService_Delete(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, error)
	}{
		{
			name: "delete_ok",
			id:   1,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("Delete", 1).Return(nil)
			},
			validateResult: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "delete_non_existent",
			id:   999,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockRepo.On("Delete", 999).Return(notFoundErr)
			},
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "not found")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "999")
			},
		},
		{
			name: "delete_buyer_repository_error",
			id:   1,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("Delete", 1).Return(errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
		{
			name: "delete_buyer_foreign_key_constraint",
			id:   1,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				constraintErr := errors.WrapErrConflict("buyer", "id", 1)
				mockRepo.On("Delete", 1).Return(constraintErr)
			},
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "conflict")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "1")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			err := service.Delete(tt.id)

			// Assert
			tt.validateResult(t, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBuyerService_GetByIdWithOrderCount(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, *models.BuyerWithOrderCount, error)
	}{
		{
			name: "get_buyer_by_id_with_order_count_success",
			id:   1,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				expectedBuyer := &models.BuyerWithOrderCount{
					Id:                  1,
					CardNumberId:        "1",
					FirstName:           "Pepito",
					LastName:            "Perez",
					PurchaseOrdersCount: 5,
				}
				mockRepo.On("GetByIdWithOrderCount", 1).Return(expectedBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "1", buyer.CardNumberId)
				assert.Equal(t, "Pepito", buyer.FirstName)
				assert.Equal(t, "Perez", buyer.LastName)
				assert.Equal(t, 5, buyer.PurchaseOrdersCount)
			},
		},
		{
			name: "get_buyer_by_id_with_zero_order_count",
			id:   2,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				expectedBuyer := &models.BuyerWithOrderCount{
					Id:                  2,
					CardNumberId:        "2",
					FirstName:           "Juan",
					LastName:            "Lopez",
					PurchaseOrdersCount: 0,
				}
				mockRepo.On("GetByIdWithOrderCount", 2).Return(expectedBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 2, buyer.Id)
				assert.Equal(t, "Juan", buyer.FirstName)
				assert.Equal(t, 0, buyer.PurchaseOrdersCount)
			},
		},
		{
			name: "get_buyer_by_id_with_order_count_non_existent",
			id:   999,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockRepo.On("GetByIdWithOrderCount", 999).Return(nil, notFoundErr)
			},
			validateResult: func(t *testing.T, buyer *models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "not found")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "999")
			},
		},
		{
			name: "get_buyer_by_id_with_order_count_repository_error",
			id:   1,
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("GetByIdWithOrderCount", 1).Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, buyer *models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.GetByIdWithOrderCount(tt.id)

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBuyerService_GetAllWithOrderCount(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, []*models.BuyerWithOrderCount, error)
	}{
		{
			name: "get_all_buyers_with_order_count_success",
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				expectedBuyers := []*models.BuyerWithOrderCount{
					{
						Id:                  1,
						CardNumberId:        "1",
						FirstName:           "Pepito",
						LastName:            "Perez",
						PurchaseOrdersCount: 5,
					},
					{
						Id:                  2,
						CardNumberId:        "2",
						FirstName:           "Juan",
						LastName:            "Lopez",
						PurchaseOrdersCount: 3,
					},
					{
						Id:                  3,
						CardNumberId:        "3",
						FirstName:           "Ana",
						LastName:            "Garcia",
						PurchaseOrdersCount: 0,
					},
				}
				mockRepo.On("GetAllWithOrderCount").Return(expectedBuyers, nil)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyers)
				assert.Len(t, buyers, 3)
				assert.Equal(t, 1, buyers[0].Id)
				assert.Equal(t, "1", buyers[0].CardNumberId)
				assert.Equal(t, "Pepito", buyers[0].FirstName)
				assert.Equal(t, "Perez", buyers[0].LastName)
				assert.Equal(t, 5, buyers[0].PurchaseOrdersCount)
				assert.Equal(t, 3, buyers[2].Id)
				assert.Equal(t, "Ana", buyers[2].FirstName)
				assert.Equal(t, 0, buyers[2].PurchaseOrdersCount)
			},
		},
		{
			name: "get_all_buyers_with_order_count_empty_list",
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("GetAllWithOrderCount").Return(nil, nil)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.Nil(t, buyers)
				assert.Len(t, buyers, 0)
			},
		},
		{
			name: "get_all_buyers_with_order_count_repository_error",
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				mockRepo.On("GetAllWithOrderCount").Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.GetAllWithOrderCount()

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBuyerService_PatchUpdate(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		updateDTO      *models.BuyerUpdateDTO
		setupMock      func(*buyer.MockBuyerRepository)
		validateResult func(*testing.T, *models.Buyer, error)
	}{
		{
			name: "patch_update_success_all_fields",
			id:   1,
			updateDTO: &models.BuyerUpdateDTO{
				CardNumberId: stringPtr("1-updated"),
				FirstName:    stringPtr("Pepito Updated"),
				LastName:     stringPtr("Perez Updated"),
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				existingBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockRepo.On("GetById", 1).Return(existingBuyer, nil)
				updatedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1-updated",
					FirstName:    "Pepito Updated",
					LastName:     "Perez Updated",
				}
				mockRepo.On("Update", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.Id == 1 &&
						buyer.CardNumberId == "1-updated" &&
						buyer.FirstName == "Pepito Updated" &&
						buyer.LastName == "Perez Updated"
				})).Return(updatedBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "1-updated", buyer.CardNumberId)
				assert.Equal(t, "Pepito Updated", buyer.FirstName)
				assert.Equal(t, "Perez Updated", buyer.LastName)
			},
		},
		{
			name: "patch_update_success_partial_fields",
			id:   1,
			updateDTO: &models.BuyerUpdateDTO{
				FirstName: stringPtr("Only First Name Updated"),
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				existingBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockRepo.On("GetById", 1).Return(existingBuyer, nil)
				updatedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Only First Name Updated",
					LastName:     "Perez",
				}
				mockRepo.On("Update", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.Id == 1 &&
						buyer.CardNumberId == "1" &&
						buyer.FirstName == "Only First Name Updated" &&
						buyer.LastName == "Perez"
				})).Return(updatedBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "1", buyer.CardNumberId)
				assert.Equal(t, "Only First Name Updated", buyer.FirstName)
				assert.Equal(t, "Perez", buyer.LastName)
			},
		},
		{
			name: "patch_update_buyer_non_existent",
			id:   999,
			updateDTO: &models.BuyerUpdateDTO{
				FirstName: stringPtr("New Name"),
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockRepo.On("GetById", 999).Return(nil, notFoundErr)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "not found")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "999")
			},
		},
		{
			name: "patch_update_no_fields_provided",
			id:   1,
			updateDTO: &models.BuyerUpdateDTO{
				CardNumberId: nil,
				FirstName:    nil,
				LastName:     nil,
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				existingBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockRepo.On("GetById", 1).Return(existingBuyer, nil)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "unprocessable entity")
				assert.Contains(t, err.Error(), "no fields provided for update")
			},
		},
		{
			name: "patch_update_update_repository_error",
			id:   1,
			updateDTO: &models.BuyerUpdateDTO{
				FirstName: stringPtr("New Name"),
			},
			setupMock: func(mockRepo *buyer.MockBuyerRepository) {
				existingBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockRepo.On("GetById", 1).Return(existingBuyer, nil)
				mockRepo.On("Update", mock.AnythingOfType("*models.Buyer")).Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Equal(t, errors.ErrGeneral, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(buyer.MockBuyerRepository)
			service := NewBuyerService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.PatchUpdate(tt.id, tt.updateDTO)

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
