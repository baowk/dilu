package middleware

import (
	"time"

	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
)

type Access struct {
	beginTime time.Time
	accessCnt int
}

var (
	accessMap = make(map[string]*Access, 0)
)

func AccessLimitfunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := ips.GetIP(c)
		v, ok := accessMap[ip]
		if !ok { //首次访
			accessMap[ip] = &Access{beginTime: time.Now(), accessCnt: 1}
		} else {
			curT := time.Now()
			if curT.Sub(v.beginTime) > core.Cfg.AccessLimit.Duration { //当前时间和开始时间的周期
				v.accessCnt = 1
				v.beginTime = curT
			} else if v.accessCnt > core.Cfg.AccessLimit.GetTotal() { //时间范围内数量超标
				v.accessCnt++
				if v.accessCnt/core.Cfg.AccessLimit.GetTotal() > 1 {
					v.beginTime = curT
				}
				//http.Error(c.Writer, "too many requests", http.StatusTooManyRequests)
				Fail(c, 429, "too many requests")
				return
			} else {
				v.accessCnt++
			}
		}
		c.Next()
	}
}
