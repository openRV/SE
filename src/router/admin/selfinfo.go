package admin

import (
	"SE/src/database"
	"SE/src/interface"
	"SE/src/interface/admin/index"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SelfInfo(c *gin.Context) {
	userName := c.GetHeader("Uaername")

	if !database.SearchUser(database.User{Username: userName, Password: ""}).Exist {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "user does not exist"})
		return
	}

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
