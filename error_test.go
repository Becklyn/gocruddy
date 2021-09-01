package gocruddy_test

import (
	"errors"
	gocruddy "github.com/Becklyn/go-cruddy"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCRUDError(t *testing.T) {
	assert.NotNil(t, gocruddy.NewCRUDError(fiber.StatusOK, errors.New("test")))
}

func TestCRUDError_Error(t *testing.T) {
	err := errors.New("test")
	crudErr := gocruddy.NewCRUDError(fiber.StatusOK, err)

	assert.Equal(t, err.Error(), crudErr.Error())
}

func TestCRUDError_Unwrap(t *testing.T) {
	err := errors.New("test")
	crudErr := gocruddy.NewCRUDError(fiber.StatusOK, err)

	assert.Equal(t, err, crudErr.Unwrap())
}
