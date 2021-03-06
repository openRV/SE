// @Title openSearch.go
// @Description 关于 搜索公开文章功能 的实现
// @Author 杜沛然 ${DATE} ${TIME}

package open

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"SE/src/interface/open/index"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OpenSearch(c *gin.Context) {

	var searchInfo database.OpenDocSearchInfo

	searchInfo.Title = c.Query("searchTitle")
	searchInfo.Author = c.Query("searchAuthor")
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	if curPage < 1 {
		curPage = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize < 1 {
		pageSize = 15
	}

	var result index.OpenSearchResult

	docs, err := database.OpenSearch(searchInfo)
	if err != "" {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{
				Success: false,
				Msg:     err,
			})
		return
	}
	num := len(docs)
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

	result.Success = true
	for i := from; i < to; i++ {
		doc := docs[i]
		var docData index.OpenSearchData
		docData.DocsId = doc.DocsId
		docData.DocsName = doc.DocsName
		docData.Author = doc.Author
		docData.LastUpdate = doc.LastUpdate
		docData.ViewCounts = doc.ViewCounts
		result.Data = append(result.Data, docData)
	}

	c.IndentedJSON(http.StatusOK, result)

}
