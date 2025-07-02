package di

import (
	handler "ProyectoFinal/internal/handler"
	repositorySection "ProyectoFinal/internal/repository/section"
	serviceSection "ProyectoFinal/internal/service/section"
	"ProyectoFinal/pkg/models"
)

func GetSectionHandler(db map[int]models.Section) *handler.SectionDefault {
	sectionRepository := repositorySection.NewSectionMap(db)
	sectionService := serviceSection.NewSectionDefault(sectionRepository)
	sectionHandler := handler.NewSectionDefault(sectionService)
	return sectionHandler
}
