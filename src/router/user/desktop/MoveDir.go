package desktop

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MoveDir(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	dirId := json["dirId"].(string)
	moveTo := json["moveTo"].(string)
	username := c.Request.Header.Get("Username")

	var info database.MoveDirInfo

	info.Id = dirId
	info.Username = username
	info.MoveTo = moveTo

	err := database.MoveDir(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.MoveDirResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
