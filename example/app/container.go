package app

import (
	"github.com/Becklyn/gocruddy"
	"github.com/ao-concepts/logging"
	"github.com/ao-concepts/storage"
	"gorm.io/driver/sqlite"
)

// Container service container
type Container struct {
	config   []gocruddy.CrudConfig
	logger   logging.Logger
	db       gocruddy.Database
	crudRepo gocruddy.CrudRepository
}

// NewContainer constructor for a new service container
func NewContainer(config []gocruddy.CrudConfig) *Container {
	logger := logging.New(logging.Debug, nil)
	db, err := storage.New(sqlite.Open(":memory:"), logger)

	if err != nil {
		logger.ErrFatal(err)
	}

	// For production use, please have a look at github.com/go-gormigrate/gormigrate
	if err := db.Gorm().AutoMigrate(&User{}); err != nil {
		logger.ErrFatal(err)
	}

	c := &Container{
		config:   config,
		logger:   logger,
		db:       db,
		crudRepo: &gocruddy.Repository{},
	}

	for _, conf := range config {
		conf.UseContainer(c)
	}

	return c
}

// GetLogger returns the logging service
func (c *Container) GetLogger() logging.Logger {
	return c.logger
}

// GetDatabase returns the database service
func (c *Container) GetDatabase() gocruddy.Database {
	return c.db
}

// GetCrudConfigs returns all crud configurations
func (c *Container) GetCrudConfigs() []gocruddy.CrudConfig {
	return c.config
}

// GetCrudRepo returns the crud repository service
func (c *Container) GetCrudRepo() gocruddy.CrudRepository {
	return c.crudRepo
}
