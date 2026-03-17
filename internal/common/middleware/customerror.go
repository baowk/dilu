package middleware

import (
	"dilu/internal/common/config"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core/logger"
	"github.com/gin-gonic/gin"
)

func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if c.IsAborted() {
				c.Status(200)
			}
			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "CustomError" {
					statusCode, e := strconv.Atoi(p[1])
					if e != nil {
						break
					}
					c.Status(statusCode)

					logger.Warn("request", "ip", ips.GetIP(c), "method", c.Request.Method, "path", c.Request.RequestURI,
						"query", c.Request.URL.RawQuery, "source", config.Get().Server.Name, "reqId", utils.GetReqId(c),
						"error", p[2])

					c.JSON(statusCode, gin.H{
						"code": statusCode,
						"msg":  p[2],
					})
				} else {
					logger.Error("unexpected panic", "error", errStr, "ip", ips.GetIP(c),
						"method", c.Request.Method, "path", c.Request.RequestURI)
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "Internal Server Error",
					})
				}
			case error:
				logger.Error("unexpected panic", "error", errStr.Error(), "ip", ips.GetIP(c),
					"method", c.Request.Method, "path", c.Request.RequestURI,
					"stack", string(debug.Stack()))
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
				})
			default:
				logger.Error("unexpected panic", "error", err, "ip", ips.GetIP(c),
					"method", c.Request.Method, "path", c.Request.RequestURI,
					"stack", string(debug.Stack()))
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
				})
			}
		}
	}()
	c.Next()
}
