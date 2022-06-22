package admin

import (
	"SE/src/Interface/admin/index"
	"SE/src/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetStorageInfo(c *gin.Context) {
	var res index.StorageInfoResult

	res.Data.UsingStorage = strconv.Itoa(database.GetAllDocSize() / 1048576)
	res.Data.TotalStorage = "100"

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
