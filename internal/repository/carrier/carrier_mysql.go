package carrier

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
)

type carrierMySql struct {
	db *sql.DB
}

func NewSqlCarrierRepository(newDB *sql.DB) CarrierRepository {
	return &carrierMySql{
		db: newDB,
	}
}

func (r *carrierMySql) Create(carrier *models.Carrier) (*models.Carrier, error) {
	result, execErr := r.db.Exec(CREATE_CARRIER, carrier.Cid, carrier.CompanyName, carrier.Address, carrier.Telephone, carrier.LocalityId)
	if execErr != nil {
		return nil, execErr
	}

	id, idErr := result.LastInsertId()
	if idErr != nil {
		return nil, idErr
	}

	carrier.Id = int(id)
	return carrier, nil
}
