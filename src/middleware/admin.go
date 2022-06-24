// @Title admin.go
// @Description 用于管理员身份认证的中间件
// @Author 杜沛然 ${DATE} ${TIME}

package middleware

import (
	Interface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

//@title func AdminOnly
//@description 返回一个函数闭包，用于判断用户的身份是否是管理员，若不是，则拦截请求
//@result func gin.HanlderFunc 用于判断并拦截不符合身份的用户的事务函数

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
