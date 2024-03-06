package gocruddy

import (
	"errors"

	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Delete(c Container, config CrudConfig) fiber.Handler {
	log := c.GetLogger()
	db := c.GetDatabase()
	crudRepo := c.GetCrudRepo()

	return func(ctx *fiber.Ctx) error {
		entryId, err := ctx.ParamsInt("id")

		if err != nil {
			log.ErrWarn(err)
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		if err := db.UseTransaction(func(tx *storage.Transaction) (err error) {
			entry, err := crudRepo.GetByID(tx, uint(entryId), config.CreateDeleteFilter(ctx), config.GetEntityEntry())

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					log.Info("Did not find entry of type %T with id %d", config.GetEntityEntry(), entryId)
					return ctx.SendStatus(fiber.StatusNotFound)
				}

				return err
			}

			if err := crudRepo.Remove(tx, entry); err != nil {
				return err
			}

			return ctx.SendStatus(fiber.StatusOK)
		}); err != nil {
			log.ErrError(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return nil
	}
}
