package middleware

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func PermHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetInt("a_rid")
		if rid == 1 {
			c.Next()
			return
		}
		curUri := c.Request.URL.Path
		curMethod := c.Request.Method
		uid := c.GetInt("a_uid")

		fmt.Println(curUri, curMethod, uid, rid)

		c.Next()
	}
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
