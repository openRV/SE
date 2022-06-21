package admin

import (
	"SE/src/Interface"
	"SE/src/Interface/admin/index"
	"SE/src/database"
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
