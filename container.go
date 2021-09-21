package gocruddy

import (
	"github.com/ao-concepts/logging"
	"github.com/ao-concepts/storage"
	"gorm.io/gorm"
)

// Container service container interface
type Container interface {
	GetCrudConfigs() []CrudConfig
	GetLogger() logging.Logger
	GetDatabase() Database
	GetCrudRepo() CrudRepository
}

// Database interface
type Database interface {
	UseTransaction(fn storage.HandlerFunc) error
	Gorm() *gorm.DB
	Begin() (tx *storage.Transaction, err error)
}

type Repo interface {
	Insert(tx *storage.Transaction, entry interface{}) (err error)
	Update(tx *storage.Transaction, entry interface{}) (err error)
	Delete(tx *storage.Transaction, entry interface{}) (err error)
	Remove(tx *storage.Transaction, entry interface{}) (err error)
}

type CrudRepository interface {
	Repo
	GetAllEntries(tx *storage.Transaction, filter DatabaseFilter, t interface{}) (entries []interface{}, err error)
	GetByID(tx *storage.Transaction, id uint, filter DatabaseFilter, t interface{}) (entry interface{}, err error)
}

// CrudContainer is a basic service container that can be used with gocruddy
type CrudContainer struct {
	crudConfigs []CrudConfig
}

// UseCrudConfig register a crud configuration
func (c *CrudContainer) UseCrudConfig(crud CrudConfig) {
	c.crudConfigs = append(c.crudConfigs, crud)
}

// GetCrudConfigs return all registered crud configurations
func (c *CrudContainer) GetCrudConfigs() []CrudConfig {
	return c.crudConfigs
}
