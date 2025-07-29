package errors

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"testing"
)

// Tests for MySQL error handling
func TestHandleMysqlError(t *testing.T) {
	tests := []struct {
		name          string
		inputError    error
		expectedError string
		expectedType  error
	}{
		{
			name:          "MySQL duplicate entry error (1062)",
			inputError:    &mysql.MySQLError{Number: 1062, Message: "Duplicate entry '123' for key 'users.email'"},
			expectedError: "users with email 123 already exists",
			expectedType:  ErrConflict,
		},
		{
			name:          "MySQL foreign key violation error (1452)",
			inputError:    &mysql.MySQLError{Number: 1452, Message: "Cannot add or update a child row: a foreign key constraint fails (`db`.`table`, CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))"},
			expectedError: "invalid value: user_id refers to a non-existent or deleted record",
			expectedType:  ErrConflict,
		},
		{
			name:          "MySQL column required error (1048)",
			inputError:    &mysql.MySQLError{Number: 1048, Message: "Column 'email' cannot be null"},
			expectedError: "email cannot be null",
			expectedType:  ErrBadRequest,
		},
		{
			name:          "MySQL parent row error (1451)",
			inputError:    &mysql.MySQLError{Number: 1451, Message: "Cannot delete or update a parent row: a foreign key constraint fails (`db`.`orders`, CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))"},
			expectedError: "cannot delete or update record: users is referenced by existing id records in orders",
			expectedType:  ErrConflict,
		},
		{
			name:          "unknown MySQL error code",
			inputError:    &mysql.MySQLError{Number: 9999, Message: "Unknown error"},
			expectedError: "Unknown error",
			expectedType:  nil,
		},
		{
			name:          "non-MySQL error",
			inputError:    errors.New("generic error"),
			expectedError: "generic error",
			expectedType:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HandleMysqlError(test.inputError)
			require.Contains(t, result.Error(), test.expectedError)

			if test.expectedType != nil {
				require.True(t, errors.Is(result, test.expectedType))
			}
		})
	}
}

func TestHandleDuplicatedEntryError(t *testing.T) {
	tests := []struct {
		name          string
		inputError    error
		expectedError string
	}{
		{
			name:          "standard duplicate entry with domain and property",
			inputError:    errors.New("Duplicate entry '123' for key 'users.email'"),
			expectedError: "conflict : users with email 123 already exists",
		},
		{
			name:          "duplicate entry with PRIMARY key",
			inputError:    errors.New("Duplicate entry '456' for key 'products.PRIMARY'"),
			expectedError: "conflict : products with id 456 already exists",
		},
		{
			name:          "malformed duplicate entry error",
			inputError:    errors.New("Duplicate entry without proper format"),
			expectedError: "conflict: Duplicate entry",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HandleDuplicatedEntryError(test.inputError)
			require.Contains(t, result.Error(), test.expectedError)
			require.True(t, errors.Is(result, ErrConflict))
		})
	}
}

func TestHandleViolationFkError(t *testing.T) {
	tests := []struct {
		name          string
		inputError    error
		expectedError string
	}{
		{
			name:          "foreign key violation with field",
			inputError:    errors.New("Cannot add or update a child row: a foreign key constraint fails (`db`.`table`, CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))"),
			expectedError: "invalid value: user_id refers to a non-existent or deleted record",
		},
		{
			name:          "foreign key violation without extractable field",
			inputError:    errors.New("Foreign key violation without proper format"),
			expectedError: "invalid reference: one of the linked objects was not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HandleViolationFkError(test.inputError)
			require.Contains(t, result.Error(), test.expectedError)
			require.True(t, errors.Is(result, ErrConflict))
		})
	}
}

func TestHandleColumnRequired(t *testing.T) {
	tests := []struct {
		name          string
		inputError    error
		expectedError string
	}{
		{
			name:          "column required with field name",
			inputError:    errors.New("Column 'email' cannot be null"),
			expectedError: "email cannot be null",
		},
		{
			name:          "column required without extractable field",
			inputError:    errors.New("Column cannot be null without proper format"),
			expectedError: "a required field is missing or null",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HandleColumnRequired(test.inputError)
			require.Contains(t, result.Error(), test.expectedError)
			require.True(t, errors.Is(result, ErrBadRequest))
		})
	}
}

func TestHandleParentRowError(t *testing.T) {
	tests := []struct {
		name          string
		inputError    error
		expectedError string
	}{
		{
			name:          "parent row error with full match",
			inputError:    errors.New("Cannot delete or update a parent row: a foreign key constraint fails (`db`.`orders`, CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))"),
			expectedError: "cannot delete or update record: users is referenced by existing id records in orders",
		},
		{
			name:          "parent row error without extractable information",
			inputError:    errors.New("Cannot delete or update a parent row without proper format"),
			expectedError: "cannot delete or update record: it is referenced by other records",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HandleParentRowError(test.inputError)
			require.Contains(t, result.Error(), test.expectedError)

			require.True(t, errors.Is(result, ErrConflict))
		})
	}
}
