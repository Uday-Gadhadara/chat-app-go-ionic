package router

import (
	"chat/backend/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/check", controller.CheckUser)
}
