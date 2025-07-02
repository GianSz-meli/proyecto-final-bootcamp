package handler

import (
	"net/http"
	"strconv"

	employeeService "ProyectoFinal/internal/service/employee"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

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

		employeeDocs := make([]models.EmployeeDoc, len(employees))
		for i, emp := range employees {
			employeeDocs[i] = emp.ModelToDoc()
		}

		body := map[string][]models.EmployeeDoc{
			"data": employeeDocs,
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
			pkgErrors.HandleError(w, err)
			return
		}

		body := map[string]models.EmployeeDoc{
			"data": employee.ModelToDoc(),
		}

		response.JSON(w, http.StatusOK, body)
	}
}

func (h *EmployeeHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.EmployeeRequest

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		if err := validateEmployee.Struct(reqBody); err != nil {
			newError := pkgErrors.WrapErrUnprocessableEntity(err)
			pkgErrors.HandleError(w, newError)
			return
		}

		employee := models.Employee{
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

		body := map[string]models.EmployeeDoc{
			"data": createdEmployee.ModelToDoc(),
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
			pkgErrors.HandleError(w, err)
			return
		}

		var reqBody models.EmployeeUpdateRequest
		if err := request.JSON(r, &reqBody); err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		updatedEmployee := currentEmployee
		if reqBody.CardNumberID != nil {
			updatedEmployee.CardNumberID = *reqBody.CardNumberID
		}
		if reqBody.FirstName != nil {
			updatedEmployee.FirstName = *reqBody.FirstName
		}
		if reqBody.LastName != nil {
			updatedEmployee.LastName = *reqBody.LastName
		}
		if reqBody.WarehouseID != nil {
			updatedEmployee.WarehouseID = *reqBody.WarehouseID
		}

		finalEmployee, err := h.service.Update(id, updatedEmployee)
		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		body := map[string]models.EmployeeDoc{
			"data": finalEmployee.ModelToDoc(),
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
			pkgErrors.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
