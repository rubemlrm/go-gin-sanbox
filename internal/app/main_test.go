package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rubemlrm/go-gin-sanbox/config"
	"github.com/Rubemlrm/go-gin-sanbox/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	cfg := config.Config{
		TrustedProxies: "0.0.0.0",
		Address:        ":8080",
	}

	router, _ := app.SetupRouter(cfg)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
