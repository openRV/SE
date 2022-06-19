package middleware

import (
	"github.com/gin-gonic/gin"
)

func SafetyCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
