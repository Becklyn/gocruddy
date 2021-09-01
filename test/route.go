package test

import (
	"bytes"
	"encoding/json"
	"github.com/Becklyn/go-cruddy/test/mock"
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func ExpectStatus(t *testing.T, app *fiber.App, req *http.Request, status int) {
	resp, _ := app.Test(req)
	assert.Equal(t, status, resp.StatusCode)
}

type RouteTest struct {
	app    *fiber.App
	locals map[string]interface{}
}

// New constructor
func New() *RouteTest {
	return &RouteTest{
		app:    fiber.New(),
		locals: make(map[string]interface{}),
	}
}

// SetLocal sets a local variable for the test request context
func (r *RouteTest) SetLocal(key string, value interface{}) {
	r.locals[key] = value
}

// RouteTestHandler handler func for a route test
type RouteTestHandler func(ctx *fiber.Ctx, tx *storage.Transaction)

// Test a route action
func (r *RouteTest) Test(handler RouteTestHandler, data *fiber.Map) {
	tx, rollback := mock.Transaction()
	defer rollback()

	wg := sync.WaitGroup{}
	r.app.Post("/", func(ctx *fiber.Ctx) error {
		for key, value := range r.locals {
			ctx.Locals(key, value)
		}

		handler(ctx, tx)

		wg.Done()
		return nil
	})

	wg.Add(1)
	postJSON(r.app, "/", data)
	wg.Wait()
}

func postJSON(app *fiber.App, path string, data *fiber.Map) *http.Response {
	payloadBytes, _ := json.Marshal(data)

	req := httptest.NewRequest("POST", "http://example.com"+path, bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	return resp
}
