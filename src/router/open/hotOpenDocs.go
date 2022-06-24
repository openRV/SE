// @Title hotOpenDocs.go
// @Description 关于 查看热门公开文章功能 的实现
// @Author 杜沛然 ${DATE} ${TIME}

package open

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"SE/src/interface/open/index"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HotOpenDocs(c *gin.Context) {

	var result index.HotdocsResult

	docs, err := database.HotOpenDocs()
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

	if num > 10 {
		num = 10
	}

	result.Success = true
	result.Total = num
	for i := 0; i < num; i++ {
		doc := docs[i]
		var docData index.HotdocsData
		docData.Id = doc.DocsId
		docData.Title = doc.DocsName
		docData.DownloadNum = doc.ViewCounts
		result.Data = append(result.Data, docData)
	}

	c.IndentedJSON(http.StatusOK, result)

}
