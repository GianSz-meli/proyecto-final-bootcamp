package handler

import (
	service "ProyectoFinal/internal/service/section"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"net/http"

	"ProyectoFinal/internal/handler/utils"

	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
)

var sectionValidator = validator.New()

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
			sectionsDoc = append(sectionsDoc, section.ToSectionDoc())
		}

		body := models.SuccessResponse{
			Data: sectionsDoc,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *SectionDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idNum, err := utils.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		section, err := h.sv.GetByID(idNum)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{
			Data: section.ToSectionDoc(),
		})
	}
}

func (h *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var section models.Section
		if err := json.NewDecoder(r.Body).Decode(&section); err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		if err := sectionValidator.Struct(section); err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		createdSection, err := h.sv.Create(section)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusCreated, models.SuccessResponse{
			Data: createdSection.ToSectionDoc(),
		})
	}
}

func (h *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idNum, err := utils.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		var section models.Section
		if err := json.NewDecoder(r.Body).Decode(&section); err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		if err := sectionValidator.Struct(section); err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		updatedSection, err := h.sv.Update(idNum, section)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{
			Data: updatedSection.ToSectionDoc(),
		})
	}
}

func (h *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idNum, err := utils.GetParamInt(r, "id")
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
