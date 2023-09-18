package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtHandler() gin.HandlerFunc {
	//func JwtHandler(jwtCfg config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := GetAccessToken(authorization)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}
		customClaims := new(CustomClaims)
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
			fmt.Println(diff.Seconds(), refreshTTL)
			if diff < refreshTTL {
				exp := time.Now().Add(time.Duration(core.Cfg.JWT.Expires) * time.Minute)
				customClaims.ExpiresAt(exp)
				newToken, _ := Refresh(customClaims, core.Cfg.JWT.SignKey)
				c.Writer.Header().Set("refresh-access-token", newToken)
				c.Writer.Header().Set("refresh-exp", strconv.FormatInt(exp.Unix(), 10))
			}
		}

		c.Set("a_uid", customClaims.UserId)
		c.Set("a_mobile", customClaims.Phone)
		c.Set("a_nickname", customClaims.Nickname)
		c.Set("jwt_data", customClaims.JwtData)
		c.Next()
	}
}

func GetUserId(c *gin.Context) int {
	uid := c.GetInt("a_uid")
	if uid == 0 {
		suid := c.GetHeader("a_uid")
		if suid != "" {
			uid, _ = strconv.Atoi(suid)
		}
	}
	return uid
}

func GetPhone(c *gin.Context) string {
	phone := c.GetString("a_mobile")
	if phone == "" {
		phone = c.GetHeader("phone")
	}
	return phone
}

func GetNickname(c *gin.Context) string {
	nickname := c.GetString("a_nickname")
	if nickname == "" {
		nickname = c.GetHeader("a_nickname")
	}
	return nickname
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// Generate 生成JWT Token
func Generate(claims jwt.Claims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名字符串
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// Refresh 刷新JWT Token
func Refresh(claims jwt.Claims, secretKey string) (string, error) {
	return Generate(claims, secretKey)
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

// AdminCustomClaims 自定义格式内容
type CustomClaims struct {
	UserId               int    `json:"uid,omitempty"`
	RoleId               int    `json:"rid,omitempty"`
	Phone                string `json:"mob,omitempty"`
	Nickname             string `json:"nick,omitempty"`
	JwtData              map[string]any
	jwt.RegisteredClaims // 内嵌标准的声明
}

func (c *CustomClaims) AddData(key string, val any) *CustomClaims {
	if c.JwtData == nil {
		c.JwtData = make(map[string]any, 0)
	}
	c.JwtData[key] = val
	return c
}

func (c *CustomClaims) GetInt(key string) int {
	if val, ok := c.JwtData[key]; ok {
		return utils.GetInterfaceToInt(val)
	}
	return 0
}

func (c *CustomClaims) GetString(key string) string {
	if val, ok := c.JwtData[key]; ok {
		return fmt.Sprintf("%s", val)
	}
	return ""
}

func (c *CustomClaims) ExpiresAt(expiresAt time.Time) *CustomClaims {
	c.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	return c
}

// NewAdminCustomClaims 初始化AdminCustomClaims
func NewClaims(userId int, expiresAt time.Time, issuer, subject string) CustomClaims {
	//now := time.Now()
	return CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt), // 定义过期时间
			Issuer:    issuer,                        // 签发人
			//IssuedAt:  jwt.NewNumericDate(now),       // 签发时间
			Subject: subject, // 签发主体
			//NotBefore: jwt.NewNumericDate(now),       // 生效时间
		},
	}
}
