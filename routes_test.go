package gocruddy_test

import (
	"github.com/Becklyn/go-cruddy"
	"github.com/Becklyn/go-cruddy/test/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterCrudRoutes(t *testing.T) {
	app := fiber.New()
	c := &mock.Container{}
	log := &mock.Log{}
	db := &mock.Database{}
	crudRepo := &mock.CrudRepo{}
	config := &mock.CrudConfig{}
	configurations := []gocruddy.CrudConfig{config}

	c.On("GetCrudConfigs").Return(configurations).Once()
	config.On("UseContainer", c).Once()
	config.On("GetBasePath").Return("users").Once()

	// GET
	c.On("GetLogger").Return(log).Once()
	c.On("GetDatabase").Return(db).Once()
	c.On("GetCrudRepo").Return(crudRepo).Once()

	// POST
	c.On("GetLogger").Return(log).Once()
	c.On("GetDatabase").Return(db).Once()
	c.On("GetCrudRepo").Return(crudRepo).Once()

	// PUT
	c.On("GetLogger").Return(log).Once()
	c.On("GetDatabase").Return(db).Once()
	c.On("GetCrudRepo").Return(crudRepo).Once()

	// DELETE
	c.On("GetLogger").Return(log).Once()
	c.On("GetDatabase").Return(db).Once()
	c.On("GetCrudRepo").Return(crudRepo).Once()

	assert.NotPanics(t, func() {
		gocruddy.RegisterCrudRoutes(app, c)
	})

	// check if all routes got registered properly
	for i, handler := range app.Stack() {
		// HEAD
		if i == 0 {
			assert.Len(t, handler, 1)
			continue
		}

		// GET
		if i == 1 {
			assert.Len(t, handler, 1)
			continue
		}

		// POST
		if i == 2 {
			assert.Len(t, handler, 1)
			continue
		}

		// PUT
		if i == 3 {
			assert.Len(t, handler, 1)
			continue
		}

		// DELETE
		if i == 4 {
			assert.Len(t, handler, 1)
			continue
		}

		assert.Len(t, handler, 0)
	}

	c.AssertExpectations(t)
	log.AssertExpectations(t)
	db.AssertExpectations(t)
	crudRepo.AssertExpectations(t)
	config.AssertExpectations(t)
}
