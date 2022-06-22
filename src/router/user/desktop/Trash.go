package desktop

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Trash(c *gin.Context) {

	// parse request
	username := c.Request.Header.Get("Username")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	if curPage < 1 {
		curPage = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	var info database.TrashInfo

	info.Username = username

	res := database.Trash(info)
	if !res.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: res.Msg})
		return
	}

	num := len(res.Data)
	var result desktop.TrashResult
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
		var data desktop.TrashData
		data.DocsId = res.Data[i].Id
		data.DeleteDate = res.Data[i].DeleteDate
		data.Author = res.Data[i].Author
		result.Data = append(result.Data, data)
	}

	c.IndentedJSON(http.StatusOK, result)

}