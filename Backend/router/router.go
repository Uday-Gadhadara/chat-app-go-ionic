package router

import "github.com/gin-gonic/gin"

func IndexRouter(r *gin.RouterGroup) {

	user := r.Group("/user")
	{
		UserRoutes(user)
	}
}
