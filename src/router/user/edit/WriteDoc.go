package edit

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/edit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteDoc(c *gin.Context) {
	// parse request

	json := make(map[string]interface{})
	c.BindJSON(&json)

	docsId := json["docsId"].(string)
	docsContent := json["docsContent"].(string)
	username := c.Request.Header.Get("Username")

	var info database.WriteDocsInfo

	info.Id = docsId
	info.Username = username
	info.Content = docsContent

	res := database.WriteDocs(info)
	if !res.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: res.Msg})
		return
	}

	var result edit.WriteDocsResult
	result.Success = true

	c.IndentedJSON(http.StatusOK, result)
}
