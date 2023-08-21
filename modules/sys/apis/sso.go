package apis

// import (
// 	"dilu/modules/sys/models"
// 	"dilu/modules/sys/service/dto"
// 	"errors"
// 	"fmt"

// 	"github.com/baowk/dilu-core/core/base"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"
// 	"github.com/jinzhu/copier"
// 	"gorm.io/gorm"
// )

// type SSO struct {
// 	base.BaseApi
// }

// // GenerateCaptchaHandler 获取验证码
// // @Summary 获取验证码
// // @Description 获取验证码
// // @Tags sso
// // @Success 200 {object} response.Response{data=string,id=string,msg=string} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/captcha [post]
// func (e SSO) GenerateCaptchaHandler(c *gin.Context) {
// 	err := e.MakeContext(c).Errors
// 	if err != nil {
// 		e.Error(500, err, "服务初始化失败！")
// 		return
// 	}
// 	id, b64s, err := captcha.DriverDigitFunc()
// 	if err != nil {
// 		e.Logger.Errorf("DriverDigitFunc error, %s", err.Error())
// 		e.Error(500, err, "验证码获取失败")
// 		return
// 	}
// 	e.Custom(gin.H{
// 		"code": 200,
// 		"data": b64s,
// 		"id":   id,
// 		"msg":  "success",
// 	})
// }

// // 发送验证码
// // SendCode 发送邮箱或者短信验证码
// // @Summary 发送邮箱或者短信验证码
// // @Description 发送邮箱或者短信验证码
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.CodeSendReq true "data"
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/sendCode [post]
// func (e SSO) SendCode(c *gin.Context) {
// 	req := dto.CodeSendReq{}
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}

// 	// if req.Code == "" || req.UUID == "" {
// 	// 	e.Error(c,constn.ParamsErr)
// 	// 	return
// 	// }

// 	// if !captcha.Verify(req.UUID, req.Code, true) {
// 	// 	e.Error(c,constn.ErrVerifyCode)
// 	// 	return
// 	// }

// 	//是否手机
// 	if regexputil.CheckMobile(req.Username) {
// 		if req.CheckExist {
// 			s := service.User{}
// 			s.Orm = e.Orm
// 			var count int64
// 			s.CountByMobile(req.Username, &count)
// 			if count > 0 {
// 				e.Error(c, constn.ErrUserExist)
// 				return
// 			}
// 		}
// 		s := service.SmsLog{}
// 		if err := e.MakeService(&s.Service).Errors; err != nil {
// 			e.Logger.Error(err)
// 			e.Error(500, err, err.Error())
// 			return
// 		}
// 		s.Send(req.Username)

// 	} else if regexputil.CheckEmail(req.Username) { //是否邮箱
// 		if req.CheckExist {
// 			s := service.User{}
// 			s.Orm = e.Orm
// 			var count int64
// 			s.CountByEmail(req.Username, &count)
// 			if count > 0 {
// 				e.Error(c, constn.ErrUserExist)
// 				return
// 			}
// 		}
// 		s := service.EmailLog{}
// 		if err := e.MakeService(&s.Service).Errors; err != nil {
// 			e.Logger.Error(err)
// 			e.Error(500, err, err.Error())
// 			return
// 		}
// 		s.Send(req.Username)
// 	} else {
// 		e.Error(c, constn.ErrMobileOrEmail)
// 		return
// 	}

// 	e.OK("", "验证码发送成功")
// }

// // 用户注册
// // Register 用户注册
// // @Summary 用户注册
// // @Description 用户注册
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.RegisterReq true "data"
// // @Success 200 {object} response.Response{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/register [post]
// func (e SSO) Register(c *gin.Context) {
// 	req := dto.RegisterReq{}
// 	e.MakeContext(c).Bind(&req)

// 	if req.Password != req.RePassword {
// 		e.Error(c, constn.ErrRePassword)
// 		return
// 	}

// 	//密码规则
// 	if req.Password != "" && !regexputil.CheckPwd(req.Password) {
// 		e.Error(c, constn.ErrPasswordFMT)
// 		return
// 	}

// 	loginType := 1
// 	//是否手机
// 	if regexputil.CheckMobile(req.Username) {
// 		s := service.SmsLog{}
// 		e.MakeOrm().MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 		loginType = 1
// 	} else if regexputil.CheckEmail(req.Username) { //是否邮箱
// 		s := service.EmailLog{}
// 		e.MakeOrm().MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 		loginType = 2
// 	} else {
// 		e.Error(c, constn.ErrMobileOrEmail)
// 		return
// 	}

// 	s := service.User{}
// 	err := e.
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)
// 	if logOk, err := s.Register(loginType, &req, ip); err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, fmt.Sprintf("用户注册失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.OK(logOk, "注册成功")
// 	}
// }

// // 验证码校验
// // VerifyCode 验证码校验
// // @Summary 验证码校验
// // @Description 验证码校验
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.VerifyCodeReq true "data"
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/verify/code [post]
// func (e SSO) VerifyCode(c *gin.Context) {
// 	req := dto.VerifyCodeReq{}
// 	e.MakeContext(c).Bind(&req)

// 	//是否手机
// 	if regexputil.CheckMobile(req.Username) {
// 		s := service.SmsLog{}
// 		e.MakeOrm().MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 	} else if regexputil.CheckEmail(req.Username) { //是否邮箱
// 		s := service.EmailLog{}
// 		e.MakeOrm().MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 	} else {
// 		e.Error(c, constn.ErrMobileOrEmail)
// 		return
// 	}
// 	e.OK("OK", "OK")

// }

// // 用户登录
// // Login 用户登录
// // @Summary 用户登录
// // @Description 用户登录
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginReq true "data"
// // @Success 200 {object} response.Response{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/login [post]
// func (e SSO) Login(c *gin.Context) {
// 	req := dto.LoginReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)

// 	if req.Password == "" {
// 		//是否手机
// 		if regexputil.CheckMobile(req.Username) {
// 			s := service.SmsLog{}
// 			e.MakeService(&s.Service)
// 			if !s.Verify(req.Username, req.Code) {
// 				e.Error(c, constn.ErrVerifyCode)
// 				return
// 			}
// 		} else if regexputil.CheckEmail(req.Username) { //是否邮箱
// 			s := service.EmailLog{}
// 			e.MakeService(&s.Service)
// 			if !s.Verify(req.Username, req.Code) {
// 				e.Error(c, constn.ErrVerifyCode)
// 				return
// 			}
// 		} else {
// 			e.Error(c, constn.ErrMobileOrEmail)
// 			return
// 		}
// 	} else {
// 		if logOk, err := s.LoginPwd(&req, ip); err != nil {
// 			e.Logger.Error(err)
// 			if err != nil {
// 				if errors.Is(err, gorm.ErrRecordNotFound) {
// 					e.Error(c, constn.ErrUserNotExist)
// 					return
// 				} else if errors.Is(err, service.ErrPwdNotExist) {
// 					e.Error(c, constn.ErrPwdNotExist)
// 					return
// 				} else {
// 					e.Error(500, err, err.Error())
// 					return
// 				}
// 			}
// 			return
// 		} else {
// 			e.OK(logOk, "登录成功")
// 			return
// 		}
// 	}
// 	if logOk, err := s.LoginCode(&req, ip); err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.OK(logOk, "登录成功")
// 	}
// }

// // Union用户注册
// // UnionReg Union用户注册
// // @Summary Union用户注册
// // @Description Union用户注册
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.UniRegReq true "data"
// // @Success 200 {object} response.Response{data=models.User} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/uniReg [post]
// func (e SSO) UnionReg(c *gin.Context) {
// 	req := dto.UniRegReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	ip := common.GetClientIP(c)
// 	if logOk, err := s.UniReg(&req, ip); err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.OK(logOk, "成功")
// 	}
// }

// // 忘记密码
// // ForgetPwd 忘记密码
// // @Summary 忘记密码
// // @Description 忘记密码
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.ForgetPwdReq true "data"
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/forgetPwd [post]
// func (e SSO) ForgetPwd(c *gin.Context) {
// 	req := dto.ForgetPwdReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}

// 	if !regexputil.CheckPwd(req.Password) {
// 		e.Error(c, constn.ErrPasswordFMT)
// 		return
// 	}

// 	var mobile, email string
// 	//是否手机
// 	if regexputil.CheckMobile(req.Username) {
// 		s := service.SmsLog{}
// 		e.MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 		mobile = req.Username
// 	} else if regexputil.CheckEmail(req.Username) { //是否邮箱
// 		s := service.EmailLog{}
// 		e.MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 		email = req.Username
// 	} else {
// 		e.Error(c, constn.ErrMobileOrEmail)
// 		return
// 	}
// 	if err := s.ChangePwd(mobile, email, req.Password); err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.OK("", "修改成功")
// 	}

// }

// // 获取用户信息
// // GetUserInfo 获取用户信息
// // @Summary 获取用户信息
// // @Description 获取用户信息
// // @Tags sso
// // @Param data body dto.IdReq true "data"
// // @Success 200 {object} response.Response{data=models.User} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/getUserinfo [post]
// func (e SSO) GetUserInfo(c *gin.Context) {
// 	req := dto.IdReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		MakeOrm().
// 		Bind(&req).
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	var object models.User

// 	err = s.GetByUserId(req.Id, &object)
// 	if err != nil {
// 		e.Error(500, err, fmt.Sprintf("获取用户表失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	resp := dto.UserinfoResp{}
// 	if err := copier.Copy(&resp, object); err != nil {
// 		e.Error(500, err, fmt.Sprintf("拷贝失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}

// 	e.OK(resp, "查询成功")
// }

// // 获取个人信息
// // GetUserInfo 获取个人信息
// // @Summary 获取个人信息
// // @Description 获取个人信息
// // @Tags sso
// // @Success 200 {object} response.Response{data=models.User} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/auth/myUserinfo [post]
// // @Security Bearer
// func (e SSO) MyUserInfo(c *gin.Context) {
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Error(c, constn.ErrUnLogin)
// 		return
// 	}
// 	var object models.User
// 	err = s.GetByUserId(e.GetUserId(), &object)
// 	if err != nil {
// 		e.Error(500, err, fmt.Sprintf("获取用户表失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	resp := dto.MyUserinfoResp{}
// 	if err := copier.Copy(&resp, object); err != nil {
// 		e.Error(500, err, fmt.Sprintf("拷贝失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	e.OK(resp, "查询成功")

// }

// // 修改密码
// // ChangePwd 修改密码
// // @Summary 修改密码
// // @Description 修改密码
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.ChangePwdReq true "data"
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/auth/changePwd [post]
// // @Security Bearer
// func (e SSO) ChangePwd(c *gin.Context) {
// 	req := dto.ChangePwdReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}

// 	if !regexputil.CheckPwd(req.NewPassword) {
// 		e.Error(c, constn.ErrPasswordFMT)
// 		return
// 	}

// 	if e.GetUserId() == "" {
// 		e.Error(c, constn.ErrUnLogin)
// 		return
// 	}
// 	if err := s.ChangePwdByOld(e.GetUserId(), req.OldPassword, req.NewPassword, req.InviteCode); err != nil {
// 		e.Error(500, err, fmt.Sprintf("修改密码失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	}
// 	e.OK("", "修改成功")
// }

// // 绑定手机号或者邮箱
// // Bind 绑定手机号或者邮箱
// // @Summary 绑定手机号或者邮箱
// // @Description 绑定手机号或者邮箱
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.BindReq true "data"
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/auth/bind [post]
// // @Security Bearer
// func (e SSO) Bind(c *gin.Context) {
// 	req := dto.BindReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Error(c, constn.ErrUnLogin)
// 		return
// 	}
// 	//是否手机
// 	if regexputil.CheckMobile(req.Username) {
// 		s := service.SmsLog{}
// 		e.MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 	} else if regexputil.CheckEmail(req.Username) { //是否邮箱
// 		s := service.EmailLog{}
// 		e.MakeService(&s.Service)
// 		if !s.Verify(req.Username, req.Code) {
// 			e.Error(c, constn.ErrVerifyCode)
// 			return
// 		}
// 	} else {
// 		e.Error(c, constn.ErrMobileOrEmail)
// 		return
// 	}

// 	if err := s.Bind(e.GetUserId(), &req); err != nil {
// 		e.Error(c, constn.ErrBind)
// 		return
// 	}

// 	e.OK("", "绑定成功")
// }

// // 修改用户信息
// // ChangeUserinfo 修改用户信息
// // @Summary 修改用户信息
// // @Description 修改用户信息
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.ChangeUserinfoReq true "data"
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/auth/changeUserinfo [post]
// // @Security Bearer
// func (e SSO) ChangeUserinfo(c *gin.Context) {
// 	req := dto.ChangeUserinfoReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Error(c, constn.ErrUnLogin)
// 		return
// 	}
// 	//是否手机

// 	if err := s.ChangeUserinfo(e.GetUserId(), &req); err != nil {
// 		e.Error(c, constn.ErrBind)
// 		return
// 	}
// 	e.OK("", "修改成功")
// }

// // // 退出登录
// // // Logout 退出登录
// // // @Summary 退出登录
// // // @Description 退出登录（调用后清空本地token）
// // // @Tags sso
// // // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // // @Router /v2/sso/auth/logout [post]
// // // @Security Bearer
// // func (e SSO) Logout(c *gin.Context) {

// // 	e.OK("", "退出成功")
// // }

// // 微信登录
// // LoginByWechat 微信登录
// // @Summary 微信登录
// // @Description 微信登录
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} response.Response{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/loginWechat [post]
// func (e SSO) LoginByWechat(c *gin.Context) {
// 	req := dto.LoginWechatReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)

// 	if logOk, err := s.LoginWechat(&req, ip); err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.OK(logOk, "登录成功")
// 	}
// }

// // 获取钉钉登录配置信息
// // GeDingCfg 获取钉钉登录配置信息
// // @Summary 获取钉钉登录配置信息
// // @Description 获取钉钉登录配置信息
// // @Tags sso
// // @Success 200 {object} response.Response{data=dto.DingCfgResp} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/getDingCfg [post]
// func (e SSO) GetDingCfg(c *gin.Context) {
// 	err := e.MakeContext(c).MakeOrm().
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	domain := "http://" + c.Request.Host
// 	var cfg dto.DingCfgResp
// 	service.GetDingConfig(domain, &cfg)
// 	e.OK(cfg, "OK")
// }

// // 钉钉
// // LoginByDing 钉钉登录
// // @Summary 钉钉
// // @Description 钉钉登录
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginDingReq true "data"
// // @Success 200 {object} response.Response{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/loginDing [post]
// func (e SSO) LoginByDing(c *gin.Context) {
// 	req := dto.LoginDingReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)

// 	if logOk, err := s.LoginDing(&req, ip); err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.OK(logOk, "登录成功")
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
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/auth/bindDing [post]
// // @Security Bearer
// func (e SSO) BindDing(c *gin.Context) {
// 	req := dto.LoginDingReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Error(c, constn.ErrUnLogin)
// 		return
// 	}

// 	if err := s.BindDing(&req, e.GetUserId()); err != nil {
// 		e.Error(c, constn.ErrBind)
// 		return
// 	}
// 	e.OK("", "成功")
// }

// // 绑定微信
// // BindWechat 绑定微信
// // @Summary 绑定微信
// // @Description 绑定微信
// // @Tags sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} response.Response{} "{"code": 200, "data": [...]}"
// // @Router /v2/sso/auth/bindWechat [post]
// // @Security Bearer
// func (e SSO) BindWechat(c *gin.Context) {
// 	req := dto.LoginWechatReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		e.Logger.Error(err)
// 		e.Error(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Error(c, constn.ErrUnLogin)
// 		return
// 	}

// 	if err := s.BindWechat(&req, e.GetUserId()); err != nil {
// 		e.Error(c, constn.ErrBind)
// 		return
// 	}
// 	e.OK("", "成功")
// }
