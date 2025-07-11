package di

import (
	"ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/section"
	service "ProyectoFinal/internal/service/section"
	"ProyectoFinal/pkg/models"
)

func GetSectionHandler(db map[int]models.Section) *handler.SectionDefault {
	sectionRepository := repository.NewSectionMap(db)
	sectionService := service.NewSectionDefault(sectionRepository)
	sectionHandler := handler.NewSectionDefault(sectionService)
	return sectionHandler
}
