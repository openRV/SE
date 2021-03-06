// @Title setinfo.go
// @Description 关于 管理员更改自身信息（账户名、密码、头像）功能 的实现
// @Author 矫晓佳 ${DATE} ${TIME}

package admin

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"SE/src/interface/admin/index"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetInfo(c *gin.Context) {
	oldUserName := c.GetHeader("Username")

	json := make(map[string]interface{})
	c.BindJSON(&json)

	params := index.SetInfoParams{
		UserName: json["userName"].(string),
		Password: json["password"].(string),
		Avatar:   json["avatar"].(string),
	}
	if !database.SearchUser(database.User{Username: oldUserName, Password: ""}).Exist {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "user does not exist"})
		return
	}

	ret := database.UpadateInfo(oldUserName, params)
	if ret.Success {
		c.IndentedJSON(http.StatusOK, index.SetInfoResult{Success: true})
	}

}
