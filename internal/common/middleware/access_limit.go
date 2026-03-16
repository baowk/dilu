package middleware

import (
	"dilu/internal/common/config"
	"sync"
	"time"

	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/gin-gonic/gin"
)

type Access struct {
	beginTime time.Time
	accessCnt int
}

var (
	accessMap = make(map[string]*Access, 0)
	accessMu  sync.Mutex
)

func AccessLimitfunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := ips.GetIP(c)
		blocked := false

		accessMu.Lock()
		v, ok := accessMap[ip]
		if !ok {
			accessMap[ip] = &Access{beginTime: time.Now(), accessCnt: 1}
		} else {
			curT := time.Now()
			if curT.Sub(v.beginTime) > config.Get().AccessLimit.Duration {
				v.accessCnt = 1
				v.beginTime = curT
			} else if v.accessCnt > config.Get().AccessLimit.GetTotal() {
				v.accessCnt++
				if v.accessCnt/config.Get().AccessLimit.GetTotal() > 1 {
					v.beginTime = curT
				}
				blocked = true
			} else {
				v.accessCnt++
			}
		}
		accessMu.Unlock()

		if blocked {
			Fail(c, 429, "too many requests")
			return
		}
		c.Next()
	}
}
