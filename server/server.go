package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gucarletto/go-gin-rest-boilerplate/routes"
)

func Start() {
	router := gin.Default()

	routes.AddRootRoutes(router)
	routes.AddRoutesV1(router)

	router.Run(":4242")
}
