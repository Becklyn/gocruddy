package app

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// User entity
type User struct {
	gorm.Model
	Name string
}

// GetID return the id of the user
func (u *User) GetID() uint {
	return u.ID
}

// Serialize an user
func (u *User) Serialize() fiber.Map {
	return fiber.Map{
		"id":   u.ID,
		"name": u.Name,
	}
}
