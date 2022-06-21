package desktop

import (
	"SE/src/database"
	comInterface "SE/src/interface"
	"SE/src/interface/user/desktop"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDir(c *gin.Context) {

	username := c.Request.Header.Get("Username")
	if username == "" {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: "empty username"})
		return
	}

	dir := database.UserDir(username, true)
	if !dir.Success {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: dir.Msg})
		return
	}

	var resDir desktop.Dir

	resDir.DirId = username
	resDir.DirName = dir.Name
	fillinDir(dir.Data, resDir.Subdir)

	res := []desktop.Dir{resDir}

	c.IndentedJSON(http.StatusOK, desktop.UserDirResult{Success: true, Data: res})

}

func fillinDir(dir []database.Dir, resDir []desktop.Dir) {
	if len(dir) == 0 {
		return
	}
	var newDir desktop.Dir
	for i := 0; i < len(dir); i = i + 1 {
		newDir.DirId = dir[i].Id
		newDir.DirName = dir[i].Name
		fillinDir(dir[i].Subdir, newDir.Subdir)
		resDir = append(resDir, newDir)
	}
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
