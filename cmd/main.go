package main

import (
	"github.com/gin-gonic/gin"
	"github.com/precioussilence/go-gin-starter/internal/config"
	"github.com/precioussilence/go-gin-starter/internal/handler"
	"github.com/precioussilence/go-gin-starter/internal/repository"
	"github.com/precioussilence/go-gin-starter/internal/router"
	"github.com/precioussilence/go-gin-starter/internal/service"
)

func main() {
	// load global configuration
	cfg, err := config.LocalConfig()
	if err != nil {
		panic("failed to local config: " + err.Error())
	}

	// create engine instance
	engine := gin.Default()

	// init repositories
	repositories := repository.NewRepositories(cfg)
	// init services
	services := service.NewServices(repositories)
	// init handlers
	handlers := handler.NewHandlers(services)
	// init routers
	routers := router.NewRouters(handlers)
	routers.SetupRouters(engine)

	// Listen and Server in 0.0.0.0:${port}
	engine.Run(":" + cfg.Server.Port)
}
