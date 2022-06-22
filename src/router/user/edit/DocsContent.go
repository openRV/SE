package edit

import (
	"SE/src/Interface/user/edit"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocsContent(c *gin.Context) {
	// parse request

	docsId := c.Query("docsId")
	username := c.Request.Header.Get("Username")

	var info database.DocsContentInfo

	info.Id = docsId
	info.Username = username

	res := database.DocsContent(info)
	if !res.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: res.Msg})
		return
	}

	var result edit.DocContentResult
	result.Success = true
	result.Data.DocContent = res.Data

	c.IndentedJSON(http.StatusOK, res)
}
