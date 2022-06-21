package admin

import (
	"SE/src/Interface"
	"SE/src/Interface/admin/index"
	"SE/src/database"
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
