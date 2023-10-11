package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"dilu/common/utils"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := GetAccessToken(authorization)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}
		customClaims := new(utils.CustomClaims)
		// 解析Token
		err = Parse(accessToken, customClaims, core.Cfg.JWT.SignKey, jwt.WithSubject(core.Cfg.JWT.Subject))
		if err != nil || customClaims == nil {
			Fail(c, 401, err.Error())
			return
		}

		exp, err := customClaims.GetExpirationTime()
		// 获取过期时间返回err,或者exp为nil都返回错误
		if err != nil || exp == nil {
			Fail(c, 401, err.Error())
			return
		}

		// 刷新时间大于0则判断剩余时间小于刷新时间时刷新Token并在Response header中返回
		if core.Cfg.JWT.Refresh > 0 {
			now := time.Now()
			diff := exp.Time.Sub(now)
			refreshTTL := time.Duration(core.Cfg.JWT.Refresh) * time.Minute
			//fmt.Println(diff.Seconds(), refreshTTL)
			if diff < refreshTTL {
				exp := time.Now().Add(time.Duration(core.Cfg.JWT.Expires) * time.Minute)
				customClaims.ExpiresAt(exp)
				newToken, _ := Refresh(customClaims, core.Cfg.JWT.SignKey)
				c.Writer.Header().Set("refresh-access-token", newToken)
				c.Writer.Header().Set("refresh-exp", strconv.FormatInt(exp.Unix(), 10))
			}
		}

		c.Set("a_uid", customClaims.UserId)
		c.Set("a_rid", customClaims.RoleId)
		c.Set("a_mobile", customClaims.Phone)
		c.Set("a_nickname", customClaims.Nickname)
		c.Set("jwt_data", customClaims.JwtData)
		c.Next()
	}
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// Refresh 刷新JWT Token
func Refresh(claims jwt.Claims, secretKey string) (string, error) {
	return utils.Generate(claims, secretKey)
}

// Parse 解析token
func Parse(accessToken string, claims jwt.Claims, secretKey string, options ...jwt.ParserOption) error {
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), err
	}, options...)
	if err != nil {
		return err
	}

	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return nil
	}

	return errors.New("Invalid Token")
}

// GetAccessToken 获取jwt的Token
func GetAccessToken(authorization string) (accessToken string, err error) {
	if authorization == "" {
		return "", errors.New("authorization header is missing")
	}

	// 检查 Authorization 头的格式
	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("invalid Authorization header format")
	}

	// 提取 Token 的值
	accessToken = strings.TrimPrefix(authorization, "Bearer ")
	return
}
