package app

import (
	"github.com/Rubemlrm/go-gin-sanbox/config"
	"github.com/Rubemlrm/go-gin-sanbox/internal/controller"
	"github.com/gin-gonic/gin"
)

// StartApp Exported to deploy the gin server to main entrypoint
func StartApp() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	router, err := SetupRouter(cfg)
	if err != nil {
		return err
	}
	err = router.Run(cfg.Address)
	if err != nil {
		return err
	}
	return nil
}

// SetupRouter ...
func SetupRouter(cfg config.Config) (*gin.Engine, error) {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{cfg.TrustedProxies})

	if err != nil {
		return nil, err
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/v1")
	{
		h := new(controller.HomeController)
		v1.GET("hello", h.Get)
	}
	return router, nil
}
