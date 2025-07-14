package di

import (
	"ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/section"
	service "ProyectoFinal/internal/service/section"
	"database/sql"
)

func GetSectionHandler(db *sql.DB) *handler.SectionDefault {
	sectionRepository := repository.NewSectionMySQL(db)
	sectionService := service.NewSectionDefault(sectionRepository)
	sectionHandler := handler.NewSectionDefault(sectionService)
	return sectionHandler
}
