package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/precioussilence/go-gin-starter/internal/handler"
)

type Routers struct {
	user *UserRouter
}

func NewRouters(handlers *handler.Handlers) *Routers {
	return &Routers{
		user: NewUserRouter(handlers.User),
	}
}

func (r *Routers) SetupRouters(engine *gin.Engine) {
	// Ping test
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := engine.Group("/api/v1")
	// register user apis
	r.user.Register(api)
	// register other apis group, such as role, dept and more
}
