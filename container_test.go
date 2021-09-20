package gocruddy_test

import (
	"github.com/Becklyn/gocruddy"
	"github.com/Becklyn/gocruddy/test/mock"
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
