package log

import (
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ModelName = "GIN"
	UseRoot   = true
)

func GinLog(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	if !UseRoot {
		if path == "/" {
			c.AbortWithStatus(200)
			return
		}
	}

	c.Next()

	end := time.Now()
	latency := end.Sub(start)

	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()

	comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
	if raw != "" {
		path = path + "?" + raw
	}

	if method != "OPTIONS" {
		Debugf("[%s] %v  %3d  %13v | %15s |  %s  %s\n%s",
			ModelName,
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
			comment)
	}
}
