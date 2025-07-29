package di

import (
	"ProyectoFinal/internal/handler/locality"
	repository "ProyectoFinal/internal/repository/locality"
	service "ProyectoFinal/internal/service/locality"
	"database/sql"
)

// GetLocalityHandler initializes and returns a LocalityHandler with the provided database connection.
func GetLocalityHandler(db *sql.DB) *locality.LocalityHandler {
	repo := repository.NewLocalityMysqlRepository(db)
	localityServce := service.NewLocalityService(repo)
	localityHandler := locality.NewLocalityHandler(localityServce)
	return localityHandler
}
