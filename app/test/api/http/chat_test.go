package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"app/src/web/http/app"
	"github.com/stretchr/testify/assert"
)

func TestGetChatByID(t *testing.T) {
	srv := app.NewApplication()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/chat/1", nil)
	srv.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetChatMessages(t *testing.T) {
	srv := app.NewApplication()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/chat/1/messages", nil)
	srv.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, "[]", rec.Body.String())
}
