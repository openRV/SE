// @Title ImportFile.go
// @Description 关于 上传文件功能 的实现
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImportFile(c *gin.Context) {

	json := make(map[string]interface{})
	c.BindJSON(&json)

	username := c.Request.Header.Get("Username")
	dirId := json["dirId"].(string)

	file, err := c.FormFile("file")
	if err != nil {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: "upload error"})
		return
	}

	fileName := file.Filename

	f, err := file.Open()
	if err != nil {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: "upload error"})
		return
	}
	defer f.Close()
	buf := make([]byte, file.Size)
	var chunk []byte
	for {
		n, err := f.Read(buf)
		if err != nil && err.Error() != "EOF" {
			fmt.Println(err)
			c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: "upload error"})
			return
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}

	var info database.ImportFileInfo
	info.DirId = dirId
	info.File = chunk
	info.FileName = fileName
	info.Username = username

	res := database.ImportFile(info)
	if !res.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: res.Msg})
	}

	var result desktop.ImportFileResult
	result.Success = true
	result.Data.DocsId = res.Id
	result.Data.DocsName = res.Name

	c.IndentedJSON(http.StatusOK, result)
}
