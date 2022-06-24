// @Title UserDir.go
// @Description 关于 查看属于当前登录用户的文件夹功能 的实现
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

import (
	"SE/src/database"
	cominterface "SE/src/interface"
	"SE/src/interface/user/desktop"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDir(c *gin.Context) {

	username := c.Request.Header.Get("Username")
	if username == "" {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: "empty username"})
		return
	}

	dir := database.UserDir(username, true)
	if !dir.Success {
		c.IndentedJSON(http.StatusOK, cominterface.ErrorRes{Success: false, Msg: dir.Msg})
		return
	}

	var resDir desktop.Dir

	resDir.DirId = username
	resDir.DirName = dir.Name
	resDir.Subdir = fillinDir(dir.Data[0].Subdir)

	res := []desktop.Dir{resDir}

	c.IndentedJSON(http.StatusOK, desktop.UserDirResult{Success: true, Data: res})

}

func fillinDir(dir []database.Dir) []desktop.Dir {

	var resDir []desktop.Dir
	if len(dir) == 0 {
		return nil
	}
	var newDir desktop.Dir
	for i := 0; i < len(dir); i = i + 1 {
		newDir.DirId = dir[i].Id
		newDir.DirName = dir[i].Name
		newDir.Subdir = fillinDir(dir[i].Subdir)
		resDir = append(resDir, newDir)
	}
	return resDir
}

//// 文件树
//type Dir struct {
//	DirId   string `json:"dirId"`
//	DirName string `json:"dirName"`
//	Subdir  []Dir  `json:"subdir"`
//}

//type Dir struct {
//	Id         string
//	Name       string
//	Owner      string
//	CreateDate string
//	LastView   string
//	Subdir []Dir
//}
