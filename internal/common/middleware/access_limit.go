package middleware

import (
	"dilu/internal/common/config"
	"sync"
	"sync/atomic"
	"time"

	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/gin-gonic/gin"
)

type accessEntry struct {
	count     atomic.Int64
	resetTime atomic.Int64 // unix nano
	mu        sync.Mutex   // 仅用于重置操作
}

var accessMap sync.Map // map[string]*accessEntry

func AccessLimitfunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := ips.GetIP(c)
		limit := int64(config.Get().AccessLimit.GetTotal())
		window := config.Get().AccessLimit.Duration

		val, _ := accessMap.LoadOrStore(ip, &accessEntry{})
		entry := val.(*accessEntry)

		now := time.Now().UnixNano()
		resetAt := entry.resetTime.Load()

		// 时间窗口过期，重置计数
		if now-resetAt > int64(window) {
			entry.mu.Lock()
			// double-check 避免多个 goroutine 同时重置
			if now-entry.resetTime.Load() > int64(window) {
				entry.count.Store(1)
				entry.resetTime.Store(now)
			}
			entry.mu.Unlock()
			c.Next()
			return
		}

		// 原子递增计数
		cnt := entry.count.Add(1)
		if cnt > limit {
			Fail(c, 429, "too many requests")
			return
		}

		c.Next()
	}
}
