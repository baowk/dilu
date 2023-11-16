package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

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

func GetTeamId(c *gin.Context) int {
	if GetRoleId(c) != 0 {
		return -1
	}
	sTeamId := c.GetHeader("teamId")
	if sTeamId != "" {
		teamId, _ := strconv.Atoi(sTeamId)
		return teamId
	}
	return c.GetInt("teamId")
}

func GetReqTeamId(c *gin.Context, reqTeamId int) int {
	teamId := GetTeamId(c)
	if teamId == -1 {
		if reqTeamId == 0 {
			return teamId
		}
		return reqTeamId
	}
	return teamId
}

func GetRoleId(c *gin.Context) int {
	rid := c.GetInt("a_rid")
	if rid == 0 {
		suid := c.GetHeader("a_rid")
		if suid != "" {
			rid, _ = strconv.Atoi(suid)
		}
	}
	return rid
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
