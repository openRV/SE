package desktop

import (
	"SE/src/database"
	comInterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MoveDoc(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	docsId := json["docsId"].(string)
	moveTo := json["moveTo"].(string)
	username := c.Request.Header.Get("Username")

	var info database.MoveDocInfo

	info.Id = docsId
	info.Username = username
	info.MoveTo = moveTo

	err := database.MoveDoc(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.MoveFileResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
