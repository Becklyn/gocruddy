package gocruddy

import (
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
)

func Get(c Container, config CrudConfig) fiber.Handler {
	log := c.GetLogger()
	db := c.GetDatabase()
	crudRepo := c.GetCrudRepo()

	return func(ctx *fiber.Ctx) error {
		if err := db.UseTransaction(func(tx *storage.Transaction) (err error) {
			entries, err := crudRepo.GetAllEntries(tx, config.CreateGetFilter(ctx), config.GetEntityEntry())

			if err != nil {
				return err
			}

			serialized, err := config.SerializeList(entries, ctx, tx)

			if err != nil {
				if crudError, ok := err.(Error); ok {
					log.ErrWarn(crudError)

					if crudError.respond {
						ctx.SendString(crudError.Error())
					}

					return ctx.SendStatus(crudError.responseCode)
				}

				return err
			}

			return ctx.JSON(serialized)
		}); err != nil {
			log.ErrError(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return nil
	}
}
