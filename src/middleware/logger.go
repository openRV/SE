// @Title logger.go
// @Description 用于向指定的文件中写入制定格式的日志的函数
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

//@title func Logger
//@description 返回一个函数闭包，用于插入事务队列。该返回函数用于向指定的文件中写入制定格式的日志。
//@param file *os.File 要写入日志的文件
//@result func gin.HandlerFunc 用于向指定文件写入日志信息的闭包

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
