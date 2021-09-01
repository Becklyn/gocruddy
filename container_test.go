package gocruddy_test

import (
	gocruddy "github.com/Becklyn/go-cruddy"
	"github.com/Becklyn/go-cruddy/test/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrudContainer(t *testing.T) {
	c := gocruddy.CrudContainer{}

	assert.Len(t, c.GetCrudConfigs(), 0)

	c.UseCrudConfig(&mock.CrudConfig{})
	c.UseCrudConfig(&mock.CrudConfig{})
	c.UseCrudConfig(&mock.CrudConfig{})

	assert.Len(t, c.GetCrudConfigs(), 3)
}
