package errors

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"regexp"
)

// HandleMysqlError inspects a MySQL error and returns a wrapped application-specific error
// based on the error number. If the error does not match any known MySQL error code, it returns the original error.
func HandleMysqlError(err error) error {
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) {
		switch mysqlError.Number {
		case 1062:
			return HandleDuplicatedEntryError(err)
		case 1452:
			return HandleViolationFkError(err)
		case 1048:
			return HandleColumnRequired(err)
		case 1451:
			return HandleParentRowError(err)
		}
	}
	return err
}

// HandleDuplicatedEntryError parses a MySQL duplicate entry error and wraps it in a conflict error
// with detailed information about the conflicting value, domain, and property.
func HandleDuplicatedEntryError(err error) error {
	re := regexp.MustCompile(`Duplicate entry '([^']*)' for key '([^\.]+)\.([^\']*)'`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) == 4 {
		value := matches[1]    // "66557"
		domain := matches[2]   // Seller, warehouse, etc
		property := matches[3] // "cid"
		if property == "PRIMARY" {
			property = "id"
		}
		return WrapErrConflict(domain, property, value)
	}
	return fmt.Errorf("%w: %s", ErrConflict, "Duplicate entry")
}

// HandleViolationFkError parses a MySQL foreign key violation error and wraps it
// with information about the violated foreign key field, indicating that the reference is invalid.
func HandleViolationFkError(err error) error {
	re := regexp.MustCompile("FOREIGN KEY \\(`([^`]*)`\\)")
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		fkField := matches[1]
		return fmt.Errorf("%w: invalid value: %s refers to a non-existent or deleted record", ErrConflict, fkField)
	}
	return fmt.Errorf("%w: invalid reference: one of the linked objects was not found", ErrConflict)
}

// HandleColumnRequired parses a MySQL error about required (non-nullable) columns
// and wraps it as a bad request error indicating which column is missing or null.
func HandleColumnRequired(err error) error {
	re := regexp.MustCompile(`Column '([^']*)'`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		return fmt.Errorf("%w: %s cannot be null", ErrBadRequest, matches[1])
	}
	return fmt.Errorf("%w: a required field is missing or null", ErrBadRequest)
}

// HandleParentRowError parses a MySQL error about parent row deletion/update constraint violations
// and wraps it as a conflict error with detailed information about the involved tables and columns.
func HandleParentRowError(err error) error {
	re := regexp.MustCompile(
		"fails \\(`[^`]+`\\.`([^`]*)`, CONSTRAINT `([^`]*)` FOREIGN KEY \\(`([^`]*)`\\) REFERENCES `([^`]*)` \\(`([^`]*)`\\)",
	)
	matches := re.FindStringSubmatch(err.Error())

	if len(matches) == 6 {
		fkColumn := matches[5]
		referencedTable := matches[4]
		affectedTable := matches[1]
		return fmt.Errorf("%w: cannot delete or update record: %s is referenced by existing %s records in %s", ErrConflict, referencedTable, fkColumn, affectedTable)
	}
	return fmt.Errorf("%w: cannot delete or update record: it is referenced by other records", ErrConflict)
}
