package gocruddy_test

import (
	"testing"

	"github.com/Becklyn/gocruddy"
	"github.com/Becklyn/gocruddy/test/mock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRepository_GetAllEntries(t *testing.T) {
	repo := gocruddy.Repository{}
	tx, rollback := mock.Transaction()
	defer rollback()

	entries, err := repo.GetAllEntries(tx, func(gorm *gorm.DB) *gorm.DB {
		return gorm
	}, &mock.Entity{})
	assert.Nil(t, err)
	assert.Len(t, entries, 0)

	assert.Nil(t, repo.Insert(tx, &mock.Entity{
		Model: gorm.Model{
			ID: 123,
		},
	}))

	entries, err = repo.GetAllEntries(tx, func(gorm *gorm.DB) *gorm.DB {
		return gorm
	}, &mock.Entity{})
	assert.Nil(t, err)
	assert.Len(t, entries, 1)

	// filter error test
	_, err = repo.GetAllEntries(tx, func(gorm *gorm.DB) *gorm.DB {
		return gorm.Where("not-existing = ?")
	}, &mock.Entity{})
	assert.NotNil(t, err)

	// filter test
	entries, err = repo.GetAllEntries(tx, func(gorm *gorm.DB) *gorm.DB {
		return gorm.Where("id = ?", 1)
	}, &mock.Entity{})
	assert.Nil(t, err)
	assert.Len(t, entries, 0)

	entries, err = repo.GetAllEntries(tx, func(gorm *gorm.DB) *gorm.DB {
		return gorm.Where("id = ?", 123)
	}, &mock.Entity{})
	assert.Nil(t, err)
	assert.Len(t, entries, 1)
}

func TestRepository_GetByID(t *testing.T) {
	repo := gocruddy.Repository{}
	tx, rollback := mock.Transaction()
	defer rollback()

	_, err := repo.GetByID(tx, 1, func(gorm *gorm.DB) *gorm.DB {
		return gorm
	}, &mock.Entity{})
	assert.NotNil(t, err)

	assert.Nil(t, repo.Insert(tx, &mock.Entity{
		Model: gorm.Model{
			ID: 1,
		},
	}))
	assert.Nil(t, repo.Insert(tx, &mock.Entity{
		Model: gorm.Model{
			ID: 2,
		},
	}))

	entry, err := repo.GetByID(tx, 1, func(gorm *gorm.DB) *gorm.DB {
		return gorm
	}, &mock.Entity{})
	assert.Nil(t, err)
	assert.Equal(t, uint(1), entry.(*mock.Entity).ID)

	entry, err = repo.GetByID(tx, 2, func(gorm *gorm.DB) *gorm.DB {
		return gorm
	}, &mock.Entity{})
	assert.Nil(t, err)
	assert.Equal(t, uint(2), entry.(*mock.Entity).ID)

	_, err = repo.GetByID(tx, 1, func(gorm *gorm.DB) *gorm.DB {
		return gorm.Where("not-existing = ?")
	}, &mock.Entity{})
	assert.NotNil(t, err)
}
