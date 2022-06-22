package Router

import (
	"SE/src/middleware"
	adminPackage "SE/src/router/admin"
	openPackage "SE/src/router/open"
	userPackageDesktop "SE/src/router/user/desktop"
	userPackageEdit "SE/src/router/user/edit"
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
	user.POST("/newdir", userPackageDesktop.NewDir)
	user.POST("/newfile", userPackageDesktop.NewDoc)
	user.POST("/setvisibility", userPackageDesktop.SetVisibility)
	user.POST("/movedir", userPackageDesktop.MoveDir)
	user.POST("/movefile", userPackageDesktop.MoveDoc)
	user.POST("/rename", userPackageEdit.Rename)
	user.POST("/deleteitem", userPackageDesktop.DeleteItem)
	user.GET("/trash", userPackageDesktop.Trash)
	user.GET("/dircontent", userPackageDesktop.DirContent)
	user.GET("/lastview", userPackageDesktop.LastView)

	admin.GET("/userlist", adminPackage.UserList)
	admin.DELETE("/deleteuser", adminPackage.DeleteUser)
	admin.GET("/selfInfo", adminPackage.SetInfo)
	admin.POST("/setinfo", adminPackage.SetInfo)
	admin.POST("/newadmin", adminPackage.NewAdmin)
	admin.GET("/storageinfo", adminPackage.GetStorageInfo)
	//admin.GET("/userinfo",adminPackage.)
	// 需要验证 Token 的部分，在验证token以后可以按照如下方法获取 username password role
	admin.POST("/test", func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("Username"))
		fmt.Println(c.Request.Header.Get("Password"))
		fmt.Println(c.Request.Header.Get("Role"))
		c.IndentedJSON(http.StatusOK, "OK")
	})

	router.Run(":8080")
}
