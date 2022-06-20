package Router

import (
	"SE/src/middleware"
	openPackage "SE/src/router/open"

	"github.com/gin-gonic/gin"
)

func MainRouter(router *gin.Engine) {

	// router group
	admin := router.Group("/admin")
	open := router.Group("/open")
	user := router.Group("/user")

	admin.Use(middleware.TokenCheck())
	user.Use(middleware.TokenCheck())

	open.POST("/login", openPackage.Login)

	router.Run(":8080")
}
