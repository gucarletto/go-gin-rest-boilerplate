package routes

import "github.com/gin-gonic/gin"

func AddRootRoutes(router *gin.Engine) {
	root := router.Group("/")
	root.GET("status", getStatus)
}

func AddRoutesV1(router *gin.Engine) {
	v1 := router.Group("v1")

	v1.POST("/auth")

	v1.GET("users/:id")
	v1.GET("users")
	v1.POST("users")
	v1.PUT("users")

	v1.GET("transactions/:id")
	v1.GET("transactions")
	v1.POST("transactions")
}

func getStatus(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
