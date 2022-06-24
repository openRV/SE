// @Title newadmin.go
// @Description 关于 管理员新建管理员功能 的实现
// @Author 矫晓佳 ${DATE} ${TIME}

package admin

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"SE/src/interface/admin/index"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAdmin(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	userName := json["userName"].(string)
	password := json["password"].(string)
	if database.SearchUser(database.User{Username: userName, Password: ""}).Exist {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "user name already exit"})
		return
	}

	ret := database.RegisterAdmin(userName, password)
	if ret.Success {
		c.IndentedJSON(http.StatusOK, index.SetInfoResult{Success: true})
	}
}
