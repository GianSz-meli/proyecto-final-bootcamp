package handler

import (
	service "ProyectoFinal/internal/service/section"
	utilsService "ProyectoFinal/internal/service/utils"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"net/http"

	utilsHandler "ProyectoFinal/internal/handler/utils"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

func NewSectionDefault(sv service.SectionService) *SectionDefault {
	return &SectionDefault{sv: sv}
}

type SectionDefault struct {
	sv service.SectionService
}

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

		sectionToUpdate, err := h.sv.GetById(idNum)
		if err != nil {
			newError := errors.WrapErrNotFound("section", "id", idNum)
			errors.HandleError(w, newError)
			return
		}

		if updated := utilsService.UpdateFields(&sectionToUpdate, &reqBody); !updated {
			newError := fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
			errors.HandleError(w, newError)
			return
		}

		updatedSection, err := h.sv.Update(idNum, sectionToUpdate)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{
			Data: updatedSection.ModelToDoc(),
		})
	}
}

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
