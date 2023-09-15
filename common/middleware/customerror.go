package middleware

import (
	"net/http"
	"runtime"
	"strconv"
	"strings"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Printf("Custom error %v \n", err)
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

					core.Log.Warn("request", zap.String("ip", ips.GetIP(c)), zap.String("method", c.Request.Method), zap.String("path", c.Request.RequestURI),
						zap.String("query", c.Request.URL.RawQuery), zap.String("source", core.Cfg.Server.Name), zap.String("reqId", utils.GetReqId(c)),
						zap.String("error", p[2]))

					c.JSON(http.StatusOK, gin.H{
						"code": statusCode,
						"msg":  p[2],
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"code": 500,
						"msg":  errStr,
					})
				}
			case runtime.Error:
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  errStr.Error(),
				})
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}
