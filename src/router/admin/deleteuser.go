// @Title deleteuser.go
// @Description 关于 管理员删除用户功能 的实现
// @Author 矫晓佳 ${DATE} ${TIME}

package admin

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"SE/src/interface/admin/index"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO:Debug need
func DeleteUser(c *gin.Context) {

	json := make(map[string]interface{})
	c.BindJSON(&json)

	userName := json["userName"].(string)

	if userName == "" {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "empty username"})
		return
	}

	ret := database.DeteleUser(userName)
	if !ret.Success {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{
				Success: false,
				Msg:     ret.Msg,
			})
		return
	}

	c.IndentedJSON(http.StatusOK, index.DeleteUserResult{Success: true})
}
