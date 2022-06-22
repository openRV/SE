package desktop

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewDir(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	dirName := json["dirName"].(string)
	fatherId := json["fatherDirId"].(string)
	username := c.Request.Header.Get("Username")

	var info database.NewDirInfo

	info.FatherDirId = fatherId
	info.Name = dirName
	info.Owner = username

	err := database.NewDir(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	c.IndentedJSON(http.StatusOK, desktop.NewDirResult{Success: true})

}
