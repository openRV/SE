package Router

import (
	"SE/src/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func MainRouter(router *gin.Engine) {

	// router group
	admin := router.Group("/admin")
	open := router.Group("/open")
	user := router.Group("/user")

	admin.Use(middleware.TokenCheck())
	user.Use(middleware.TokenCheck())

	open.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "OK")
	})

	admin.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "OK")
	})

	router.Run(":8080")
}
