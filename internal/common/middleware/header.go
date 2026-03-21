package middleware

import (
	"net/http"
	"time"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/core/logger"
	"github.com/gin-gonic/gin"
)

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// ReqId 生成/提取请求 ID，注入到 zerolog context，下游所有日志自动携带 reqId
func ReqId(c *gin.Context) {
	reqId := utils.GetReqId(c)
	log := logger.Default().With().Str("reqId", reqId).Logger()
	c.Request = c.Request.WithContext(log.WithContext(c.Request.Context()))
	c.Next()
}

