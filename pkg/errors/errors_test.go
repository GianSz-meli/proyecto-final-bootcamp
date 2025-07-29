package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
)

// Test constructors for ApiError types
func TestNewErrInternalServer(t *testing.T) {
	err := NewErrInternalServer()
	require.Equal(t, http.StatusInternalServerError, err.StatusCode)
}

func TestNewErrNotFound(t *testing.T) {
	err := NewErrNotFound()
	require.Equal(t, http.StatusNotFound, err.StatusCode)
}

func TestNewErrConflict(t *testing.T) {
	err := NewErrConflict()
	require.Equal(t, http.StatusConflict, err.StatusCode)
}

func TestNewErrBadRequest(t *testing.T) {
	err := NewErrBadRequest()
	require.Equal(t, http.StatusBadRequest, err.StatusCode)
}

func TestNewErrUnprocessableEntity(t *testing.T) {
	err := NewErrUnprocessableEntity()
	require.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
}

// Test getMappedError function
func TestGetMappedError(t *testing.T) {
	tests := []struct {
		name                string
		inputError          error
		expectedMappedError *ApiError
	}{
		{
			name:                "ErrGeneral maps to internal server error",
			inputError:          ErrGeneral,
			expectedMappedError: &ApiError{StatusCode: http.StatusInternalServerError},
		},
		{
			name:                "ErrNotFound maps to not found error",
			inputError:          ErrNotFound,
			expectedMappedError: &ApiError{StatusCode: http.StatusNotFound},
		},
		{
			name:                "ErrConflict maps to conflict error",
			inputError:          ErrConflict,
			expectedMappedError: &ApiError{StatusCode: http.StatusConflict},
		},
		{
			name:                "ErrBadRequest maps to bad request error",
			inputError:          ErrBadRequest,
			expectedMappedError: &ApiError{StatusCode: http.StatusBadRequest},
		},
		{
			name:                "ErrUnprocessableEntity maps to unprocessable entity error",
			inputError:          ErrUnprocessableEntity,
			expectedMappedError: &ApiError{StatusCode: http.StatusUnprocessableEntity},
		},
		{
			name:                "wrapped error maps correctly",
			inputError:          fmt.Errorf("wrapped: %w", ErrNotFound),
			expectedMappedError: &ApiError{StatusCode: http.StatusNotFound},
		},
		{
			name:                "unknown error returns nil",
			inputError:          errors.New("unknown error"),
			expectedMappedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getMappedError(test.inputError)
			if test.expectedMappedError == nil {
				require.Nil(t, result)
			} else {
				require.NotNil(t, result)
				require.Equal(t, test.expectedMappedError.StatusCode, result.StatusCode)
			}
		})
	}
}

// Test HandleError function
func TestHandleError(t *testing.T) {
	tests := []struct {
		name                 string
		inputError           error
		expectedStatusCode   int
		expectedBodyContains string
	}{
		{
			name:                 "ErrGeneral returns 500",
			inputError:           ErrGeneral,
			expectedStatusCode:   http.StatusInternalServerError,
			expectedBodyContains: "internal server error",
		},
		{
			name:                 "ErrNotFound returns 404",
			inputError:           ErrNotFound,
			expectedStatusCode:   http.StatusNotFound,
			expectedBodyContains: "not found",
		},
		{
			name:                 "ErrConflict returns 409",
			inputError:           ErrConflict,
			expectedStatusCode:   http.StatusConflict,
			expectedBodyContains: "conflict",
		},
		{
			name:                 "ErrBadRequest returns 400",
			inputError:           ErrBadRequest,
			expectedStatusCode:   http.StatusBadRequest,
			expectedBodyContains: "bad request",
		},
		{
			name:                 "ErrUnprocessableEntity returns 422",
			inputError:           ErrUnprocessableEntity,
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedBodyContains: "unprocessable entity",
		},
		{
			name:                 "wrapped error returns correct status",
			inputError:           fmt.Errorf("wrapped: %w", ErrNotFound),
			expectedStatusCode:   http.StatusNotFound,
			expectedBodyContains: "wrapped: not found",
		},
		{
			name:                 "unknown error returns 500",
			inputError:           errors.New("unknown error"),
			expectedStatusCode:   http.StatusInternalServerError,
			expectedBodyContains: "Internal Server Error",
		},
		{
			name:                 "MySQL duplicate entry error returns 409",
			inputError:           &mysql.MySQLError{Number: 1062, Message: "Duplicate entry '123' for key 'users.email'"},
			expectedStatusCode:   http.StatusConflict,
			expectedBodyContains: "users with email 123 already exists",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			HandleError(w, test.inputError)

			require.Equal(t, test.expectedStatusCode, w.Code)
			require.Contains(t, strings.ToLower(w.Body.String()), strings.ToLower(test.expectedBodyContains))
		})
	}
}

// Test wrapper functions
func TestWrapErrConflict(t *testing.T) {
	err := WrapErrConflict("user", "email", "test@example.com")
	require.Contains(t, err.Error(), "conflict")
	require.Contains(t, err.Error(), "user with email test@example.com already exists")
	require.True(t, errors.Is(err, ErrConflict))
}

func TestWrapErrBadRequest(t *testing.T) {
	originalErr := errors.New("invalid input")
	err := WrapErrBadRequest(originalErr)
	require.Contains(t, err.Error(), "bad request")
	require.Contains(t, err.Error(), "invalid input")
	require.True(t, errors.Is(err, ErrBadRequest))
}

func TestWrapErrUnprocessableEntity(t *testing.T) {
	originalErr := errors.New("validation failed")
	err := WrapErrUnprocessableEntity(originalErr)
	require.Contains(t, err.Error(), "unprocessable entity")
	require.Contains(t, err.Error(), "validation failed")
	require.True(t, errors.Is(err, ErrUnprocessableEntity))
}

func TestWrapErrNotFound(t *testing.T) {
	err := WrapErrNotFound("user", "id", 123)
	require.Contains(t, err.Error(), "not found")
	require.Contains(t, err.Error(), "user with id 123 not found")
	require.True(t, errors.Is(err, ErrNotFound))
}
