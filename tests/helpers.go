package tests

import (
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

func PrepareRequest(method string, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	if body == "" {
		body = "{}"
	}

	req := httptest.NewRequest(method, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}
