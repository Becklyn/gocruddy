package mock

import (
	"github.com/Becklyn/go-cruddy"
	"github.com/ao-concepts/logging"
	"github.com/stretchr/testify/mock"
)

type Container struct {
	mock.Mock
}

func (c *Container) GetLogger() logging.Logger {
	return c.Called().Get(0).(logging.Logger)
}

func (c *Container) GetDatabase() gocruddy.Database {
	return c.Called().Get(0).(gocruddy.Database)
}

func (c *Container) GetCrudConfigs() []gocruddy.CrudConfig {
	return c.Called().Get(0).([]gocruddy.CrudConfig)
}

func (c *Container) GetCrudRepo() gocruddy.CrudRepository {
	return c.Called().Get(0).(gocruddy.CrudRepository)
}
