package open

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"SE/src/interface/open/index"
	"SE/src/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	// parse request
	json := make(map[string]interface{})
	c.BindJSON(&json)

	username := json["userName"].(string)
	password := json["password"].(string)

	if username == "" {
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "empty username"})
		return
	}

	// search database for username and password
	user := database.SearchUser(database.User{Username: username, Password: password})

	if !user.Exist {
		// when user does not exist
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "user does not exist"})
		return
	} else if !user.Password {
		// when user's password incorrect
		c.IndentedJSON(http.StatusOK,
			Interface.ErrorRes{Success: false, Msg: "password incorrect"})
	}

	token := middleware.GenToken(username, password)

	var ret index.LoginResult

	ret.Success = true
	ret.Data.Role = user.Role
	ret.Data.Token = token
	ret.Data.Name = username
	ret.Data.Avatar = user.Avatar

	c.IndentedJSON(http.StatusOK, ret)

}
