package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type httpTest struct {
	name               string
	req                *http.Request
	statusCodeExpected int
	BodyExpected       string
}

func getHttpTests() []httpTest {
	var req1 = httptest.NewRequest(http.MethodGet, "/api/ola", nil)
	req1.Header.Set("Authorization", "jU0IEM7RbOV8TU3srME7fojeAC63")

	var myHttpTest = []httpTest{
		{
			name:               "Test with authorization",
			req:                req1,
			statusCodeExpected: 200,
		},
	}

	return myHttpTest
}

func TestMiddleware(app *fiber.App, t *testing.T) {
	app.Get("/api/ola", func(c *fiber.Ctx) error {
		c.SendStatus(200)
		return nil
	})

	for _, req := range getHttpTests() {
		res, err := app.Test(req.req)
		if err != nil {
			t.Errorf("Error %s\nr: %s", req.name, err.Error())
		}

		if res.StatusCode != req.statusCodeExpected {
			t.Fatalf("Error %v\nr: status code not expected\nExpected %v but get %v", req.name, req.statusCodeExpected, res.StatusCode)
			t.Fail()
		}
	}
	t.Log("Teste finalizado middleware")

}
