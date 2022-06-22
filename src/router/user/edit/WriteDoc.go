package edit

import (
	"SE/src/Interface/user/edit"
	"SE/src/database"
	comInterface "SE/src/interface"
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
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: res.Msg})
		return
	}

	var result edit.WriteDocsResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, result)
}
