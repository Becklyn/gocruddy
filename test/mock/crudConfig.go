package mock

import (
	"github.com/Becklyn/go-cruddy"
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type CrudConfig struct {
	mock.Mock
}

func (cc *CrudConfig) UseContainer(c gocruddy.Container) {
	cc.Called(c)
}

func (cc *CrudConfig) GetEntityEntry() interface{} {
	return cc.Called().Get(0)
}

func (cc *CrudConfig) GetBasePath() string {
	return cc.Called().String(0)
}

func (cc *CrudConfig) CreateGetFilter(ctx *fiber.Ctx) gocruddy.DatabaseFilter {
	return cc.Called(ctx).Get(0).(func(gorm *gorm.DB) *gorm.DB)
}

func (cc *CrudConfig) CreateUpdateFilter(ctx *fiber.Ctx) gocruddy.DatabaseFilter {
	return cc.Called(ctx).Get(0).(func(gorm *gorm.DB) *gorm.DB)
}

func (cc *CrudConfig) CreateDeleteFilter(ctx *fiber.Ctx) gocruddy.DatabaseFilter {
	return cc.Called(ctx).Get(0).(func(gorm *gorm.DB) *gorm.DB)
}

func (cc *CrudConfig) SerializeList(entries []interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (serialized fiber.Map, err error) {
	args := cc.Called(entries, ctx)
	return args.Get(0).(fiber.Map), args.Error(1)
}

func (cc *CrudConfig) MapPostEntry(entry interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (mapped gocruddy.Entry, err error) {
	args := cc.Called(entry, ctx, tx)
	return args.Get(0).(gocruddy.Entry), args.Error(1)
}

func (cc *CrudConfig) MapPutEntry(entry interface{}, ctx *fiber.Ctx, tx *storage.Transaction) (mapped gocruddy.Entry, err error) {
	args := cc.Called(entry, ctx, tx)
	return args.Get(0).(gocruddy.Entry), args.Error(1)
}
