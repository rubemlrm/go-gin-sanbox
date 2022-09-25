package app

import (
	"github.com/Rubemlrm/go-gin-sanbox/config"
	"github.com/Rubemlrm/go-gin-sanbox/internal/controller"
	"github.com/gin-gonic/gin"
)

// StartApp Exported to deploy the gin server to main entrypoint
func StartApp() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	router := SetupRouter(cfg)
	router.Run(cfg.Address)
}

// SetupRouter ...
func SetupRouter(cfg config.Config) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{cfg.TrustedProxies})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/v1")
	{
		h := new(controller.HomeController)
		v1.GET("hello", h.Get)
	}
	return router
}
