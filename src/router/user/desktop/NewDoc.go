package desktop

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewDoc(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	dirName := json["docsName"].(string)
	fatherId := json["fatherDirId"].(string)
	username := c.Request.Header.Get("Username")

	var info database.NewDocInfo

	info.FatherDirId = fatherId
	info.DocName = dirName
	info.Username = username

	err := database.NewDoc(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.NewFileResult
	res.Success = true
	res.Data.DocsId = err.Id

	c.IndentedJSON(http.StatusOK, res)

}
