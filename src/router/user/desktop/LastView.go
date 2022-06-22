package desktop

import (
	"SE/src/Interface/user/desktop"
	"SE/src/database"
	comInterface "SE/src/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LastView(c *gin.Context) {
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

	var info database.LastViewInfo
	info.Username = username
	res := database.LastView(info)
	if !res.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: res.Msg})
	}

	num := len(res.Data)
	var result desktop.LastViewResult
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
		var doc desktop.DocsListItem
		doc.Author = res.Data[i].Author
		doc.CreateDate = res.Data[i].CreateDate
		doc.DocsId = res.Data[i].DocsId
		doc.DocsName = res.Data[i].DocsName
		doc.DocsType = res.Data[i].DocsType
		doc.LastView = res.Data[i].LastUpdate

		result.Data = append(result.Data, doc)
	}

	c.IndentedJSON(http.StatusOK, result)

}
