package apis

import (
	"dilu/common/codes"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"
	"errors"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/common/utils/regexp_util"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SSO struct {
	base.BaseApi
}

// 发送验证码
// SendCode 发送邮箱或者短信验证码
// @Summary 发送邮箱或者短信验证码
// @Description 发送邮箱或者短信验证码
// @Tags sso
// @Accept application/json
// @Product application/json
// @Param data body dto.CodeSendReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sso/sendCode [post]
func (e *SSO) SendCode(c *gin.Context) {
	req := dto.CodeSendReq{}
	if err := c.ShouldBind(&req); err != nil {
		return
	}

	if req.Code == "" || req.UUID == "" {
		e.Err(c, codes.ErrInvalidParameter(e.GetReqId(c), nil))
		return
	}

	if !service.Verify(req.UUID, req.Code, true) {
		e.Err(c, errs.Err(codes.CaptchaVerifyErr, e.GetReqId(c), nil))
		return
	}

	//是否手机
	if regexp_util.CheckMobile(req.Username) {
		if req.CheckExist {
			var count int64
			service.SysUserS.CountByPhone(req.Username, &count)
			if count > 0 {
				e.Code(c, codes.PhoneExistErr)
				return
			}
		}
		service.SmsS.Send(req.Username)

	} else if regexp_util.CheckEmail(req.Username) { //是否邮箱
		if req.CheckExist {
			var count int64
			service.SysUserS.CountByEmail(req.Username, &count)
			if count > 0 {
				e.Code(c, codes.EmailExistErr)
				return
			}
		}
		service.EmailS.Send(req.Username)
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}
	e.Ok(c, c, "验证码发送成功")
}

// 用户注册
// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags sso
// @Accept application/json
// @Product application/json
// @Param data body dto.RegisterReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/v1/sso/register [post]
func (e *SSO) Register(c *gin.Context) {
	req := dto.RegisterReq{}

	if req.Password != req.RePassword {
		e.Code(c, codes.ErrRePassword)
		return
	}
	//密码规则
	if req.Password != "" && !regexp_util.CheckPwd(req.Password) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}

	loginType := 1
	//是否手机
	if regexp_util.CheckMobile(req.Username) {
		if !service.SmsS.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		loginType = 1
	} else if regexp_util.CheckEmail(req.Username) { //是否邮箱
		if !service.EmailS.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		loginType = 2
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}

	ip := utils.GetIP(c)
	if logOk, err := service.SysUserS.Register(loginType, &req, ip); err != nil {
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
// @Tags sso
// @Accept application/json
// @Product application/json
// @Param data body dto.VerifyCodeReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sso/verify/code [post]
func (e *SSO) VerifyCode(c *gin.Context) {
	req := dto.VerifyCodeReq{}
	c.ShouldBind(&req)

	//是否手机
	if regexp_util.CheckMobile(req.Username) {
		s := service.SmsLog{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else if regexp_util.CheckEmail(req.Username) { //是否邮箱
		s := service.EmailLog{}

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
// @Tags sso
// @Accept application/json
// @Product application/json
// @Param data body dto.LoginReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/v1/sso/login [post]
func (e *SSO) Login(c *gin.Context) {
	req := dto.LoginReq{}
	c.ShouldBind(&req)
	ip := utils.GetIP(c)
	if req.Password == "" {
		//是否手机
		if regexp_util.CheckMobile(req.Username) {
			if !service.SmsS.Verify(req.Username, req.Code) {
				e.Code(c, codes.ErrVerifyCode)
				return
			}
		} else if regexp_util.CheckEmail(req.Username) { //是否邮箱
			if !service.EmailS.Verify(req.Username, req.Code) {
				e.Code(c, codes.ErrVerifyCode)
				return
			}
		} else {
			e.Code(c, codes.ErrMobileOrEmail)
			return
		}
	} else {
		if logOk, err := service.SysUserS.LoginPwd(&req, ip); err != nil {
			core.Log.Error("sso", zap.Error(err))
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					e.Code(c, codes.ErrUserNotExist)
					return
				} else {
					e.Code(c, 500)
					return
				}
			}
			return
		} else {
			e.Ok(c, logOk, "登录成功")
			return
		}
	}
	if logOk, err := service.SysUserS.LoginCode(&req, ip); err != nil {
		core.Log.Error("sso", zap.Error(err))
		e.Code(c, 500)
		return
	} else {
		e.Ok(c, logOk, "登录成功")
	}
}

// 忘记密码
// ForgetPwd 忘记密码
// @Summary 忘记密码
// @Description 忘记密码
// @Tags sso
// @Accept application/json
// @Product application/json
// @Param data body dto.ForgetPwdReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sso/forgetPwd [post]
func (e *SSO) ForgetPwd(c *gin.Context) {
	req := dto.ForgetPwdReq{}
	c.ShouldBind(&req)

	if !regexp_util.CheckPwd(req.Password) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}

	var mobile, email string
	//是否手机
	if regexp_util.CheckMobile(req.Username) {
		s := service.SmsLog{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		mobile = req.Username
	} else if regexp_util.CheckEmail(req.Username) { //是否邮箱
		s := service.EmailLog{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		email = req.Username
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}
	if err := service.SysUserS.ChangePwd(mobile, email, req.Password); err != nil {
		core.Log.Error("sso", zap.Error(err))
		e.Code(c, 500)
		return
	} else {
		e.Ok(c)
	}

}

// // 获取用户信息
// // GetUserInfo 获取用户信息
// // @Summary 获取用户信息
// // @Description 获取用户信息
// // @Tags sso
// // @Param data body dto.IdReq true "data"
// // @Success 200 {object} base.Resp{data=models.User} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/getUserinfo [post]
// func (e *SSO) GetUserInfo(c *gin.Context) {
// 	req := dto.IdReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		MakeOrm().
// 		Bind(&req).
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}
// 	var object models.User

// 	err = s.GetByUserId(req.Id, &object)
// 	if err != nil {
// 		e.Code(500, err, fmt.Sprintf("获取用户表失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	resp := dto.UserinfoResp{}
// 	if err := copier.Copy(&resp, object); err != nil {
// 		e.Code(500, err, fmt.Sprintf("拷贝失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}

// 	e.Ok(c,resp, "查询成功")
// }

// // 获取个人信息
// // GetUserInfo 获取个人信息
// // @Summary 获取个人信息
// // @Description 获取个人信息
// // @Tags sso
// // @Success 200 {object} base.Resp{data=models.User} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/auth/myUserinfo [post]
// // @Security Bearer
// func (e *SSO) MyUserInfo(c *gin.Context) {
// 	s := service.User{}
// 	err := e.MakeContext(c).
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
// 	var object models.User
// 	err = s.GetByUserId(e.GetUserId(), &object)
// 	if err != nil {
// 		e.Code(500, err, fmt.Sprintf("获取用户表失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	resp := dto.MyUserinfoResp{}
// 	if err := copier.Copy(&resp, object); err != nil {
// 		e.Code(500, err, fmt.Sprintf("拷贝失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	e.Ok(c,resp, "查询成功")

// }

// // 修改密码
// // ChangePwd 修改密码
// // @Summary 修改密码
// // @Description 修改密码
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.ChangePwdReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/auth/changePwd [post]
// // @Security Bearer
// func (e *SSO) ChangePwd(c *gin.Context) {
// 	req := dto.ChangePwdReq{}
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

// 	if !regexp_util.CheckPwd(req.NewPassword) {
// 		e.Code(c, codes.ErrPasswordFMT)
// 		return
// 	}

// 	if e.GetUserId() == "" {
// 		e.Code(c, codes.ErrUnLogin)
// 		return
// 	}
// 	if err := s.ChangePwdByOld(e.GetUserId(), req.OldPassword, req.NewPassword, req.InviteCode); err != nil {
// 		e.Code(500, err, fmt.Sprintf("修改密码失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	e.Ok(c,c, "修改成功")
// }

// // 绑定手机号或者邮箱
// // Bind 绑定手机号或者邮箱
// // @Summary 绑定手机号或者邮箱
// // @Description 绑定手机号或者邮箱
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.BindReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/auth/bind [post]
// // @Security Bearer
// func (e *SSO) Bind(c *gin.Context) {
// 	req := dto.BindReq{}
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
// 	//是否手机
// 	if regexp_util.CheckMobile(req.Username) {
// 		s := service.SmsLog{}
//
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Code(c, codes.ErrVerifyCode)
// 			return
// 		}
// 	} else if regexp_util.CheckEmail(req.Username) { //是否邮箱
// 		s := service.EmailLog{}
//
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Code(c, codes.ErrVerifyCode)
// 			return
// 		}
// 	} else {
// 		e.Code(c, codes.ErrMobileOrEmail)
// 		return
// 	}

// 	if err := s.Bind(e.GetUserId(), &req); err != nil {
// 		e.Code(c, codes.ErrBind)
// 		return
// 	}

// 	e.Ok(c,c, "绑定成功")
// }

// // 修改用户信息
// // ChangeUserinfo 修改用户信息
// // @Summary 修改用户信息
// // @Description 修改用户信息
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.ChangeUserinfoReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/auth/changeUserinfo [post]
// // @Security Bearer
// func (e *SSO) ChangeUserinfo(c *gin.Context) {
// 	req := dto.ChangeUserinfoReq{}
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
// 	//是否手机

// 	if err := s.ChangeUserinfo(e.GetUserId(), &req); err != nil {
// 		e.Code(c, codes.ErrBind)
// 		return
// 	}
// 	e.Ok(c,c, "修改成功")
// }

// // // 退出登录
// // // Logout 退出登录
// // // @Summary 退出登录
// // // @Description 退出登录（调用后清空本地token）
// // // @Tags sso
// // // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // // @Router /api/v1/sso/auth/logout [post]
// // // @Security Bearer
// // func (e *SSO) Logout(c *gin.Context) {

// // 	e.Ok(c,c, "退出成功")
// // }

// // 微信登录
// // LoginByWechat 微信登录
// // @Summary 微信登录
// // @Description 微信登录
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/loginWechat [post]
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
// 		e.Ok(c,logOk, "登录成功")
// 	}
// }

// // 获取钉钉登录配置信息
// // GeDingCfg 获取钉钉登录配置信息
// // @Summary 获取钉钉登录配置信息
// // @Description 获取钉钉登录配置信息
// // @Tags sso
// // @Success 200 {object} base.Resp{data=dto.DingCfgResp} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/getDingCfg [post]
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
// 	e.Ok(c,cfg, "OK")
// }

// // 钉钉
// // LoginByDing 钉钉登录
// // @Summary 钉钉
// // @Description 钉钉登录
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginDingReq true "data"
// // @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/loginDing [post]
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
// 		e.Ok(c,logOk, "登录成功")
// 	}

// }

// // 绑定钉钉
// // BindDing 绑定钉钉
// // @Summary 绑定钉钉
// // @Description 绑定钉钉
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginDingReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/auth/bindDing [post]
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
// 	e.Ok(c,c, "成功")
// }

// // 绑定微信
// // BindWechat 绑定微信
// // @Summary 绑定微信
// // @Description 绑定微信
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sso/auth/bindWechat [post]
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
// 	e.Ok(c,c, "成功")
// }
