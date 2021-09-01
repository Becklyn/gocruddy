package mock

import (
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
}

// GetID return the id of the entity
func (e *Entity) GetID() uint {
	return e.ID
}
