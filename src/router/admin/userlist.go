// @Title userlist.go
// @Description 关于 管理员查看用户列表功能 的实现
// @Author 矫晓佳 ${DATE} ${TIME}

package admin

import (
	"SE/src/database"
	"SE/src/interface/admin/index"
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
		for _, x := range ret.Data[(curPage-1)*pageSize:] {
			data := index.UserListData{
				UserName: x[0],
				Password: x[1],
			}
			res.Data = append(res.Data, data)
		}

	} else {
		for _, x := range ret.Data[(curPage-1)*pageSize : curPage*pageSize] {
			data := index.UserListData{
				UserName: x[0],
				Password: x[1],
			}
			res.Data = append(res.Data, data)
		}
	}
	res.Success = true
	c.IndentedJSON(http.StatusOK, res)
}
