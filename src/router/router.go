package Router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainRouter(router *gin.Engine) {

	// router group
	//admin := router.Group("/admin")
	//open := router.Group("/open")
	//user := router.Group("/user")

	router.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "OK")
	})

	router.Run(":8080")
}
