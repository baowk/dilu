package middleware

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LogLayout 日志layout
type LogLayout struct {
	//Metadata  map[string]interface{} // 存储自定义原数据
	//Time      time.Time
	Method    string //方法
	Path      string // 访问路径
	Query     string // 携带query
	Body      string // 携带body数据
	IP        string // ip地址
	UserAgent string // 代理
	Error     string // 错误
	Cost      string // 花费时间
	Source    string // 来源
	//Result    string        //返回数据
}

// type Logger struct {
// 	// Filter 用户自定义过滤
// 	Filter func(c *gin.Context) bool
// 	// FilterKeyword 关键字过滤(key)
// 	FilterKeyword func(layout *LogLayout) bool
// 	// AuthProcess 鉴权处理
// 	AuthProcess func(c *gin.Context, layout *LogLayout)
// 	// 日志处理
// 	Print func(LogLayout)
// 	// Source 服务唯一标识
// 	Source string
// }

// func (l Logger) SetLoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		start := time.Now()
// 		path := c.Request.URL.Path
// 		query := c.Request.URL.RawQuery
// 		var body []byte
// 		if l.Filter != nil && !l.Filter(c) {
// 			fmt.Println("aaaa")
// 			body, _ = c.GetRawData()
// 			// 将原body塞回去
// 			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
// 		}
// 		c.Next()
// 		cost := time.Since(start)
// 		layout := LogLayout{
// 			//Time:      time.Now(),
// 			Path:      path,
// 			Query:     query,
// 			IP:        c.ClientIP(),
// 			UserAgent: c.Request.UserAgent(),
// 			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
// 			Cost:      cost.String(),
// 			Source:    l.Source,
// 		}
// 		if l.Filter != nil && !l.Filter(c) {
// 			layout.Body = string(body)
// 		}
// 		if l.AuthProcess != nil {
// 			// 处理鉴权需要的信息
// 			l.AuthProcess(c, &layout)
// 		}
// 		if l.FilterKeyword != nil {
// 			// 自行判断key/value 脱敏等
// 			l.FilterKeyword(&layout)
// 		}
// 		// 自行处理日志
// 		l.Print(layout)
// 	}
// }

// func DefaultLogger() gin.HandlerFunc {
// 	return Logger{
// 		Filter: func(c *gin.Context) bool {
// 			return true
// 		},
// 		Print: func(layout LogLayout) {
// 			// 标准输出,k8s做收集
// 			v, _ := json.Marshal(layout)
// 			fmt.Println(string(v))
// 		},
// 		//Source: "",
// 	}.SetLoggerMiddleware()
// }

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				core.Log.Warn("copy body error", zap.Error(err))
				err = nil
			}
			rb, _ := io.ReadAll(bf)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
			body = string(rb)
		}

		c.Next()
		uri := c.Request.RequestURI
		// 结束时间
		if c.Request.Method == http.MethodOptions {
			return
		}

		// rt, bl := c.Get("result")
		// var result = ""
		// if bl {
		// 	rb, err := json.Marshal(rt)
		// 	if err != nil {
		// 		core.Log.Warn("json Marshal result error", zap.Error(err))
		// 	} else {
		// 		result = string(rb)
		// 	}
		// }
		cost := time.Since(startTime)

		fmt.Println(cost.Nanoseconds())

		layout := LogLayout{
			Path:      uri,
			Query:     c.Request.URL.RawQuery,
			IP:        ips.GetIP(c),
			UserAgent: c.Request.UserAgent(),
			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:      cost.String(),
			Body:      body,
			Method:    c.Request.Method,
			Source:    core.Cfg.Server.Name,
		}

		if cost.Milliseconds() < 200 {
			core.Log.Info("Req", zap.Any("", layout))
		} else {
			core.Log.Warn("Req", zap.Any("", layout))
		}

		// if c.Request.Method != "OPTIONS" && config.LoggerConfig.EnabledDB && statusCode != 404 {
		// 	SetDBOperLog(c, clientIP, statusCode, reqUri, reqMethod, latencyTime, body, result, statusBus)
		// }
	}
}
