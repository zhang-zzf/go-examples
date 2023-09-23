package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApiV1Videos(t *testing.T) {
	// given
	router := setupRouter()
	// when
	w := httptest.NewRecorder()
	url := "/api/v1/videos?sort_by=-popularity&limit=24&offset=48"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)
	// then
	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
func TestPingRoute(t *testing.T) {
	// given
	router := setupRouter()
	// when
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
