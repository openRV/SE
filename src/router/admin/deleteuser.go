package admin

import (
	"SE/src/Interface"
	"SE/src/Interface/admin/index"
	"SE/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUserour(c *gin.Context) {

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
