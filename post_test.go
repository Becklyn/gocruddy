package gocruddy_test

import (
	"encoding/json"
	"errors"
	"github.com/Becklyn/gocruddy"
	"github.com/Becklyn/gocruddy/test"
	"github.com/Becklyn/gocruddy/test/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

func initPostServer() (*fiber.App, *mock.Container, *mock.Log, *mock.Database, *mock.CrudRepo, *mock.CrudConfig) {
	app := fiber.New()
	c := &mock.Container{}
	log := &mock.Log{}
	db := &mock.Database{}
	crudRepo := &mock.CrudRepo{}
	config := &mock.CrudConfig{}

	c.On("GetLogger").Return(log).Once()
	c.On("GetDatabase").Return(db).Once()
	c.On("GetCrudRepo").Return(crudRepo).Once()
	app.Post("/", gocruddy.Post(c, config))

	return app, c, log, db, crudRepo, config
}

func assertAllPostExpectations(
	t *testing.T,
	c *mock.Container,
	log *mock.Log,
	db *mock.Database,
	crudRepo *mock.CrudRepo,
	config *mock.CrudConfig,
) {
	c.AssertExpectations(t)
	log.AssertExpectations(t)
	db.AssertExpectations(t)
	crudRepo.AssertExpectations(t)
	config.AssertExpectations(t)
}

func TestPost_TransactionError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPostServer()

	db.On("UseTransaction", mock.Anything).Return(errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("POST", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllPostExpectations(t, c, log, db, crudRepo, config)
}

func TestPost_MapPostEntryError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPostServer()

	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	config.On("MapPostEntry", entry, mock.Anything, mock.Anything).Return(&mock.Entity{}, errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("POST", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllPostExpectations(t, c, log, db, crudRepo, config)
}

func TestPost_MapPostEntryCRUDError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPostServer()

	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	config.On("MapPostEntry", entry, mock.Anything, mock.Anything).Return(&mock.Entity{}, gocruddy.NewError(fiber.StatusBadRequest, errors.New("test"))).Once()
	log.On("ErrWarn", mock.Anything).Once()

	req := httptest.NewRequest("POST", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusBadRequest)

	assertAllPostExpectations(t, c, log, db, crudRepo, config)
}

func TestPost_InsertError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPostServer()

	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	config.On("MapPostEntry", entry, mock.Anything, mock.Anything).Return(&mock.Entity{}, nil).Once()
	crudRepo.On("Insert", mock.Anything, entry).Return(errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("POST", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllPostExpectations(t, c, log, db, crudRepo, config)
}

func TestPost_Success(t *testing.T) {
	app, c, log, db, crudRepo, config := initPostServer()

	entry := &mock.Entity{
		Model: gorm.Model{
			ID: 123,
		},
	}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	config.On("MapPostEntry", entry, mock.Anything, mock.Anything).Return(entry, nil).Once()
	crudRepo.On("Insert", mock.Anything, entry).Return(nil).Once()

	req := httptest.NewRequest("POST", "http://example.com/", nil)

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var data struct {
		ID uint `json:"id"`
	}

	assert.Nil(t, json.NewDecoder(resp.Body).Decode(&data))
	assert.Equal(t, uint(123), data.ID)

	assertAllPostExpectations(t, c, log, db, crudRepo, config)
}
