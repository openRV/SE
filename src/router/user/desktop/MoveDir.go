package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
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
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.MoveDirResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
