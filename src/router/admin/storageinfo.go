package admin

import (
	"SE/src/Interface/admin/index"
	"SE/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetStorageInfo(c *gin.Context) {
	var res index.StorageInfoResult
	res.Success = true
	res.Data.UsingStorage = strconv.Itoa(database.GetAllDocSize() / 1048576)
	res.Data.TotalStorage = "100"
	c.IndentedJSON(http.StatusOK, res)
}
