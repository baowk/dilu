package middleware

import (
	"bytes"
	"dilu/internal/common/config"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core/logger"
	"github.com/gin-gonic/gin"
)

// sensitiveFields 需要在日志中过滤的敏感字段（小写）
var sensitiveFields = []string{
	"password", "passwd", "pwd", "secret", "token",
	"access_token", "refresh_token", "authorization",
	"credit_card", "card_number", "cvv", "ssn",
	"api_key", "apikey", "private_key",
}

const maxBodyLogLen = 1024

// hasSensitiveField 检查 body 中是否包含敏感字段
func hasSensitiveField(body []byte) bool {
	lower := bytes.ToLower(body)
	for _, field := range sensitiveFields {
		if bytes.Contains(lower, []byte(field)) {
			return true
		}
	}
	return false
}

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
			// 直接用 io.ReadAll，避免 buffer→writer→readAll 三次分配
			rb, err := io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Warn("read body error", "err", err)
			} else {
				c.Request.Body = io.NopCloser(bytes.NewReader(rb))
				if hasSensitiveField(rb) {
					body = "***[contains sensitive data, masked]***"
				} else if len(rb) > maxBodyLogLen {
					body = string(rb[:maxBodyLogLen]) + "...[truncated]"
				} else {
					body = string(rb)
				}
			}
		}

		c.Next()

		writeLog(startTime, body, c)
	}
}

func writeLog(startTime time.Time, body string, c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		return
	}
	cost := time.Since(startTime)

	args := []any{
		"ip", ips.GetIP(c),
		"method", c.Request.Method,
		"path", c.Request.RequestURI,
		"cost", cost,
		"query", c.Request.URL.RawQuery,
		"source", config.Get().Server.Name,
		"reqId", utils.GetReqId(c),
	}
	if body != "" {
		args = append(args, "body", body)
	}

	if cost.Milliseconds() < 200 {
		logger.Info("request", args...)
	} else {
		logger.Warn("request", args...)
	}
}

// bodyNeedLog 判断是否需要记录 body 的方法（预留扩展）
func bodyNeedLog(path string) bool {
	// 可根据路径过滤，如文件上传路径跳过 body 记录
	return !strings.HasPrefix(path, "/api/upload")
}
