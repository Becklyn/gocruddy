package mock

import (
	"github.com/Becklyn/gocruddy"
	"github.com/ao-concepts/storage"
)

type CrudRepo struct {
	Repo
}

func (r *CrudRepo) GetAllEntries(tx *storage.Transaction, filter gocruddy.DatabaseFilter, t interface{}) (entries []interface{}, err error) {
	args := r.Called(tx, filter, t)
	return args.Get(0).([]interface{}), args.Error(1)
}

func (r *CrudRepo) GetByID(tx *storage.Transaction, id uint, filter gocruddy.DatabaseFilter, t interface{}) (entry interface{}, err error) {
	args := r.Called(tx, id, filter, t)
	return args.Get(0).(interface{}), args.Error(1)
}
