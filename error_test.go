package gocruddy_test

import (
	"errors"
	"github.com/Becklyn/gocruddy"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCRUDError(t *testing.T) {
	assert.NotNil(t, gocruddy.NewError(fiber.StatusOK, errors.New("test")))
}

func TestCRUDError_Error(t *testing.T) {
	err := errors.New("test")
	crudErr := gocruddy.NewError(fiber.StatusOK, err)

	assert.Equal(t, err.Error(), crudErr.Error())
}

func TestCRUDError_Unwrap(t *testing.T) {
	err := errors.New("test")
	crudErr := gocruddy.NewError(fiber.StatusOK, err)

	assert.Equal(t, err, crudErr.Unwrap())
}
