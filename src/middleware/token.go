// @Title DeleteUser.go
// @Description 中间件中关于 toke 处理的函数、数据结构
// @Author 杜沛然 ${DATE} ${TIME}

package middleware

import (
	"SE/src/database"
	Interface "SE/src/interface"
	"fmt"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claim struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var key []byte
var effectTime time.Duration

func InitToken(keyStr string, hourTime int) {
	key = []byte(keyStr)
	effectTime = time.Hour * time.Duration(hourTime)
}

//@title func TokenCheck
//@description 返回一个函数闭包，用于解析用户的 token 是否合法，用户名密码时间是否正确，并在请求头中插入用户名密码等信息
//@param key string 解析 token 所需的密钥
//@param effectTime time.Duration 令牌过期的时间
//@result func gin.HanlderFunc 用于解析 token 并判断权限的事务函数

func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get token string from request
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			// empty token
			c.IndentedJSON(http.StatusUnauthorized, Interface.ErrorRes{Success: false, Msg: "Please login first"})
			c.Abort()
		} else {
			// parse token
			Claim := &Claim{}
			token, err := jwt.ParseWithClaims(token, Claim, func(tokenString *jwt.Token) (i interface{}, err error) {
				return key, nil
			})
			if err != nil || !token.Valid {
				c.IndentedJSON(http.StatusUnauthorized, Interface.ErrorRes{Success: false, Msg: "Please login first"})
				c.Abort()
			}

			user := database.SearchUser(database.User{Username: Claim.Name, Password: Claim.Password})
			if !user.Exist {
				// not a existing user
				c.IndentedJSON(http.StatusUnauthorized, Interface.ErrorRes{Success: false, Msg: "Invalid username"})
			} else if !user.Password {
				// password incorrect
				c.IndentedJSON(http.StatusUnauthorized, Interface.ErrorRes{Success: false, Msg: "Invalid password"})
			}

			c.Request.Header.Add("Role", user.Role)
			c.Request.Header.Add("Username", Claim.Name)
			c.Request.Header.Add("Password", Claim.Password)
		}
		c.Next()
	}
}

func GenToken(username string, password string) string {

	expireTime := time.Now().Add(effectTime)
	claims := &Claim{
		Name:     username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}
