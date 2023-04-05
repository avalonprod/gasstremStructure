package user

import "github.com/gin-gonic/gin"

type userController struct {
	router  *gin.RouterGroup
	service *userService
}

func NewUserController(router *gin.RouterGroup, userService *userService) *userController {
	return &userController{
		router:  router,
		service: userService,
	}
}
