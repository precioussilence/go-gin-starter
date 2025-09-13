package router

import (
	"github.com/gin-gonic/gin"
	"github.com/precioussilence/go-gin-starter/internal/handler"
)

type UserRouter struct {
	userHandler *handler.UserHandler
}

func NewUserRouter(userHandler *handler.UserHandler) *UserRouter {
	return &UserRouter{
		userHandler: userHandler,
	}
}

func (ur *UserRouter) Register(r *gin.RouterGroup) {
	users := r.Group("/user")
	{
		users.GET("/:id", ur.userHandler.GetUserByID)
		// other router for user operation
	}
}
