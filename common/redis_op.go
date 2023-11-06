package common

import (
	"fmt"

	"github.com/baowk/dilu-core/core"
)

func GetMpAccessToken(appId string) string {
	j, err := core.Cache.Get("acctoken:" + appId)
	if err == nil {
		return j
	}
	return ""
}

func SetMpAccessToken(appId string, token string) {
	core.Cache.Set("acctoken:"+appId, token, 7000)
}

func GetMpOpenId(scene string) (string, error) {
	return core.Cache.Get("mp:login:" + scene)
}

func SetMpOpenId(scene, openId string) {
	core.Cache.Set("mp:login:"+scene, openId, 400)
}

func DelMpOpenId(scene string) error {
	err := core.Cache.Del("mp:login:" + scene)
	if err != nil {
		return err
	}
	return nil
}

func TeamMemberKey(teamId, userId int) string {
	return fmt.Sprintf("t:m:%d:%d", teamId, userId)
}
