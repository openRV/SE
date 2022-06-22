package desktop

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmptyTrash(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	username := c.Request.Header.Get("Username")

	var info database.EmptyTrashInfo

	info.Username = username

	err := database.EmptyTrash(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.EmptyTrashResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}