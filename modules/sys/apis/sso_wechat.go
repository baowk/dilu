package apis

import (
	"dilu/common/codes"
	"dilu/common/config"
	"dilu/common/third/wechat"

	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"net/http"
	"strings"

	"dilu/common"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type Wechat struct {
	base.BaseApi
}

// MPCallback 微信公众号回调函数
// @Summary 微信公众号回调函数
// @Description 微信公众号回调函数
// @Tags sys-sso wechat mp
// @Accept application/json
// @Product application/json
// @Success 200 {object} string
// @Router /api/v1/sys/mp/callback [get]
func (e *Wechat) MPCallback(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	if wechat.MPCheckSign(timestamp, nonce, signature, config.Ext.WechatMp.WxToken) {
		if echostr != "" {
			c.String(http.StatusOK, echostr)
		} else {
			var req wechat.MPMsg
			err := c.ShouldBind(&req)
			if err != nil {
				e.Error(c, err)
				return
			}
			if req.MsgType == "event" {
				var qrscene string
				if req.EventKey != "" {
					if strings.HasPrefix(req.EventKey, "qrscene_") {
						qrscene = req.EventKey[8:]
					} else {
						qrscene = req.EventKey
					}
				}

				if qrscene != "" {
					common.SetMpOpenId(qrscene, req.FromUserName)
				}
			}
			return
		}
	} else {
		c.String(http.StatusOK, "fail")
	}
}

// GetMpQrcode 获取公众号二维码
// @Summary 获取公众号二维码
// @Description 获取公众号二维码
// @Tags sys-sso wechat mp
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=wechat.QrCodeResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/mp/qrCode [post]
func (e *Wechat) GetMpQrcode(c *gin.Context) {
	accT, err := getAccessToken(config.Ext.WechatMp.AppId, config.Ext.WechatMp.AppSecret)
	if err != nil {
		e.Error(c, err)
		return
	}
	rstr := utils.GenUUid()
	qr, err := wechat.GetQrCode(accT, 300, 0, rstr)
	if err != nil {
		e.Error(c, err)
		return
	}
	qr.Scene = rstr
	e.Ok(c, qr)
}

// LoginMp  轮询用户登录成功与否
// @Summary  轮询用户登录成功与否
// @Description 获取公众号二维码
// @Tags sys-sso wechat mp
// @Accept application/json
// @Product application/json
// @Param data body dto.MpSceneReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/mp/login [post]
func (e *Wechat) LoginMp(c *gin.Context) {
	var req dto.MpSceneReq

	str, err := common.GetMpOpenId(req.Scene)
	if err != nil {
		e.Code(c, codes.ThirdNotScan)
		return
	}
	if str == "" {
		e.Code(c, codes.ThirdExpire)
		return
	}

	ip := ips.GetIP(c)
	lok, err := service.SerSysUser.LoginWechatMp(req, str, ip)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, lok)
}

func getAccessToken(appId, appSecret string) (string, error) {
	at := common.GetMpAccessToken(appId)
	if at != "" {
		return at, nil
	}
	data, err := wechat.GetAccessToken(appId, appSecret)
	if err != nil {
		return "", err
	}
	common.SetMpAccessToken(appId, data.AccessToken)
	return data.AccessToken, nil
}
