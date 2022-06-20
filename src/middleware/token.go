package middleware

import (
	Interface "SE/src/interface"

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

func InitToken(keyStr string) {
	key = keyStr
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
