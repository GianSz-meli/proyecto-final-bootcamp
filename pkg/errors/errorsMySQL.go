package errors

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"regexp"
)

func HandleMysqlError(err error) error {
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) {
		fmt.Println(err)
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

func HandleDuplicatedEntryError(err error) error {
	re := regexp.MustCompile(`Duplicate entry '([^']*)' for key '([^\.]+)\.([^\']*)'`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) == 4 {
		value := matches[1]    // "66557"
		domain := matches[2]   // Seller, warehouse, etc
		property := matches[3] // "cid"
		return WrapErrConflict(domain, property, value)
	}
	return fmt.Errorf("%w: %s", ErrConflict, "Duplicate entry")
}

func HandleViolationFkError(err error) error {
	re := regexp.MustCompile("FOREIGN KEY \\(`([^`]*)`\\)")
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		fkField := matches[1]
		return fmt.Errorf("%w: invalid value: %s refers to a non-existent or deleted record", ErrConflict, fkField)
	}
	return fmt.Errorf("%w: invalid reference: one of the linked objects was not found", ErrConflict)
}

func HandleColumnRequired(err error) error {
	re := regexp.MustCompile(`Column '([^']*)'`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		return fmt.Errorf("%w: %s cannot be null", ErrBadRequest, matches[1])
	}
	return fmt.Errorf("%w: a required field is missing or null", ErrBadRequest)
}

func HandleParentRowError(err error) error {
	re := regexp.MustCompile("CONSTRAINT\\s+`([^`]*)`\\s+FOREIGN KEY\\s+\\(`([^`]*)`\\)\\s+REFERENCES\\s+`([^`]*)`\\s+\\(`([^`]*)`\\)")
	matches := re.FindStringSubmatch(err.Error())

	if len(matches) == 5 {
		fkColumn := matches[4]
		referencedTable := matches[3]

		return fmt.Errorf("%w: cannot delete or update record: %s is referenced by existing %s records", ErrConflict, referencedTable, fkColumn)
	}
	return fmt.Errorf("%w: cannot delete or update record: it is referenced by other records", ErrConflict)
}
