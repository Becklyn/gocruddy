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

func initGetServer() (*fiber.App, *mock.Container, *mock.Log, *mock.Database, *mock.CrudRepo, *mock.CrudConfig) {
	app := fiber.New()
	c := &mock.Container{}
	log := &mock.Log{}
	db := &mock.Database{}
	crudRepo := &mock.CrudRepo{}
	config := &mock.CrudConfig{}

	c.On("GetLogger").Return(log).Once()
	c.On("GetDatabase").Return(db).Once()
	c.On("GetCrudRepo").Return(crudRepo).Once()
	app.Get("/", gocruddy.Get(c, config))

	return app, c, log, db, crudRepo, config
}

func assertAllGetExpectations(
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

func TestGet_TransactionError(t *testing.T) {
	app, c, log, db, crudRepo, config := initGetServer()

	db.On("UseTransaction", mock.Anything).Return(errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("GET", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllGetExpectations(t, c, log, db, crudRepo, config)
}

func TestGet_GetAllEntriesError(t *testing.T) {
	app, c, log, db, crudRepo, config := initGetServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entity := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateGetFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entity).Once()
	crudRepo.On("GetAllEntries", mock.Anything, mock.AnythingOfType("gocruddy.DatabaseFilter"), entity).Return([]interface{}{}, errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("GET", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllGetExpectations(t, c, log, db, crudRepo, config)
}

func TestGet_SerializeListError(t *testing.T) {
	app, c, log, db, crudRepo, config := initGetServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entity := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateGetFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entity).Once()
	crudRepo.On("GetAllEntries", mock.Anything, mock.AnythingOfType("gocruddy.DatabaseFilter"), entity).Return([]interface{}{}, nil).Once()
	config.On("SerializeList", []interface{}{}, mock.Anything).Return(fiber.Map{}, errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("GET", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllGetExpectations(t, c, log, db, crudRepo, config)
}

func TestGet_SerializeListCRUDError(t *testing.T) {
	app, c, log, db, crudRepo, config := initGetServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entity := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateGetFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entity).Once()
	crudRepo.On("GetAllEntries", mock.Anything, mock.AnythingOfType("gocruddy.DatabaseFilter"), entity).Return([]interface{}{}, nil).Once()
	config.On("SerializeList", []interface{}{}, mock.Anything).Return(fiber.Map{}, gocruddy.NewCRUDError(fiber.StatusBadRequest, errors.New("test"))).Once()
	log.On("ErrWarn", mock.Anything).Once()

	req := httptest.NewRequest("GET", "http://example.com/", nil)
	test.ExpectStatus(t, app, req, fiber.StatusBadRequest)

	assertAllGetExpectations(t, c, log, db, crudRepo, config)
}

func TestGet_Success(t *testing.T) {
	app, c, log, db, crudRepo, config := initGetServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entity := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateGetFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entity).Once()
	crudRepo.On("GetAllEntries", mock.Anything, mock.AnythingOfType("gocruddy.DatabaseFilter"), entity).Return([]interface{}{}, nil).Once()
	config.On("SerializeList", []interface{}{}, mock.Anything).Return(fiber.Map{"key": "value"}, nil).Once()

	req := httptest.NewRequest("GET", "http://example.com/", nil)

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var data struct {
		Key string `json:"key"`
	}

	assert.Nil(t, json.NewDecoder(resp.Body).Decode(&data))
	assert.Equal(t, "value", data.Key)

	assertAllGetExpectations(t, c, log, db, crudRepo, config)
}
