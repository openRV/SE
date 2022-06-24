package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DirContent(c *gin.Context) {
	// parse request
	username := c.Request.Header.Get("Username")
	dirId := c.Query("dirId")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	if curPage < 1 {
		curPage = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	var info database.DirContentInfo
	info.Id = dirId
	info.Username = username
	res := database.DirContent(info)
	if !res.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: res.Msg})
		c.Abort()
	}

	numDir := len(res.Data.Dir)
	numDoc := len(res.Data.Docs)

	num := numDir + numDoc
	var result desktop.DirContentResult
	result.Success = true
	result.Total = num

	from := (curPage - 1) * pageSize
	to := curPage * pageSize

	if from >= num {
		// if required more than exist, return null
		from = num
		to = num
	}

	if to >= num {
		// else if end is logger than exist, shorten end
		to = num
	}

	for i := from; i < to; i += 1 {

		if i >= numDir {
			result.Data.Docs = append(result.Data.Docs, res.Data.Docs[i-numDir])
		} else {
			result.Data.Dir = append(result.Data.Dir, res.Data.Dir[i])
		}

	}

	c.IndentedJSON(http.StatusOK, result)

}
