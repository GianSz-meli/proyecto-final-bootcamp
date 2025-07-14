package locality

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
)

type SqlLocalityRepository struct {
	db *sql.DB
}

func NewSqlLocalityRepository(db *sql.DB) *SqlLocalityRepository {
	return &SqlLocalityRepository{
		db: db,
	}
}

func (r *SqlLocalityRepository) GetById(id int) (*models.Locality, error) {
	query := `
		SELECT 
			l.id, l.locality_name, 
			p.id, p.province_name, 
			c.id, c.country_name
		FROM localities l
		LEFT JOIN provinces p ON l.province_id = p.id
		LEFT JOIN countries c ON p.country_id = c.id
		WHERE l.id = ?
	`

	var (
		locality    models.Locality
		province    models.Province
		country     models.Country
	)

	err := r.db.QueryRow(query, id).Scan(
		&locality.Id,
		&locality.LocalityName,
		&province.Id,
		&province.ProvinceName,
		&country.Id,
		&country.CountryName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	province.Country = &country
	locality.Province = &province

	return &locality, nil
}
