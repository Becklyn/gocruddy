package mock

import (
	"github.com/ao-concepts/storage"
	"github.com/stretchr/testify/mock"
)

type Repo struct {
	mock.Mock
}

func (r *Repo) Insert(tx *storage.Transaction, entry interface{}) (err error) {
	return r.Called(tx, entry).Error(0)
}

func (r *Repo) Update(tx *storage.Transaction, entry interface{}) (err error) {
	return r.Called(tx, entry).Error(0)
}

func (r *Repo) Delete(tx *storage.Transaction, entry interface{}) (err error) {
	return r.Called(tx, entry).Error(0)
}

func (r *Repo) Remove(tx *storage.Transaction, entry interface{}) (err error) {
	return r.Called(tx, entry).Error(0)
}
