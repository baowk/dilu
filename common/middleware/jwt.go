package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/baowk/dilu-core/common/utils/token"
	"github.com/baowk/dilu-core/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtHandler(jwtCfg config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := token.GetAccessToken(authorization)
		if err != nil {
			Fail(c, 555, err.Error())
			return
		}
		customClaims := new(token.CustomClaims)
		// 解析Token
		err = token.Parse(accessToken, customClaims, jwtCfg.SignKey, jwt.WithSubject(jwtCfg.Subject))
		if err != nil || customClaims == nil {
			Fail(c, 555, err.Error())
			return
		}

		exp, err := customClaims.GetExpirationTime()
		// 获取过期时间返回err,或者exp为nil都返回错误
		if err != nil || exp == nil {
			Fail(c, 555, err.Error())
			return
		}

		// 刷新时间大于0则判断剩余时间小于刷新时间时刷新Token并在Response header中返回
		if jwtCfg.Refresh > 0 {
			now := time.Now()
			diff := exp.Time.Sub(now)
			refreshTTL := time.Duration(jwtCfg.Refresh) * time.Minute
			fmt.Println(diff.Seconds(), refreshTTL)
			if diff < refreshTTL {
				exp := time.Now().Add(time.Duration(jwtCfg.Expires) * time.Minute)
				newToken, _ := token.Refresh(customClaims, jwtCfg.SignKey)
				c.Writer.Header().Set("refresh-access-token", newToken)
				c.Writer.Header().Set("refresh-exp", strconv.FormatInt(exp.Unix(), 10))
			}
		}

		c.Set("a_uid", customClaims.UserID)
		c.Set("a_mobile", customClaims.Mobile)
		c.Set("a_nickname", customClaims.Nickname)
		c.Set("jwt_data", customClaims.JwtData)
		c.Next()
	}
}

func refresh(customClaims token.CustomClaims, exp time.Time, jwtCfg config.JWT) (string, error) {
	jd := token.JwtData{
		UserID:   customClaims.UserID,
		Nickname: customClaims.Nickname,
		Mobile:   customClaims.Mobile,
	}
	cc := token.NewCustomClaims(jd, exp, jwtCfg.Issuer, jwtCfg.Subject)
	return token.Refresh(cc, jwtCfg.SignKey)
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
