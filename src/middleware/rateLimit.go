// @Title rateLimit.go
// @Description 用于控制流量的中间件
// @Author 杜沛然 ${DATE} ${TIME}

package middleware

// 中间件
// 用于控制流量
// 采用令牌桶方法

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

//@title func RateLimit
//@description 返回一个函数闭包，用于利用令牌桶原理限制服务器 并发/单位时间并发 协程的数量
//@param fillInterval time.Duration 每格多长时间向令牌桶增加令牌
//@param cap int64 令牌桶的最大容量
//@param quantum int64 每次向令牌桶新增的令牌数量
//@result func gin.HandlerFunc 用于限制并发量的事务函数

func RateLimit(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit ...")
			c.Abort()
			return
		}
		c.Next()
	}
}
