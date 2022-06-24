package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteItem(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	docsId := json["docsId"].(string)
	isDir := json["isDir"].(bool)
	username := c.Request.Header.Get("Username")

	var info database.DeleteItemInfo

	info.Id = docsId
	info.Username = username
	info.IsDir = isDir

	err := database.DeleteItem(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.DeleteItemResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
