package edit

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Rename(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	docsId := json["docsId"].(string)
	newName := json["newName"].(string)
	isDir := json["isDir"].(bool)
	username := c.Request.Header.Get("Username")

	var info database.RenameInfo

	info.Id = docsId
	info.Username = username
	info.Newname = newName
	info.IsDir = isDir

	err := database.Rename(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.MoveFileResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)
}
