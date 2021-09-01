package gocruddy_test

import (
	"errors"
	"github.com/Becklyn/go-cruddy"
	"github.com/Becklyn/go-cruddy/test"
	"github.com/Becklyn/go-cruddy/test/mock"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

func initPutServer() (*fiber.App, *mock.Container, *mock.Log, *mock.Database, *mock.CrudRepo, *mock.CrudConfig) {
	app := fiber.New()
	c := &mock.Container{}
	log := &mock.Log{}
	db := &mock.Database{}
	crudRepo := &mock.CrudRepo{}
	config := &mock.CrudConfig{}

	c.On("GetLogger").Return(log).Once()
	c.On("GetDatabase").Return(db).Once()
	c.On("GetCrudRepo").Return(crudRepo).Once()
	app.Put("/:id", gocruddy.Put(c, config))

	return app, c, log, db, crudRepo, config
}

func assertAllPutExpectations(
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

func TestPut_IdError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	log.On("ErrWarn", mock.Anything).Once()

	req := httptest.NewRequest("PUT", "http://example.com/asd", nil)
	test.ExpectStatus(t, app, req, fiber.StatusBadRequest)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}

func TestPut_TransactionError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	db.On("UseTransaction", mock.Anything).Return(errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("PUT", "http://example.com/123", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}

func TestPut_GetByIDError_NotFound(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateUpdateFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entry).Twice()
	crudRepo.On("GetByID", mock.Anything, uint(123), mock.AnythingOfType("gocruddy.DatabaseFilter"), entry).Return(entry, gorm.ErrRecordNotFound).Once()
	log.On("Info", mock.Anything, mock.Anything).Once()

	req := httptest.NewRequest("PUT", "http://example.com/123", nil)
	test.ExpectStatus(t, app, req, fiber.StatusNotFound)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}

func TestPut_GetByIDError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateUpdateFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	crudRepo.On("GetByID", mock.Anything, uint(123), mock.AnythingOfType("gocruddy.DatabaseFilter"), entry).Return(entry, errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("PUT", "http://example.com/123", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}

func TestPut_MapPutEntryError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateUpdateFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	crudRepo.On("GetByID", mock.Anything, uint(123), mock.AnythingOfType("gocruddy.DatabaseFilter"), entry).Return(entry, nil).Once()
	config.On("MapPutEntry", entry, mock.Anything, mock.Anything).Return(entry, errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("PUT", "http://example.com/123", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}

func TestPut_MapPutEntryCRUDError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateUpdateFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	crudRepo.On("GetByID", mock.Anything, uint(123), mock.AnythingOfType("gocruddy.DatabaseFilter"), entry).Return(entry, nil).Once()
	config.On("MapPutEntry", entry, mock.Anything, mock.Anything).Return(entry, gocruddy.NewCRUDError(fiber.StatusBadRequest, errors.New("test"))).Once()
	log.On("ErrWarn", mock.Anything).Once()

	req := httptest.NewRequest("PUT", "http://example.com/123", nil)
	test.ExpectStatus(t, app, req, fiber.StatusBadRequest)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}

func TestPut_UpdateError(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateUpdateFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	crudRepo.On("GetByID", mock.Anything, uint(123), mock.AnythingOfType("gocruddy.DatabaseFilter"), entry).Return(entry, nil).Once()
	config.On("MapPutEntry", entry, mock.Anything, mock.Anything).Return(entry, nil).Once()
	crudRepo.On("Update", mock.Anything, entry).Return(errors.New("test")).Once()
	log.On("ErrError", mock.Anything).Once()

	req := httptest.NewRequest("PUT", "http://example.com/123", nil)
	test.ExpectStatus(t, app, req, fiber.StatusInternalServerError)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}

func TestPut_Success(t *testing.T) {
	app, c, log, db, crudRepo, config := initPutServer()

	filter := func(db *gorm.DB) *gorm.DB {
		return db
	}
	entry := &mock.Entity{}

	db.On("UseTransaction", mock.Anything).Return(nil).Once()
	config.On("CreateUpdateFilter", mock.Anything).Return(filter).Once()
	config.On("GetEntityEntry").Return(entry).Once()
	crudRepo.On("GetByID", mock.Anything, uint(123), mock.AnythingOfType("gocruddy.DatabaseFilter"), entry).Return(entry, nil).Once()
	config.On("MapPutEntry", entry, mock.Anything, mock.Anything).Return(entry, nil).Once()
	crudRepo.On("Update", mock.Anything, entry).Return(nil).Once()

	req := httptest.NewRequest("PUT", "http://example.com/123", nil)
	test.ExpectStatus(t, app, req, fiber.StatusOK)

	assertAllPutExpectations(t, c, log, db, crudRepo, config)
}
