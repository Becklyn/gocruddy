package gocruddy

import (
	"github.com/ao-concepts/storage"
	"reflect"
)

// Repository crud data repository
type Repository struct {
	storage.Repository
}

// GetAllEntries returns all entities matching a filter
func (r *Repository) GetAllEntries(tx *storage.Transaction, filter DatabaseFilter, t interface{}) (entries []interface{}, err error) {
	slice := reflect.New(reflect.SliceOf(reflect.TypeOf(t))).Interface()

	if err := filter(tx.Gorm().Model(t)).Find(slice).Error; err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	sliceElement := sliceValue.Elem()

	for i := 0; i < sliceElement.Len(); i++ {
		entries = append(entries, sliceElement.Index(i).Interface())
	}

	return entries, nil
}

// GetByID fetches an entity by its unique id
func (r *Repository) GetByID(tx *storage.Transaction, id uint, filter DatabaseFilter, t interface{}) (entry interface{}, err error) {
	err = filter(tx.Gorm().Model(t)).Where("id = ?", id).First(t).Error
	return t, err
}
