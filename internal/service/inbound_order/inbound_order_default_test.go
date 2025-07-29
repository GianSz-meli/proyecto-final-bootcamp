package inbound_order

import (
	"ProyectoFinal/mocks"
	"ProyectoFinal/mocks/inbound_order"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_Create(t *testing.T) {
	tests := []struct {
		name                  string
		inboundOrder          models.InboundOrder
		setupInboundOrderMock func(*inbound_order.MockInboundOrderRepository)
		setupEmployeeMock     func(*mocks.MockEmployeeRepository)
		validateResult        func(*testing.T, models.InboundOrder, error)
	}{
		{
			name: "create_success",
			inboundOrder: models.InboundOrder{
				OrderDate:      time.Date(2023, 10, 15, 10, 30, 0, 0, time.UTC),
				OrderNumber:    "ORD001",
				EmployeeID:     1,
				ProductBatchID: 1,
				WarehouseID:    1,
			},
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				mockRepo.On("Create", mock.AnythingOfType("*models.InboundOrder")).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(0).(*models.InboundOrder)
					arg.ID = 1
				})
			},
			setupEmployeeMock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				expectedEmployee := models.Employee{
					ID:           1,
					CardNumberID: "EMP001",
					FirstName:    "John",
					LastName:     "Doe",
					WarehouseID:  &[]int{1}[0],
				}
				mockEmployeeRepo.On("GetById", 1).Return(expectedEmployee, nil)
			},
			validateResult: func(t *testing.T, result models.InboundOrder, err error) {
				require.NoError(t, err)
				assert.Equal(t, 1, result.ID)
				assert.Equal(t, "ORD001", result.OrderNumber)
				assert.Equal(t, 1, result.EmployeeID)
				assert.Equal(t, 1, result.ProductBatchID)
				assert.Equal(t, 1, result.WarehouseID)
			},
		},
		{
			name: "create_employee_not_found",
			inboundOrder: models.InboundOrder{
				OrderDate:      time.Date(2023, 10, 15, 10, 30, 0, 0, time.UTC),
				OrderNumber:    "ORD001",
				EmployeeID:     999,
				ProductBatchID: 1,
				WarehouseID:    1,
			},
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				// This mock should not be called
			},
			setupEmployeeMock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				notFoundErr := errors.WrapErrNotFound("employee", "id", 999)
				mockEmployeeRepo.On("GetById", 999).Return(models.Employee{}, notFoundErr)
			},
			validateResult: func(t *testing.T, result models.InboundOrder, err error) {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "conflict")
				assert.Contains(t, err.Error(), "employees")
				assert.Equal(t, models.InboundOrder{}, result)
			},
		},
		{
			name: "create_employee_repository_error",
			inboundOrder: models.InboundOrder{
				OrderDate:      time.Date(2023, 10, 15, 10, 30, 0, 0, time.UTC),
				OrderNumber:    "ORD001",
				EmployeeID:     1,
				ProductBatchID: 1,
				WarehouseID:    1,
			},
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {

			},
			setupEmployeeMock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.On("GetById", 1).Return(models.Employee{}, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, result models.InboundOrder, err error) {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "conflict")
				assert.Contains(t, err.Error(), "employees")
				assert.Equal(t, models.InboundOrder{}, result)
			},
		},
		{
			name: "create_inbound_order_repository_error",
			inboundOrder: models.InboundOrder{
				OrderDate:      time.Date(2023, 10, 15, 10, 30, 0, 0, time.UTC),
				OrderNumber:    "ORD001",
				EmployeeID:     1,
				ProductBatchID: 1,
				WarehouseID:    1,
			},
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				mockRepo.On("Create", mock.AnythingOfType("*models.InboundOrder")).Return(errors.ErrGeneral)
			},
			setupEmployeeMock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				expectedEmployee := models.Employee{
					ID:           1,
					CardNumberID: "EMP001",
					FirstName:    "John",
					LastName:     "Doe",
					WarehouseID:  &[]int{1}[0],
				}
				mockEmployeeRepo.On("GetById", 1).Return(expectedEmployee, nil)
			},
			validateResult: func(t *testing.T, result models.InboundOrder, err error) {
				require.Error(t, err)
				assert.Equal(t, errors.ErrGeneral, err)
				assert.Equal(t, models.InboundOrder{}, result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockInboundOrderRepo := new(inbound_order.MockInboundOrderRepository)
			mockEmployeeRepo := new(mocks.MockEmployeeRepository)
			service := NewService(mockInboundOrderRepo, mockEmployeeRepo)
			tt.setupInboundOrderMock(mockInboundOrderRepo)
			tt.setupEmployeeMock(mockEmployeeRepo)

			// Act
			result, err := service.Create(tt.inboundOrder)

			// Assert
			tt.validateResult(t, result, err)
			mockInboundOrderRepo.AssertExpectations(t)
			mockEmployeeRepo.AssertExpectations(t)
		})
	}
}

func TestService_GetEmployeeInboundOrdersReportByEmployeeId(t *testing.T) {
	tests := []struct {
		name                  string
		employeeId            int
		setupInboundOrderMock func(*inbound_order.MockInboundOrderRepository)
		validateResult        func(*testing.T, models.EmployeeInboundOrdersReport, error)
	}{
		{
			name:       "get_report_by_employee_id_success",
			employeeId: 1,
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				expectedReport := models.EmployeeInboundOrdersReport{
					ID:                 1,
					CardNumberID:       "EMP001",
					FirstName:          "John",
					LastName:           "Doe",
					WarehouseID:        1,
					InboundOrdersCount: 5,
				}
				mockRepo.On("GetEmployeeInboundOrdersReportByEmployeeId", 1).Return(expectedReport, nil)
			},
			validateResult: func(t *testing.T, result models.EmployeeInboundOrdersReport, err error) {
				require.NoError(t, err)
				assert.Equal(t, 1, result.ID)
				assert.Equal(t, "EMP001", result.CardNumberID)
				assert.Equal(t, "John", result.FirstName)
				assert.Equal(t, "Doe", result.LastName)
				assert.Equal(t, 1, result.WarehouseID)
				assert.Equal(t, 5, result.InboundOrdersCount)
			},
		},
		{
			name:       "get_report_by_employee_id_not_found",
			employeeId: 999,
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				notFoundErr := errors.WrapErrNotFound("employee", "id", 999)
				mockRepo.On("GetEmployeeInboundOrdersReportByEmployeeId", 999).Return(models.EmployeeInboundOrdersReport{}, notFoundErr)
			},
			validateResult: func(t *testing.T, result models.EmployeeInboundOrdersReport, err error) {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "not found")
				assert.Contains(t, err.Error(), "employee")
				assert.Equal(t, models.EmployeeInboundOrdersReport{}, result)
			},
		},
		{
			name:       "get_report_by_employee_id_repository_error",
			employeeId: 1,
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				mockRepo.On("GetEmployeeInboundOrdersReportByEmployeeId", 1).Return(models.EmployeeInboundOrdersReport{}, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, result models.EmployeeInboundOrdersReport, err error) {
				require.Error(t, err)
				assert.Equal(t, errors.ErrGeneral, err)
				assert.Equal(t, models.EmployeeInboundOrdersReport{}, result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockInboundOrderRepo := new(inbound_order.MockInboundOrderRepository)
			mockEmployeeRepo := new(mocks.MockEmployeeRepository)
			service := NewService(mockInboundOrderRepo, mockEmployeeRepo)
			tt.setupInboundOrderMock(mockInboundOrderRepo)

			// Act
			result, err := service.GetEmployeeInboundOrdersReportByEmployeeId(tt.employeeId)

			// Assert
			tt.validateResult(t, result, err)
			mockInboundOrderRepo.AssertExpectations(t)
			mockEmployeeRepo.AssertExpectations(t)
		})
	}
}

func TestService_GetEmployeeInboundOrdersReportAll(t *testing.T) {
	tests := []struct {
		name                  string
		setupInboundOrderMock func(*inbound_order.MockInboundOrderRepository)
		validateResult        func(*testing.T, []models.EmployeeInboundOrdersReport, error)
	}{
		{
			name: "get_all_reports_success",
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				expectedReports := []models.EmployeeInboundOrdersReport{
					{
						ID:                 1,
						CardNumberID:       "EMP001",
						FirstName:          "John",
						LastName:           "Doe",
						WarehouseID:        1,
						InboundOrdersCount: 5,
					},
					{
						ID:                 2,
						CardNumberID:       "EMP002",
						FirstName:          "Jane",
						LastName:           "Smith",
						WarehouseID:        2,
						InboundOrdersCount: 3,
					},
				}
				mockRepo.On("GetEmployeeInboundOrdersReportAll").Return(expectedReports, nil)
			},
			validateResult: func(t *testing.T, result []models.EmployeeInboundOrdersReport, err error) {
				require.NoError(t, err)
				require.Len(t, result, 2)
				assert.Equal(t, 1, result[0].ID)
				assert.Equal(t, "EMP001", result[0].CardNumberID)
				assert.Equal(t, "John", result[0].FirstName)
				assert.Equal(t, "Doe", result[0].LastName)
				assert.Equal(t, 1, result[0].WarehouseID)
				assert.Equal(t, 5, result[0].InboundOrdersCount)
				assert.Equal(t, 2, result[1].ID)
				assert.Equal(t, "EMP002", result[1].CardNumberID)
				assert.Equal(t, "Jane", result[1].FirstName)
				assert.Equal(t, "Smith", result[1].LastName)
				assert.Equal(t, 2, result[1].WarehouseID)
				assert.Equal(t, 3, result[1].InboundOrdersCount)
			},
		},
		{
			name: "get_all_reports_empty_result",
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				mockRepo.On("GetEmployeeInboundOrdersReportAll").Return([]models.EmployeeInboundOrdersReport{}, nil)
			},
			validateResult: func(t *testing.T, result []models.EmployeeInboundOrdersReport, err error) {
				require.NoError(t, err)
				assert.Empty(t, result)
			},
		},
		{
			name: "get_all_reports_nil_result",
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				mockRepo.On("GetEmployeeInboundOrdersReportAll").Return(nil, nil)
			},
			validateResult: func(t *testing.T, result []models.EmployeeInboundOrdersReport, err error) {
				require.NoError(t, err)
				assert.Nil(t, result)
			},
		},
		{
			name: "get_all_reports_repository_error",
			setupInboundOrderMock: func(mockRepo *inbound_order.MockInboundOrderRepository) {
				mockRepo.On("GetEmployeeInboundOrdersReportAll").Return(nil, errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, result []models.EmployeeInboundOrdersReport, err error) {
				require.Error(t, err)
				assert.Equal(t, errors.ErrGeneral, err)
				assert.Nil(t, result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockInboundOrderRepo := new(inbound_order.MockInboundOrderRepository)
			mockEmployeeRepo := new(mocks.MockEmployeeRepository)
			service := NewService(mockInboundOrderRepo, mockEmployeeRepo)
			tt.setupInboundOrderMock(mockInboundOrderRepo)

			// Act
			result, err := service.GetEmployeeInboundOrdersReportAll()

			// Assert
			tt.validateResult(t, result, err)
			mockInboundOrderRepo.AssertExpectations(t)
			mockEmployeeRepo.AssertExpectations(t)
		})
	}
}
