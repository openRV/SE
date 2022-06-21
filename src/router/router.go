package Router

import (
	"SE/src/middleware"
	openPackage "SE/src/router/open"
	userPackageDesktop "SE/src/router/user/desktop"
	"fmt"
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

	admin.Use(middleware.AdminOnly())

	open.POST("/login", openPackage.Login)
	open.POST("/register", openPackage.Register)
	open.GET("/openSearch", openPackage.OpenSearch)
	open.GET("/hotdocs", openPackage.HotOpenDocs)

	user.GET("/userdir", userPackageDesktop.UserDir)

	// 需要验证 Token 的部分，在验证token以后可以按照如下方法获取 username password role
	admin.POST("/test", func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("Username"))
		fmt.Println(c.Request.Header.Get("Password"))
		fmt.Println(c.Request.Header.Get("Role"))
		c.IndentedJSON(http.StatusOK, "OK")
	})

	router.Run(":8080")
}
