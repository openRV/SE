package admin

import (
	"SE/src/database"
	"SE/src/interface/admin/index"
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

	c.IndentedJSON(http.StatusOK, index.UserInfoResult{
		Success: true,
		Data: index.UserInfoData{
			UserNumbers:    userNumRet.UserNum,
			OnlineNumbers:  activeUserNumRet.Num,
			D_UserIncrease: newUserNumRet.D_data,
			M_UserIncrease: newUserNumRet.M_data}})
}
