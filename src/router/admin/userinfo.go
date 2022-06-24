// @Title userinfo.go
// @Description 关于 管理员查看用户信息功能 的实现
// @Author 矫晓佳 ${DATE} ${TIME}

package admin

import (
	"SE/src/database"
	"SE/src/interface/admin/index"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {

	userNumRet := database.GetUserNum()
	newUserNumRet := database.GetNewUserNum()
	activeUserNumRet := database.GetActiveUserNum()

	if !(userNumRet.Success && newUserNumRet.Success && activeUserNumRet.Success) {
		c.IndentedJSON(http.StatusOK, userNumRet.Msg+"\n"+newUserNumRet.Msg+"\n"+activeUserNumRet.Msg)
	}
	for _, element := range newUserNumRet.M_data {
		fmt.Println(element.Month, element.Num)
	}
	c.IndentedJSON(http.StatusOK, index.UserInfoResult{
		Success: true,
		Data: index.UserInfoData{
			UserNumbers:    userNumRet.UserNum,
			OnlineNumbers:  activeUserNumRet.Num,
			D_UserIncrease: newUserNumRet.D_data,
			M_UserIncrease: newUserNumRet.M_data}})
}
