package errors

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"regexp"
)

func HandleMysqlError(err error) error {
	var mysqlError *mysql.MySQLError
	log.Println(err)
	if errors.As(err, &mysqlError) {
		switch mysqlError.Number {
		case 1062:
			return HandleDuplicatedEntryError(err) //Conflict
		case 1452:
			return HandleViolationFkError(err)
		case 1451:
			//TODO: Cannot delete or update parent row (viola FK padre) 409 CONFLICT
			return err
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
		return WrapErrAlreadyExist(domain, property, value)
	}
	return fmt.Errorf("%w: %s", ErrAlreadyExists, "Duplicate entry")
}
func HandleViolationFkError(err error) error {
	re := regexp.MustCompile("FOREIGN KEY \\(`([^`]*)`\\)")
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		fkField := matches[1]
		return fmt.Errorf("%w: foreign key constraint failed on field '%s'", ErrBadRequest, fkField)
	}
	return fmt.Errorf("%w: foreign key constraint failed (unknown field)", ErrBadRequest)
}
