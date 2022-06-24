// @Title logger.go
// @Description
// @Author 杜沛然 ${DATE} ${TIME}

package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newLogger(file *os.File) *logrus.Logger {

	logger := logrus.New()

	logger.Out = file
	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}

func Logger(file *os.File) gin.HandlerFunc {
	logger := newLogger(file)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
