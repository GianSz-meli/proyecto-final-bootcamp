package employee

import (
	"net/http"

	"ProyectoFinal/internal/handler/utils"
	employeeService "ProyectoFinal/internal/service/employee"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

type EmployeeHandler struct {
	service employeeService.Service
}

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

		responseBody := models.SuccessResponse{
			Data: employeeDocs,
		}

		response.JSON(w, http.StatusOK, responseBody)
	}
}

func (h *EmployeeHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetParamInt(r, "id")
		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		employee, err := h.service.GetById(id)
		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		responseBody := models.SuccessResponse{
			Data: employee.ModelToDoc(),
		}

		response.JSON(w, http.StatusOK, responseBody)
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

		if err := utils.ValidateRequestData(reqBody); err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		employee := reqBody.DocToModel()

		createdEmployee, err := h.service.Create(employee)

		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		responseBody := models.SuccessResponse{
			Data: createdEmployee.ModelToDoc(),
		}
		response.JSON(w, http.StatusCreated, responseBody)
	}
}

func (h *EmployeeHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetParamInt(r, "id")
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

		if err := utils.ValidateRequestData(reqBody); err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		employeeUpdated, err := h.service.PatchUpdate(id, &reqBody)
		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		responseBody := models.SuccessResponse{
			Data: employeeUpdated.ModelToDoc(),
		}
		response.JSON(w, http.StatusOK, responseBody)
	}
}

func (h *EmployeeHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetParamInt(r, "id")
		if err != nil {
			pkgErrors.HandleError(w, err)
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
