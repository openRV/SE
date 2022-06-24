package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserSearch(c *gin.Context) {
	// parse request
	username := c.Request.Header.Get("Username")
	searchType := c.Query("searchType")
	searchContent := c.Query("searchContent")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	if curPage < 1 {
		curPage = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	var info database.UserSearchInfo
	info.SearchContent = searchContent
	info.SearchType = searchType
	info.Username = username
	res := database.UserSearch(info)
	if !res.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: res.Msg})
		return
	}

	num := len(res.Data)
	var result desktop.UserSearchResult
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
		var data desktop.DocsListItem
		data.Author = res.Data[i].Author
		data.CreateDate = res.Data[i].CreateDate
		data.DocsId = res.Data[i].DocsId
		data.DocsName = res.Data[i].DocsName
		data.DocsType = res.Data[i].DocsType
		data.LastView = res.Data[i].LastUpdate
		result.Data = append(result.Data, data)
	}

	c.IndentedJSON(http.StatusOK, result)

}
