package app

import (
	"errors"
	"fmt"
	"github.com/Becklyn/gocruddy"
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserCrud configuration for user crud operations
type UserCrud struct {
	container *Container
}

// UseContainer register the service container
func (uc *UserCrud) UseContainer(c gocruddy.Container) {
	// You can also extend your container by additional functions. Just cast it to your container type.
	// You may want to use your own interface type to avoid circular references.
	// This example just uses the struct type directly for simplicity.
	uc.container = c.(*Container)
}

// GetEntityEntry this database entity is used for this crud
func (uc *UserCrud) GetEntityEntry() interface{} {
	// Just return an empty instance of your entity type
	return &User{}
}

// GetBasePath return the base path of your crud routes
func (uc *UserCrud) GetBasePath() string {
	// The base path for this crud will be `/user`
	return "user"
}

// CreateGetFilter restrict access for GET routes
func (uc *UserCrud) CreateGetFilter(ctx *fiber.Ctx) gocruddy.DatabaseFilter {
	// You can do anything you want here. E.g. get data from the fiber.Ctx
	// By that data you can add a where clause to your query as in: `db.Where("name = ?", ctxUser.Name)`

	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// CreateUpdateFilter restrict access for POST and PUT routes
func (uc *UserCrud) CreateUpdateFilter(ctx *fiber.Ctx) gocruddy.DatabaseFilter {
	// You can do anything you want here. E.g. get data from the fiber.Ctx
	// By that data you can add a where clause to your query as in: `db.Where("name = ?", ctxUser.Name)`

	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// CreateDeleteFilter restrict access for DELETE routes
func (uc *UserCrud) CreateDeleteFilter(ctx *fiber.Ctx) gocruddy.DatabaseFilter {
	// You can do anything you want here. E.g. get data from the fiber.Ctx
	// By that data you can add a where clause to your query as in: `db.Where("name = ?", ctxUser.Name)`

	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// SerializeList serialize data that is returned by the GET route
func (uc *UserCrud) SerializeList(entries []interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (serialized interface{}, err error) {
	var serializedData []fiber.Map

	for _, entry := range entries {
		user, ok := entry.(*User)

		if !ok {
			return nil, errors.New(fmt.Sprintf("Cannot serialize %T as user", entry))
		}

		// if you need services inside your serialize method, just pass the service container.
		serializedData = append(serializedData, user.Serialize())
	}

	return serializedData, nil
}

type PostUserPayload struct {
	Name string `json:"name"`
}

// MapPostEntry maps data from a POST request to a fresh entry instance
func (uc *UserCrud) MapPostEntry(entry interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (mapped gocruddy.Entry, err error) {
	user, ok := entry.(*User)

	if !ok {
		return nil, gocruddy.NewError(fiber.StatusBadRequest, err)
	}

	payload := new(PostUserPayload)

	if err := ctx.BodyParser(payload); err != nil {
		return nil, gocruddy.NewError(fiber.StatusBadRequest, err)
	}

	// You can also do additional access checks here or manipulate other data.
	// You have access to the fiber.Ctx, the database transaction and the service container here.

	user.Name = payload.Name

	return user, nil
}

type PutUserPayload struct {
	PostUserPayload
}

// MapPutEntry maps data from a PUT request to a fresh entry instance
func (uc *UserCrud) MapPutEntry(entry interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (mapped gocruddy.Entry, err error) {
	user, ok := entry.(*User)

	if !ok {
		return nil, gocruddy.NewError(fiber.StatusBadRequest, err)
	}

	payload := new(PutUserPayload)

	if err := ctx.BodyParser(payload); err != nil {
		return nil, gocruddy.NewError(fiber.StatusBadRequest, err)
	}

	// You can also do additional access checks here or manipulate other data.
	// You have access to the fiber.Ctx, the database transaction and the service container here.

	user.Name = payload.Name

	return user, nil
}
