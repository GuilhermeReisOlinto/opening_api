package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	router := gin.Default()

	routes(router)

	router.Run(":8080")
}
