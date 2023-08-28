package apis

import (
	"dilu/common/third/wechat"

	"net/http"
	"strings"

	"dilu/common"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type Wechat struct {
	base.BaseApi
}

var (
	appId          = ""
	appSecret      = ""
	wxToken        = ""
	encodingAESKey = ""
)

// MPCallback 微信公众号回调函数
// @Summary 微信公众号回调函数
// @Description 微信公众号回调函数
// @Tags 微信公众号
// @Accept application/json
// @Product application/json
// @Success 200 {object} string
// @Router /api/v1/sso/mp/callback [get]
func (e *Wechat) MPCallback(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	if wechat.MPCheckSign(timestamp, nonce, signature, wxToken) {
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
// @Tags 微信公众号
// @Accept application/json
// @Product application/json
// @Success 200 {object} response.Response{data=wechat.QrCodeResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sso/mp/qrCode [post]
func (e *Wechat) GetMpQrcode(c *gin.Context) {
	accT, err := getAccessToken(appId, appSecret)
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
// @Tags 微信公众号
// @Accept application/json
// @Product application/json
// @Param data body dto.MpSceneReq true "data"
// @Success 200 {object} response.Response{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/v1/sso/mp/login [post]
// func (e *Wechat) LoginMp(c *gin.Context) {
// 	var req dto.MpSceneReq
// 	s := service.User{}

// 	str, err := common.GetMpOpenId(req.Scene)
// 	if err != nil {
// 		e.ErrorN(constn.ErrMpNotScan)
// 		return
// 	}
// 	if str == "" {
// 		e.ErrorN(constn.ErrMpExpire)
// 		return
// 	}

// 	ip := utils.GetIP(c)
// 	lok, err := s.LoginWechatMp(req, str, ip)
// 	if err != nil {
// 		e.Err(c, err)
// 		return
// 	}
// 	e.Ok(c, lok)
// }

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
