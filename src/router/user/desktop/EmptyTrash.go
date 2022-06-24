package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmptyTrash(c *gin.Context) {
	// parse request

	username := c.Request.Header.Get("Username")

	var info database.EmptyTrashInfo

	info.Username = username

	err := database.EmptyTrash(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.EmptyTrashResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
