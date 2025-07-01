package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	employeeService "ProyectoFinal/internal/service/employee"
	pkgErrors "ProyectoFinal/pkg/errors"
	employeemodel "ProyectoFinal/pkg/models/employee"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type EmployeeHandler struct {
	service employeeService.Service
}

var validateEmployee = validator.New()

func NewEmployeeHandler(service employeeService.Service) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
	}
}

func (h *EmployeeHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := h.service.GetAll()

		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		body := map[string][]employeemodel.Employee{
			"data": employees,
		}

		response.JSON(w, http.StatusOK, body)
	}
}

func (h *EmployeeHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		employee, err := h.service.GetById(id)
		if err != nil {
			if errors.Is(err, pkgErrors.ErrNotFound) {
				newErr := fmt.Errorf("%w : employee with id %d not found", pkgErrors.ErrNotFound, id)
				pkgErrors.HandleError(w, newErr)
			} else {
				pkgErrors.HandleError(w, err)
			}
			return
		}

		body := map[string]employeemodel.Employee{
			"data": employee,
		}

		response.JSON(w, http.StatusOK, body)
	}
}

func (h *EmployeeHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody employeemodel.EmployeeRequest

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		if err := validateEmployee.Struct(reqBody); err != nil {
			newError := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newError)
			return
		}

		employee := employeemodel.Employee{
			CardNumberID: reqBody.CardNumberID,
			FirstName:    reqBody.FirstName,
			LastName:     reqBody.LastName,
			WarehouseID:  reqBody.WarehouseID,
		}

		createdEmployee, err := h.service.Create(employee)

		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		body := map[string]employeemodel.Employee{
			"data": createdEmployee,
		}
		response.JSON(w, http.StatusCreated, body)
	}
}

func (h *EmployeeHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		currentEmployee, err := h.service.GetById(id)
		if err != nil {
			if errors.Is(err, pkgErrors.ErrNotFound) {
				newErr := fmt.Errorf("%w : employee with id %d not found", pkgErrors.ErrNotFound, id)
				pkgErrors.HandleError(w, newErr)
			} else {
				pkgErrors.HandleError(w, err)
			}
			return
		}

		var reqBody employeemodel.EmployeeUpdateRequest
		if err := request.JSON(r, &reqBody); err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		updatedEmployee := currentEmployee
		if reqBody.CardNumberID != "" {
			updatedEmployee.CardNumberID = reqBody.CardNumberID
		}
		if reqBody.FirstName != "" {
			updatedEmployee.FirstName = reqBody.FirstName
		}
		if reqBody.LastName != "" {
			updatedEmployee.LastName = reqBody.LastName
		}
		if reqBody.WarehouseID != 0 {
			updatedEmployee.WarehouseID = reqBody.WarehouseID
		}

		finalEmployee, err := h.service.Update(id, updatedEmployee)
		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		body := map[string]employeemodel.Employee{
			"data": finalEmployee,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *EmployeeHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		err = h.service.Delete(id)
		if err != nil {
			if errors.Is(err, pkgErrors.ErrNotFound) {
				pkgErrors.HandleError(w, err)
			} else {
				pkgErrors.HandleError(w, err)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
