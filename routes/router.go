package routes

import "github.com/gin-gonic/gin"

func GetRoutesV1(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("v1")

	v1.GET("users/:id")
	v1.GET("users")
	v1.POST("users")
	v1.PUT("users")

	v1.GET("transactions/:id")
	v1.GET("transactions")
	v1.POST("transactions")

	return v1
}
