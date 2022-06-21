package admin

import (
	"SE/src/Interface/admin/index"
	"SE/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserList(c *gin.Context) {

	curPage, _ := strconv.Atoi(c.Query("curPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if curPage < 1 {
		curPage = 1
	}

	ret := database.GetAllUser()
	var res index.UserListResult
	res.Total = len(ret.Data)
	if curPage*pageSize > len(ret.Data) {
		res.Data = ret.Data[(curPage-1)*pageSize:]
	} else {
		res.Data = ret.Data[(curPage-1)*pageSize : curPage*pageSize]
	}

	c.IndentedJSON(http.StatusOK, res)
}
