package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
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
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	c.IndentedJSON(http.StatusOK, desktop.NewDirResult{Success: true})

}
