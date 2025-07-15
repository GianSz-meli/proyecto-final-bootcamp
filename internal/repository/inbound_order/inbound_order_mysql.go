package inbound_order

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &mysqlRepository{
		db: db,
	}
}

func (r *mysqlRepository) Create(inboundOrder *models.InboundOrder) error {
	result, err := r.db.Exec(QueryCreateInboundOrder,
		inboundOrder.OrderDate,
		inboundOrder.OrderNumber,
		inboundOrder.EmployeeID,
		inboundOrder.ProductBatchID,
		inboundOrder.WarehouseID)
	if err != nil {
		return pkgErrors.HandleMysqlError(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return pkgErrors.HandleMysqlError(err)
	}

	inboundOrder.ID = int(lastInsertID)
	return nil
}

func (r *mysqlRepository) ExistsByOrderNumber(orderNumber string) (bool, error) {
	var exists bool

	err := r.db.QueryRow(QueryExistsByOrderNumber, orderNumber).Scan(&exists)
	if err != nil {
		return false, pkgErrors.HandleMysqlError(err)
	}

	return exists, nil
}

func (r *mysqlRepository) GetEmployeeInboundOrdersReportByEmployeeId(employeeId int) (models.EmployeeInboundOrdersReport, error) {
	row := r.db.QueryRow(QueryGetEmployeeInboundOrdersReportByEmployeeId, employeeId)

	var report models.EmployeeInboundOrdersReport

	err := row.Scan(&report.ID, &report.CardNumberID, &report.FirstName, &report.LastName, &report.WarehouseID, &report.InboundOrdersCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.EmployeeInboundOrdersReport{}, pkgErrors.WrapErrNotFound("employee", "id", employeeId)
		}
		return models.EmployeeInboundOrdersReport{}, pkgErrors.HandleMysqlError(err)
	}

	return report, nil
}

func (r *mysqlRepository) GetEmployeeInboundOrdersReportAll() ([]models.EmployeeInboundOrdersReport, error) {
	rows, err := r.db.Query(QueryGetEmployeeInboundOrdersReportAll)
	if err != nil {
		return nil, pkgErrors.HandleMysqlError(err)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.HandleMysqlError(err)
	}

	var reports []models.EmployeeInboundOrdersReport
	for rows.Next() {
		var report models.EmployeeInboundOrdersReport

		err := rows.Scan(&report.ID, &report.CardNumberID, &report.FirstName, &report.LastName, &report.WarehouseID, &report.InboundOrdersCount)
		if err != nil {
			return nil, pkgErrors.HandleMysqlError(err)
		}

		reports = append(reports, report)
	}

	return reports, nil
}
