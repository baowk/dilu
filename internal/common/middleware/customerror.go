package middleware

import (
	"dilu/internal/common/codes"
	"dilu/internal/common/config"
	"errors"
	"net/http"
	"runtime/debug"

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

			// 优先匹配类型化的 AppError
			var appErr *codes.AppError
			switch v := err.(type) {
			case *codes.AppError:
				appErr = v
			case error:
				if errors.As(v, &appErr) {
					// wrapped AppError
				}
			}

			if appErr != nil {
				logger.Warn("request", "ip", ips.GetIP(c), "method", c.Request.Method, "path", c.Request.RequestURI,
					"query", c.Request.URL.RawQuery, "source", config.Get().Server.Name, "reqId", utils.GetReqId(c),
					"error", appErr.Msg)
				c.JSON(appErr.Code, gin.H{
					"code": appErr.Code,
					"msg":  appErr.Msg,
				})
				return
			}

			// 未知 panic，记录堆栈，返回通用错误
			logger.Error("unexpected panic", "error", err, "ip", ips.GetIP(c),
				"method", c.Request.Method, "path", c.Request.RequestURI,
				"stack", string(debug.Stack()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "Internal Server Error",
			})
		}
	}()
	c.Next()
}
