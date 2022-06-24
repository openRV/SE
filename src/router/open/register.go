// @Title register.go
// @Description 关于 注册功能 的实现
// @Author 杜沛然 ${DATE} ${TIME}

package open

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"SE/src/interface/open/index"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	username := json["userName"].(string)
	password := json["password"].(string)

	if username == "" {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "empty username"})
		return
	}

	ret := database.RegisterUser(database.User{Username: username, Password: password})
	if !ret.Success {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{
				Success: false,
				Msg:     ret.Msg,
			})
		return
	}

	c.IndentedJSON(http.StatusOK, index.RegisterResult{Success: true})

}
