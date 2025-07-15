package di

import (
	"ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/locality"
	service "ProyectoFinal/internal/service/locality"
	"database/sql"
)

func GetLocalityHandler(db *sql.DB) *handler.LocalityHandler {
	repo := repository.NewLocalityMysqlRepository(db)
	localityServce := service.NewLocalityService(repo)
	localityHandler := handler.NewLocalityHandler(localityServce)
	return localityHandler
}
