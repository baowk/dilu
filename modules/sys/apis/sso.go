package apis

import (
	"dilu/common/codes"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"
	"errors"

	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/common/utils/regexps"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ApiSso = SSO{}

type SSO struct {
	base.BaseApi
}

// 发送验证码
// SendCode 发送邮箱或者短信验证码
// @Summary 发送邮箱或者短信验证码
// @Description 发送邮箱或者短信验证码
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.CodeSendReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sendCode [post]
func (e *SSO) SendCode(c *gin.Context) {
	req := dto.CodeSendReq{}
	if err := c.ShouldBind(&req); err != nil {
		return
	}

	if req.Code == "" || req.UUID == "" {
		e.Err(c, codes.ErrInvalidParameter(e.GetReqId(c), "code or uuid is all nil"))
		return
	}

	if !service.Verify(req.UUID, req.Code, true) {
		e.Err(c, errs.Err(codes.CaptchaVerifyErr, e.GetReqId(c), nil))
		return
	}

	//是否手机
	if regexps.CheckMobile(req.Username) {
		if req.CheckExist {
			var count int64
			service.SerSysUser.CountByPhone(req.Username, &count)
			if count > 0 {
				e.Code(c, codes.PhoneExistErr)
				return
			}
		}
		var tempId string
		service.SerSms.Send(req.Username, tempId)

	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		if req.CheckExist {
			var count int64
			service.SerSysUser.CountByEmail(req.Username, &count)
			if count > 0 {
				e.Code(c, codes.EmailExistErr)
				return
			}
		}
		service.SerEmail.Send(req.Username)
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}
	e.Ok(c)
}

// 用户注册
// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.RegisterReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/register [post]
func (e *SSO) Register(c *gin.Context) {
	req := dto.RegisterReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.Password != req.RePassword {
		e.Code(c, codes.ErrRePassword)
		return
	}
	//密码规则
	if req.Password != "" && !regexps.CheckPwd(req.Password) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}

	loginType := 1
	//是否手机
	if regexps.CheckMobile(req.Username) {
		if !service.SerSms.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		loginType = 1
	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		if !service.SerEmail.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		loginType = 2
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}

	ip := ips.GetIP(c)
	if logOk, err := service.SerSysUser.Register(loginType, &req, ip); err != nil {
		core.Log.Error("sso", zap.Error(err))
		e.Error(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// 验证码校验
// VerifyCode 验证码校验
// @Summary 验证码校验
// @Description 验证码校验
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.VerifyCodeReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/verify/code [post]
func (e *SSO) VerifyCode(c *gin.Context) {
	req := dto.VerifyCodeReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	//是否手机
	if regexps.CheckMobile(req.Username) {
		s := service.SysSms{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		s := service.SysEmail{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}
	e.Ok(c)
}

// 用户登录
// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.LoginReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/login [post]
func (e *SSO) Login(c *gin.Context) {
	req := dto.LoginReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	ip := ips.GetIP(c)
	if req.Password == "" {
		//是否手机
		if regexps.CheckMobile(req.Username) {
			if !service.SerSms.Verify(req.Username, req.Code) {
				e.Code(c, codes.ErrVerifyCode)
				return
			}
		} else if regexps.CheckEmail(req.Username) { //是否邮箱
			if !service.SerEmail.Verify(req.Username, req.Code) {
				e.Code(c, codes.ErrVerifyCode)
				return
			}
			// } else {
			// 	e.Code(c, codes.ErrMobileOrEmail)
			// 	return
		}
	} else {
		if logOk, err := service.SerSysUser.LoginPwd(&req, ip); err != nil {
			core.Log.Error("sso", zap.Error(err))
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					e.Code(c, codes.UserNotExist)
					return
				} else {
					e.Err(c, err)
					return
				}
			}
			return
		} else {
			e.Ok(c, logOk)
			return
		}
	}
	if logOk, err := service.SerSysUser.LoginCode(&req, ip); err != nil {
		core.Log.Error("sso", zap.Error(err))
		e.Err(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// 忘记密码
// ForgetPwd 忘记密码
// @Summary 忘记密码
// @Description 忘记密码
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.ForgetPwdReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/forgetPwd [post]
func (e *SSO) ForgetPwd(c *gin.Context) {
	req := dto.ForgetPwdReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if !regexps.CheckPwd(req.Password) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}

	var mobile, email string
	//是否手机
	if regexps.CheckMobile(req.Username) {
		s := service.SysSms{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		mobile = req.Username
	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		s := service.SysEmail{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		email = req.Username
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}
	if err := service.SerSysUser.ChangePwd(mobile, email, req.Password); err != nil {
		core.Log.Error("sso", zap.Error(err))
		e.Error(c, err)
		return
	} else {
		e.Ok(c)
	}

}

// 获取用户信息
// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags sys-sso
// @Param data body base.ReqId true "data"
// @Success 200 {object} base.Resp{data=dto.UserinfoResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/getUserinfo [post]
func (e *SSO) GetUserInfo(c *gin.Context) {
	req := base.ReqId{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var object models.SysUser

	err := service.SerSysUser.Get(req.Id, &object)
	if err != nil {
		core.Log.Error("sso", zap.Error(err))
		e.Error(c, err)
		return
	}
	resp := dto.UserinfoResp{}
	if err := copier.Copy(&resp, object); err != nil {
		core.Log.Error("sso", zap.Error(err))
		e.Error(c, err)
		return
	}

	e.Ok(c, resp)
}

// 获取个人信息
// GetUserInfo 获取个人信息
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags sys-sso
// @Success 200 {object} base.Resp{data=dto.MyUserinfoResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/myUserinfo [get]
// @Security Bearer
func (e *SSO) MyUserInfo(c *gin.Context) {
	uid := utils.GetUserId(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var object models.SysUser
	err := service.SerSysUser.Get(uid, &object)
	if err != nil {
		e.Error(c, err)
		return
	}
	resp := dto.MyUserinfoResp{}
	if err := copier.Copy(&resp, object); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, resp)

}

// 修改密码
// ChangePwd 修改密码
// @Summary 修改密码
// @Description 修改密码
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.ChangePwdReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/changePwd [post]
// @Security Bearer
func (e *SSO) ChangePwd(c *gin.Context) {
	req := dto.ChangePwdReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.NewPassword != req.RePassword {
		e.Code(c, codes.ErrRePassword)
		return
	}
	if !regexps.CheckPwd(req.NewPassword) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}
	uid := utils.GetUserId(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}

	if err := service.SerSysUser.ChangePwdByOld(uid, req.OldPassword, req.NewPassword, req.InviteCode); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c)
}

// 绑定手机号或者邮箱
// Bind 绑定手机号或者邮箱
// @Summary 绑定手机号或者邮箱
// @Description 绑定手机号或者邮箱
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.BindReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/auth/bind [post]
// @Security Bearer
func (e *SSO) Bind(c *gin.Context) {
	req := dto.BindReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	uid := utils.GetUserId(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	//是否手机
	if regexps.CheckMobile(req.Username) {
		s := service.SysSms{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		s := service.SysEmail{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}

	if err := service.SerSysUser.Bind(uid, &req); err != nil {
		e.Code(c, codes.ErrBind)
		return
	}

	e.Ok(c)
}

// 修改用户信息
// ChangeUserinfo 修改用户信息
// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.ChangeUserinfoReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/auth/changeUserinfo [post]
// @Security Bearer
func (e *SSO) ChangeUserinfo(c *gin.Context) {
	req := dto.ChangeUserinfoReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	uid := utils.GetUserId(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}

	var user models.SysUser
	copier.Copy(&user, req)
	user.Id = uid
	//是否手机
	if err := service.SerSysUser.ChangeUserinfo(uid, user); err != nil {
		e.Code(c, codes.ErrBind)
		return
	}
	e.Ok(c)
}

// 退出登录
// Logout 退出登录
// @Summary 退出登录
// @Description 退出登录（调用后清空本地token）
// @Tags sys-sso
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/auth/logout [post]
// @Security Bearer
func (e *SSO) Logout(c *gin.Context) {

	e.Ok(c)
}

// // 微信登录
// // LoginByWechat 微信登录
// // @Summary 微信登录
// // @Description 微信登录
// // @Tags sys-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/loginWechat [post]
// func (e *SSO) LoginByWechat(c *gin.Context) {
// 	req := dto.LoginWechatReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)

// 	if logOk, err := s.LoginWechat(&req, ip); err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.Ok(c,logOk)
// 	}
// }

// // 获取钉钉登录配置信息
// // GeDingCfg 获取钉钉登录配置信息
// // @Summary 获取钉钉登录配置信息
// // @Description 获取钉钉登录配置信息
// // @Tags sys-sso
// // @Success 200 {object} base.Resp{data=dto.DingCfgResp} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/getDingCfg [post]
// func (e *SSO) GetDingCfg(c *gin.Context) {
// 	err := e.MakeContext(c).MakeOrm().
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}
// 	domain := "http://" + c.Request.Host
// 	var cfg dto.DingCfgResp
// 	service.GetDingConfig(domain, &cfg)
// 	e.Ok(c,cfg)
// }

// // 钉钉
// // LoginByDing 钉钉登录
// // @Summary 钉钉
// // @Description 钉钉登录
// // @Tags sys-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginDingReq true "data"
// // @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/loginDing [post]
// func (e *SSO) LoginByDing(c *gin.Context) {
// 	req := dto.LoginDingReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)

// 	if logOk, err := s.LoginDing(&req, ip); err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.Ok(c,logOk)
// 	}

// }

// // 绑定钉钉
// // BindDing 绑定钉钉
// // @Summary 绑定钉钉
// // @Description 绑定钉钉
// // @Tags sys-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginDingReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/auth/bindDing [post]
// // @Security Bearer
// func (e *SSO) BindDing(c *gin.Context) {
// 	req := dto.LoginDingReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Code(c, codes.ErrUnLogin)
// 		return
// 	}

// 	if err := s.BindDing(&req, e.GetUserId()); err != nil {
// 		e.Code(c, codes.ErrBind)
// 		return
// 	}
// 	e.Ok(c)
// }

// // 绑定微信
// // BindWechat 绑定微信
// // @Summary 绑定微信
// // @Description 绑定微信
// // @Tags sys-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/auth/bindWechat [post]
// // @Security Bearer
// func (e *SSO) BindWechat(c *gin.Context) {
// 	req := dto.LoginWechatReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Code(c, codes.ErrUnLogin)
// 		return
// 	}

// 	if err := s.BindWechat(&req, e.GetUserId()); err != nil {
// 		e.Code(c, codes.ErrBind)
// 		return
// 	}
// 	e.Ok(c)
// }
