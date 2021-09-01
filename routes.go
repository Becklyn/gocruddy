package gocruddy

import (
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CrudConfig configuration of a crud api for an entity
type CrudConfig interface {
	UseContainer(c Container)
	GetEntityEntry() interface{}
	GetBasePath() string
	CreateGetFilter(ctx *fiber.Ctx) DatabaseFilter
	CreateUpdateFilter(ctx *fiber.Ctx) DatabaseFilter
	CreateDeleteFilter(ctx *fiber.Ctx) DatabaseFilter
	SerializeList(entries []interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (serialized fiber.Map, err error)
	MapPostEntry(entry interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (mapped Entry, err error)
	MapPutEntry(entry interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (mapped Entry, err error)
}

// DatabaseFilter a function that filters a database query
type DatabaseFilter func(db *gorm.DB) *gorm.DB

// Entry a single instance of an entity
type Entry interface {
	GetID() uint
}

// RegisterCrudRoutes register all crud routes based on their configuration
func RegisterCrudRoutes(router fiber.Router, c Container) {
	for _, config := range c.GetCrudConfigs() {
		config.UseContainer(c)

		group := router.Group(config.GetBasePath())

		group.Get("/", Get(c, config))
		group.Post("/", Post(c, config))
		group.Put("/:id", Put(c, config))
		group.Delete("/:id", Delete(c, config))
	}
}
