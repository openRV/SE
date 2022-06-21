package desktop

import (
	"SE/src/database"
	comInterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetVisibility(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	vis := json["visibility"].(string)
	docId := json["docsId"].(string)
	username := c.Request.Header.Get("Username")

	var info database.SetVisInfo

	info.Id = docId
	info.Username = username
	info.Vis = vis

	err := database.SetVisibility(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.SetVisibilityResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
