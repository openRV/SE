package middleware

import (
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

var key string
var effectTime time.Duration

func InitToken(keyStr string, hourTime int) {
	key = keyStr
	effectTime = time.Hour * time.Duration(hourTime)
}

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
			// TODO search database for the username & password
		}
		c.Next()
	}
}

func GenToken(json map[string]interface{}) string {

	expireTime := time.Now().Add(effectTime)
	claims := &Claim{
		Name:     json["username"].(string),
		Password: json["password"].(string),
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
