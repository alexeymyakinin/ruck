package api

import (
	"encoding/json"
	"github.com/alexeymyakinin/ruck/app/src/api/http/app"
	"github.com/alexeymyakinin/ruck/app/src/api/http/schema"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCreateUser(t *testing.T) {
	raw, _ := json.Marshal(&schema.UserCreateRequest{Email: "test@domain.com", Username: "test", Password: "1234"})
	str := string(raw)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", strings.NewReader(str))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	srv := app.NewApplication()
	srv.ServeHTTP(res, req)

	require.Equal(t, http.StatusCreated, res.Code, res.Body)
}
