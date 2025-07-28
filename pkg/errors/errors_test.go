package errors

import (
	"errors"
	"testing"

	"github.com/palantir/stacktrace"
	"github.com/stretchr/testify/assert"
)

func TestAppError_Error(t *testing.T) {
	// Test through the Compile function which properly initializes AppError
	testErr := errors.New("Test error message")
	_, appErr := Compile(COMMON, testErr, "en", false)

	result := appErr.Error()
	assert.Equal(t, "Test error message", result)
}

func TestCompile_WithValidError(t *testing.T) {
	// Create a test error with stacktrace
	testErr := stacktrace.NewError("test error")

	statusCode, appErr := Compile(COMMON, testErr, "en", true)

	assert.NotZero(t, statusCode)
	assert.NotZero(t, appErr.Code)
	assert.NotEmpty(t, appErr.Message)
	assert.NotNil(t, appErr.DebugError)
}

func TestCompile_WithRegularError(t *testing.T) {
	testErr := errors.New("regular error")

	statusCode, appErr := Compile(COMMON, testErr, "en", false)

	assert.Equal(t, 500, statusCode) // Should default to internal server error
	assert.NotZero(t, appErr.Code)
	assert.NotEmpty(t, appErr.Message)
	assert.Nil(t, appErr.DebugError) // Debug mode is false
}

func TestCompile_WithDebugMode(t *testing.T) {
	testErr := errors.New("debug error")

	statusCode, appErr := Compile(COMMON, testErr, "en", true)

	assert.NotZero(t, statusCode)
	assert.NotNil(t, appErr.DebugError)
	assert.Contains(t, *appErr.DebugError, "debug error")
}

func TestCompile_WithoutDebugMode(t *testing.T) {
	testErr := errors.New("hidden error")

	statusCode, appErr := Compile(COMMON, testErr, "en", false)

	assert.NotZero(t, statusCode)
	assert.Nil(t, appErr.DebugError)
}

func TestServiceType_Constants(t *testing.T) {
	assert.Equal(t, ServiceType(1), COMMON)
	assert.Equal(t, ServiceType(2), HTTP)
	assert.Equal(t, ServiceType(3), SQL)
	assert.Equal(t, ServiceType(4), INTERNAL)
	assert.Equal(t, ServiceType(5), CACHE)
}

func TestValidationError_Struct(t *testing.T) {
	validationErr := ValidationError{
		Field:   "email",
		Message: "Invalid email format",
	}

	assert.Equal(t, "email", validationErr.Field)
	assert.Equal(t, "Invalid email format", validationErr.Message)
}
