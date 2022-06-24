// @Title admin.go
// @Description 用于管理员身份认证的中间件
// @Author 杜沛然 ${DATE} ${TIME}

package middleware

import (
	Interface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Role") == "admin" {
			c.Next()
		} else {
			c.IndentedJSON(http.StatusOK,
				Interface.ErrorRes{Success: false, Msg: "admin only"})
		}
	}
}
