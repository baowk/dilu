package middleware

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func PermHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetInt("a_rid")
		if rid == -1 {
			c.Next()
			return
		}

		curUri := c.Request.URL.Path
		curMethod := c.Request.Method
		//uid := c.GetInt("a_uid")

		//fmt.Println(curUri, curMethod, uid, rid)
		var aid int
		apis := GetApis()
		for _, api := range apis {
			if api.Method == curMethod && KeyMatch2(curUri, api.Path) {
				aid = api.Id
				break
			}
		}
		//fmt.Println(aid)
		if aid < 1 {
			Fail(c, 403, "无权限")
			return
		}

		if err := service.SerSysMenu.CanAccess(c, aid); err != nil {
			Fail(c, 403, "无权限")
			return
		}

		c.Next()
	}
}

var apis []models.SysApi

func GetApis() []models.SysApi {
	if len(apis) == 0 {
		service.SerSysApi.GetByType(3, &apis)
	}
	return apis
}

// KeyMatch2 determines whether key1 matches the pattern of key2 (similar to RESTful path), key2 can contain a *.
// For example, "/foo/bar" matches "/foo/*", "/resource1" matches "/:resource"
func KeyMatch2(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}

// RegexMatch determines whether key1 matches the pattern of key2 in regular expression.
func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}
