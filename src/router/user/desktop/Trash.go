// @Title Trash.go
// @Description 关于 查看垃圾桶中的文件功能 的实现
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
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
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: res.Msg})
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
		data.DocsName = database.GetDocsName(res.Data[i].Id).DocsName
		data.DocsId = res.Data[i].Id
		data.DeleteDate = res.Data[i].DeleteDate
		data.Author = res.Data[i].Author
		result.Data = append(result.Data, data)
	}

	c.IndentedJSON(http.StatusOK, result)

}
