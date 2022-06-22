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

	res.Data.UsingStorage = float32(database.GetAllDocSize() / 1048576)
	res.Data.TotalStorage = 100

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
