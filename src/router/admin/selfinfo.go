// @Title selfinfo.go
// @Description 关于 管理员查看自己信息功能 的实现
// @Author 矫晓佳 ${DATE} ${TIME}

package admin

import (
	"SE/src/database"
	"SE/src/interface/admin/index"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SelfInfo(c *gin.Context) {
	userName := c.Request.Header.Get("UserName")

	ret := database.GetSelfInfo(userName)
	res := index.SelfInfoResult{
		Success: true,
		Data: struct {
			UserName string `json:"userName"`
			Password string `json:"password"`
			Avatar   string `json:"avatar"`
		}(struct {
			UserName string
			Password string
			Avatar   string
		}{UserName: ret.UserName, Password: ret.Password, Avatar: ret.Avatar}),
	}

	c.IndentedJSON(http.StatusOK, res)
}
