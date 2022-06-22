package desktop

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteItem(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	docsId := json["docsId"].(string)
	isDir := json["isDir"].(string) == "true"
	username := c.Request.Header.Get("Username")

	var info database.DeleteItemInfo

	info.Id = docsId
	info.Username = username
	info.IsDir = isDir

	err := database.DeleteItem(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.DeleteItemResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
