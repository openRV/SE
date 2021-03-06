// @Title storageinfo.go
// @Description 关于 管理员查看系统存储数据功能 的实现
// @Author 矫晓佳 ${DATE} ${TIME}

package admin

import (
	"SE/src/database"
	"SE/src/interface/admin/index"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStorageInfo(c *gin.Context) {
	var res index.StorageInfoResult

	//res.Data.UsingStorage = float32(database.GetAllDocSize() / 1048576)
	res.Data.UsingStorage = 1.25
	res.Data.TotalStorage = 10

	ret := database.GetIncrease()
	if !ret.Success {
		fmt.Println(ret.Msg)
		c.IndentedJSON(http.StatusOK, ret.Msg)
	}
	res.Data.D_increase = ret.D_Data
	res.Data.M_increase = ret.M_Data
	res.Success = true
	c.IndentedJSON(http.StatusOK, res)
}
