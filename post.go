package gocruddy

import (
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
)

func Post(c Container, config CrudConfig) fiber.Handler {
	log := c.GetLogger()
	db := c.GetDatabase()
	crudRepo := c.GetCrudRepo()

	return func(ctx *fiber.Ctx) error {
		if err := db.UseTransaction(func(tx *storage.Transaction) (err error) {
			mapped, err := config.MapPostEntry(config.GetEntityEntry(), ctx, tx)

			if err != nil {
				if crudError, ok := err.(CRUDError); ok {
					log.ErrWarn(crudError)
					return ctx.SendStatus(crudError.responseCode)
				}

				return err
			}

			if err := crudRepo.Insert(tx, mapped); err != nil {
				return err
			}

			return ctx.JSON(fiber.Map{
				"id": mapped.GetID(),
			})
		}); err != nil {
			log.ErrError(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return nil
	}
}
