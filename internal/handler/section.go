package handler

import (
	service "ProyectoFinal/internal/service/section"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    sections,
		})
	}
}

func (h *SectionDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id parameter",
			})
			return
		}

		section, err := h.sv.GetByID(idNum)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    section,
		})
	}
}

func (h *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var section models.Section
		if err := json.NewDecoder(r.Body).Decode(&section); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		createdSection, err := h.sv.Create(section)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "section created successfully",
			"data":    createdSection,
		})
	}
}

func (h *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id parameter",
			})
			return
		}

		var section models.Section
		if err := json.NewDecoder(r.Body).Decode(&section); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		updatedSection, err := h.sv.Update(idNum, section)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "section updated successfully",
			"data":    updatedSection,
		})
	}
}

func (h *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id parameter",
			})
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
