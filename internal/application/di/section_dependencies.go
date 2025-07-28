package di

import (
	sectionHandler "ProyectoFinal/internal/handler/section"
	repository "ProyectoFinal/internal/repository/section"
	service "ProyectoFinal/internal/service/section"
	"database/sql"
)

func GetSectionHandler(db *sql.DB) *sectionHandler.SectionDefault {
	sectionRepository := repository.NewSectionMySQL(db)
	sectionService := service.NewSectionDefault(sectionRepository)
	sectionHandler := sectionHandler.NewSectionDefault(sectionService)
	return sectionHandler
}
