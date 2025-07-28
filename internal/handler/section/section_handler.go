package section

import (
	service "ProyectoFinal/internal/service/section"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"net/http"

	utilsHandler "ProyectoFinal/internal/handler/utils"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

// NewSectionDefault creates a new instance of SectionDefault with the provided service
func NewSectionDefault(sv service.SectionService) *SectionDefault {
	return &SectionDefault{sv: sv}
}

// SectionDefault handles HTTP requests for section operations
type SectionDefault struct {
	sv service.SectionService
}

// GetAll handles GET requests to retrieve all sections
// Returns a JSON response with all sections or an error response
func (h *SectionDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sections, err := h.sv.GetAll()
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		sectionsDoc := make([]models.SectionDoc, 0, len(sections))
		for _, section := range sections {
			sectionsDoc = append(sectionsDoc, section.ModelToDoc())
		}

		body := models.SuccessResponse{
			Data: sectionsDoc,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

// GetById handles GET requests to retrieve a section by its ID
// Extracts the ID from URL parameters and returns the section or an error response
func (h *SectionDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idNum, err := utilsHandler.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		section, err := h.sv.GetById(idNum)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{
			Data: section.ModelToDoc(),
		})
	}
}

// Create handles POST requests to create a new section
// Validates the request body and creates the section, returning the created section or an error response
func (h *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.CreateSectionRequest

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := errors.WrapErrBadRequest(err)
			errors.HandleError(w, newErr)
			return
		}

		if err := utilsHandler.ValidateRequestData(reqBody); err != nil {
			errors.HandleError(w, err)
			return
		}

		model := reqBody.DocToModel()

		createdSection, err := h.sv.Create(model)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusCreated, models.SuccessResponse{
			Data: createdSection.ModelToDoc(),
		})
	}
}

// Update handles PUT requests to update an existing section
// Extracts the ID from URL parameters, validates the request body, and updates the section
// Returns the updated section or an error response
func (h *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idNum, err := utilsHandler.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		var reqBody models.UpdateSectionRequest

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := errors.WrapErrBadRequest(err)
			errors.HandleError(w, newErr)
			return
		}

		if err := utilsHandler.ValidateRequestData(reqBody); err != nil {
			errors.HandleError(w, err)
			return
		}

		updatedSection, err := h.sv.UpdateWithValidation(idNum, reqBody)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{
			Data: updatedSection.ModelToDoc(),
		})
	}
}

// Delete handles DELETE requests to remove a section by its ID
// Extracts the ID from URL parameters and deletes the section, returning a 204 No Content response
func (h *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idNum, err := utilsHandler.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		err = h.sv.Delete(idNum)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
