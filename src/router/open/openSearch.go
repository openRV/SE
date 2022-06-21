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

	var searchInfo database.DocSearchInfo

	searchInfo.Content = c.Query("searchContent")
	searchInfo.Type = c.Query("searchType")
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
		var docStruct index.OpenSearchResult
		docStruct.Data[0].DocsId = doc.DocsId
		docStruct.Data[0].DocsName = doc.DocsName
		docStruct.Data[0].Author = doc.Author
		docStruct.Data[0].LastUpdate = doc.LastUpdate
		docStruct.Data[0].ViewCounts = doc.ViewCounts
		result.Data = append(result.Data, docStruct.Data[0])
	}

	c.IndentedJSON(http.StatusOK, result)

}
