// @Title MoveDoc.go
// @Description 关于 移动文章功能 的实现
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MoveDoc(c *gin.Context) {
	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	docsId := json["docsId"].(string)
	moveTo := json["moveTo"].(string)
	username := c.Request.Header.Get("Username")

	var info database.MoveDocInfo

	info.Id = docsId
	info.Username = username
	info.MoveTo = moveTo

	err := database.MoveDoc(info)
	if !err.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: err.Msg})
		return
	}

	var res desktop.MoveFileResult
	res.Success = true

	c.IndentedJSON(http.StatusOK, res)

}
