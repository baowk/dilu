package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	coreConst "github.com/baowk/dilu-core/common/consts"

	"github.com/baowk/dilu-core/common/utils/cryptos"
	"github.com/baowk/dilu-core/common/utils/regexps"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysUser struct {
	*base.BaseService
}

var SerSysUser = SysUser{
	base.NewService(coreConst.DB_DEF),
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	err := core.DB().Debug().Preload("Dept").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	return nil
}

// Get 获取SysUser对象
func (e *SysUser) Get(id int, model *models.SysUser) error {
	var data models.SysUser

	err := core.DB().Model(&data).Debug().
		First(model, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	if err != nil {
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	return nil
}

// // Insert 创建SysUser对象
// func (e *SysUser) Insert(c *dto.SysUserInsertReq) error {
// 	var err error
// 	var data models.SysUser
// 	var i int64
// 	err = core.DB().Model(&data).Where("username = ?", c.Username).Count(&i).Error
// 	if err != nil {
// 		core.Log.Error("db error: %s", zap.Error(err))
// 		return err
// 	}
// 	if i > 0 {
// 		err := errors.New("用户名已存在！")
// 		core.Log.Error("db error: %s", zap.Error(err))
// 		return err
// 	}
// 	c.Generate(&data)
// 	err = core.DB().Create(&data).Error
// 	if err != nil {
// 		core.Log.Error("db error: %s", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }

// // Update 修改SysUser对象
// func (e *SysUser) Update(c *dto.SysUserUpdateReq) error {
// 	var err error
// 	var model models.SysUser
// 	db := core.DB().First(&model, c.GetId())
// 	if err = db.Error; err != nil {
// 		core.Log.Error("Service UpdateSysUser error: %s", zap.Error(err))
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")

// 	}
// 	c.Generate(&model)
// 	update := core.DB().Model(&model).Where("id = ?", &model.Id).Omit("password", "salt").Updates(&model)
// 	if err = update.Error; err != nil {
// 		core.Log.Error("db error: %s", zap.Error(err))
// 		return err
// 	}
// 	if update.RowsAffected == 0 {
// 		err = errors.New("update userinfo error")
// 		core.Log.Warn("db update error")
// 		return err
// 	}
// 	return nil
// }

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq) error {
	var err error
	var model models.SysUser
	db := core.DB().First(&model, c.GetId())
	if err = db.Error; err != nil {
		core.Log.Error("Service UpdateSysUser error: %s", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = core.DB().Table(model.TableName()).Where("id =? ", c.Id).Updates(c).Error
	if err != nil {
		core.Log.Error("Service UpdateSysUser error: %s", zap.Error(err))
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq) error {
	var err error
	var model models.SysUser
	db := core.DB().First(&model, c.GetId())
	if err = db.Error; err != nil {
		core.Log.Error("Service UpdateSysUser error: %s", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = core.DB().Table(model.TableName()).Where("id =? ", c.Id).Updates(c).Error
	if err != nil {
		core.Log.Error("Service UpdateSysUser error: %s", zap.Error(err))
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq) error {
	var err error
	var model models.SysUser
	db := core.DB().First(&model, c.GetId())
	if err = db.Error; err != nil {
		core.Log.Error("At Service ResetSysUserPwd error: %s", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	err = core.DB().Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
	if err != nil {
		core.Log.Error("At Service ResetSysUserPwd error: %s", zap.Error(err))
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(id int) error {
	var err error
	var data models.SysUser

	db := core.DB().Model(&data).Delete(&data, id)
	if err = db.Error; err != nil {
		core.Log.Error("Error found in  RemoveSysUser : %s", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string) errs.IError {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = core.DB().Model(c).Select("Id", "Password").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrNotFound(strconv.Itoa(id), "sysuser", "", err)
		}
		core.Log.Error("db error: %s", zap.Error(err))
		return codes.ErrSys(err)
	}
	if !c.CompPwd(oldPassword) {
		return errs.ErrWithCode(codes.ErrPwd)
	}
	c.Password = newPassword
	db := core.DB().Model(c).Where("id = ?", id).
		Select("Password", "Salt").
		Updates(c)
	if err = db.Error; err != nil {
		core.Log.Error("db error", zap.Error(err))
		return codes.ErrSys(err)
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		core.Log.Warn("db update error")
		return codes.ErrSys(err)
	}
	return nil
}

// 手机注册次数
func (e *SysUser) CountByPhone(phone string, count *int64) error {
	var data models.SysUser
	err := core.DB().Model(&data).Where("phone = ?", phone).Count(count).Error
	if err != nil {
		core.Log.Error("db error", zap.Error(err))
		return err
	}
	return nil
}

// 手机注册次数
func (e *SysUser) CountByEmail(email string, count *int64) error {
	var data models.SysUser
	err := core.DB().Model(&data).Where("email = ?", email).Count(count).Error
	if err != nil {
		core.Log.Error("db error", zap.Error(err))
		return err
	}
	return nil
}

// 注册用户
func (e *SysUser) Register(loginType int, c *dto.RegisterReq, ip string) (dto.LoginOK, errs.IError) {
	model := models.SysUser{}
	lok := dto.LoginOK{}
	var count int64
	if loginType == 1 {
		if err := e.CountByPhone(c.Username, &count); err != nil {
			return lok, codes.ErrSys(err)
		}
		if count > 0 {
			return lok, errs.ErrWithCode(codes.PhoneExistErr)
		}
		model.Phone = c.Username
		if c.Name == "" {
			c.Name = c.Username
		}
	} else {
		if err := e.CountByEmail(c.Username, &count); err != nil {
			return lok, codes.ErrSys(err)
		}
		if count > 0 {
			return lok, errs.ErrWithCode(codes.EmailExistErr)
		}
		model.Email = c.Username
		if c.Name == "" {
			arr := strings.Split(c.Username, "@")
			c.Name = arr[0]
		}
	}
	model.Password = c.Password
	model.Nickname = c.Name
	model.CreatedAt = time.Now()
	model.UpdatedAt = model.CreatedAt
	err := core.DB().Create(&model).Error
	if err != nil {
		core.Log.Error("UserService Insert error", zap.Error(err))
		return lok, codes.ErrSys(err)
	}
	//go e.SendRegDingBot(model)
	return e.loginOK(&model, 0)
}

func (e *SysUser) loginOK(u *models.SysUser, need int) (dto.LoginOK, errs.IError) {
	exp := time.Now().Add(time.Duration(core.Cfg.JWT.Expires) * time.Minute)
	claims := utils.NewClaims(u.Id, exp, core.Cfg.JWT.Issuer, core.Cfg.JWT.Subject)
	claims.Phone = u.Phone
	claims.Nickname = u.Nickname
	claims.RoleId = u.PlatformRoleId
	token, err := utils.Generate(claims, core.Cfg.JWT.SignKey)
	lok := dto.LoginOK{}
	if err != nil {
		return lok, errs.Err(codes.FAILURE, "", err)
	}
	lok.Expire = exp
	lok.AccessToken = token
	lok.Need = need
	if u.PlatformRoleId != 0 {
		lok.Roles = []string{strconv.Itoa(u.PlatformRoleId)}
	}

	if u.Nickname != "" {
		lok.Username = u.Nickname
	} else if u.Username != "" {
		lok.Username = u.Username
	} else if u.Phone != "" {
		lok.Username = u.Phone
	} else if u.Email != "" {
		lok.Username = u.Email
	}
	claims.ExpiresAt(exp.Add(time.Hour * 24 * 7))
	refT, _ := utils.Generate(claims, core.Cfg.JWT.SignKey)
	lok.RefreshToken = refT
	return lok, nil
}

// 通过密码登录
func (e *SysUser) LoginPwd(c *dto.LoginReq, ip string) (dto.LoginOK, errs.IError) {
	model := models.SysUser{}
	lok := dto.LoginOK{}
	if regexps.CheckMobile(c.Username) {
		if err := e.GetByPhone(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
			}
			return lok, errs.ErrWithCode(codes.FAILURE)
		}
	} else if regexps.CheckEmail(c.Username) { //是否邮箱
		if err := e.GetByEmail(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
			}
			return lok, errs.ErrWithCode(codes.FAILURE)
		}
	} else { //用户名密码登录
		if err := e.GetByUsername(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
			}
			return lok, errs.ErrWithCode(codes.FAILURE)
		}
	}
	if model.Password == "" {
		return lok, errs.ErrWithCode(codes.PwdNotExist)
	}
	if !model.CompPwd(c.Password) {
		return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
	}
	if c.UUID != "" {
		e.bindById(c.UUID, model)
	}
	return e.loginOK(&model, 0)
}

// 通过验证码
func (e *SysUser) LoginCode(c *dto.LoginReq, ip string) (dto.LoginOK, errs.IError) {
	var model models.SysUser
	lok := dto.LoginOK{}
	var name string
	if regexps.CheckMobile(c.Username) {
		if err := e.GetByPhone(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, codes.ErrSys(err)
			} else {
				model.Phone = c.Username
				name = c.Username
			}
		}
	} else if regexps.CheckEmail(c.Username) { //是否邮箱
		if err := e.GetByEmail(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, codes.ErrNotFound(c.Username, "sysuser", "", err)
			} else {
				model.Email = c.Username
				arr := strings.Split(c.Username, "@")
				name = arr[0]
			}
		}
	} else {
		return lok, errs.ErrWithCode(codes.ErrMobileOrEmail)
	}
	if model.Id == 0 {
		model.CreatedAt = time.Now()
		model.UpdatedAt = model.CreatedAt
		model.Nickname = name

		err := core.DB().Create(&model).Error
		if err != nil {
			core.Log.Error("sysuser", zap.Error(err))
			return lok, codes.ErrSys(err)
		}
		if c.UUID != "" {
			e.bindById(c.UUID, model)
		}
		//go e.SendRegDingBot(model)
		return e.loginOK(&model, 2)
	}
	if c.UUID != "" {
		e.bindById(c.UUID, model)
	}
	return e.loginOK(&model, 0)
}

// Get 获取User对象
func (e *SysUser) GetByEmail(email string, model *models.SysUser) error {
	var data models.SysUser
	err := core.DB().Model(&data).Where("email = ?", email).First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		//err = errors.New("当前账号不存在，请先注册")
		core.Log.Error("sysuser", zap.Error(err))
		return err
	}
	if err != nil {
		core.Log.Error("sysuser", zap.Error(err))
		return err
	}
	return nil
}

// Get 获取User对象
func (e *SysUser) GetByPhone(mobile string, model *models.SysUser) error {
	//var data models.SysUser
	return core.DB().Where("phone = ?", mobile).First(model).Error
}

func (e *SysUser) bindById(enCode string, user models.SysUser) error {
	dstr, err := cryptos.RSA_Decrypt(enCode, consts.PriKey)
	if err != nil {
		return err
	}
	arr := strings.Split(string(dstr), "-")
	if len(arr) != 2 {
		return errors.New("参数错误")
	}

	var tlm models.ThirdLogin
	id, err := strconv.Atoi(arr[1])
	if err != nil {
		return errors.New("参数错误")
	}
	err = SerThirdLogin.GetById(id, &tlm)
	if err != nil {
		return err
	}
	if tlm.Id < 1 {
		return errors.New("参数错误")
	}

	if err := SerThirdLogin.UpdateUserId(user.Id, tlm); err != nil {
		return err
	}
	return nil

}

// 通过验证码
func (e *SysUser) ChangePwd(mobile, email, password string) errs.IError {

	var user models.SysUser
	if mobile != "" {
		if err := e.GetByPhone(mobile, &user); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return codes.ErrSys(err)
			} else {
				return errs.ErrWithCode(codes.UserNotExist)
			}
		}
	}
	if email != "" {
		if err := e.GetByEmail(email, &user); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return codes.ErrSys(err)
			} else {
				return errs.ErrWithCode(codes.UserNotExist)
			}
		}
	}
	enPwd, err := user.GenPwd(password)
	if err != nil {
		return codes.ErrSys(err)
	}
	updates := models.SysUser{
		Password: string(enPwd),
	}
	updates.UpdatedAt = time.Now()
	db := core.DB().Model(&user).Updates(updates)
	if err = db.Error; err != nil {
		core.Log.Error("sysuser", zap.Error(err))
		return codes.ErrSys(err)
	}
	return nil
}

// 微信登录
func (e *SysUser) LoginWechatMp(req dto.MpSceneReq, openId, ip string) (dto.LoginOK, errs.IError) {

	lok := dto.LoginOK{}

	var tl models.ThirdLogin
	if err := SerThirdLogin.GetTL(3, openId, "", &tl); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return lok, codes.ErrSys(err)
		}
	}
	var user models.SysUser
	if tl.Id < 1 {
		tl.OpenId = openId
		tl.Platform = 3
		tl.CreatedAt = time.Now().Unix()
		tl.UpdatedAt = tl.CreatedAt
		if err := SerThirdLogin.Create(&tl); err != nil {
			return lok, codes.ErrSys(err)
		}
		needMobile(tl.Platform, tl.Id, &lok)
		return lok, nil
	} else {
		if tl.Id == 0 {
			needMobile(tl.Platform, tl.Id, &lok)
			return lok, nil
		}
		if err := e.Get(tl.Id, &user); err != nil {
			return lok, codes.ErrSys(err)
		}
	}
	return e.loginOK(&user, 0)

}

// 钉钉登录
func (e *SysUser) LoginDing(c *dto.LoginDingReq, id string) (dto.LoginOK, errs.IError) {
	lok := dto.LoginOK{}

	if id == "" {
		return lok, errs.ErrWithCode(codes.ThirdNotScan)
	}
	var user models.SysUser
	var tlModel models.ThirdLogin
	if err := SerThirdLogin.GetTL(2, id, "", &tlModel); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return lok, codes.ErrSys(err)
		}
	}

	if tlModel.Id < 1 {
		tlModel.OpenId = id
		tlModel.Platform = 2
		tlModel.CreatedAt = time.Now().Unix()
		tlModel.UpdatedAt = tlModel.CreatedAt
		if err := SerThirdLogin.Create(&tlModel); err != nil {
			return lok, codes.ErrSys(err)
		}
		needMobile(tlModel.Platform, tlModel.Id, &lok)
		return lok, nil
	} else {
		if tlModel.Id == 0 {
			needMobile(tlModel.Platform, tlModel.Id, &lok)
			return lok, nil
		}
		if err := e.Get(tlModel.Id, &user); err != nil {
			return lok, codes.ErrSys(err)
		}
	}
	return e.loginOK(&user, 0)
}

// Get 获取User对象
func (e *SysUser) GetByUsername(username string, model *models.SysUser) errs.IError {
	// str, err := core.Cache.Get("username:" + username)
	// fmt.Println("get:" + str)
	// if err == nil && str != "" {
	// 	if err := json.Unmarshal([]byte(str), model); err == nil {
	// 		return nil
	// 	} else {
	// 		fmt.Println("err:" + err.Error())
	// 	}
	// }
	var data models.SysUser
	err := core.DB().Model(&data).Where("username = ?", username).First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		berr := errs.Err(codes.ErrUserExist, "", err)
		core.Log.Error("sysuser", zap.Error(berr))
		return berr
	}
	if err != nil {
		core.Log.Error("sysuser", zap.Error(err))
		return codes.ErrSys(err)
	}
	// if err := core.Cache.Set("username:"+username, model, time.Hour); err == nil {
	// 	fmt.Println("set================")

	// } else {
	// 	fmt.Println(err)
	// }
	return nil
}

func needMobile(platform, id int, lod *dto.LoginOK) error {
	enS, err := cryptos.RSA_Encrypt([]byte(fmt.Sprintf("%d-%d", platform, id)), consts.PubKey)
	if err != nil {
		return err
	}
	lod.Need = 1
	lod.AccessToken = enS
	return nil
}

// 通过老密码修改
func (e *SysUser) ChangePwdByOld(id int, oldPwd, newPwd, inviteCode string) errs.IError {
	var user models.SysUser
	if id != 0 {
		if err := e.Get(id, &user); err != nil {
			return codes.ErrNotFound(strconv.Itoa(id), "sysuser", "", err)
		}
	}
	if user.Password != "" { //已设置密码
		if !user.CompPwd(oldPwd) {
			return errs.ErrWithCode(codes.ErrPwd)
		}
	}
	enPwd, err := user.GenPwd(newPwd)
	if err != nil {
		return codes.ErrSys(err)
	}
	updates := models.SysUser{
		Password: string(enPwd),
	}
	db := core.DB().Model(user).Updates(updates)
	if err = db.Error; err != nil {
		core.Log.Error("UserService Save error", zap.Error(err))
		return codes.ErrSys(err)
	}
	return nil
}

// 绑定
func (e *SysUser) Bind(id int, c *dto.BindReq) error {
	var user models.SysUser
	if id != 0 {
		if err := e.Get(id, &user); err != nil {
			return errors.New("用户不存在")
		}
	}

	updates := models.SysUser{}

	if regexps.CheckMobile(c.Username) {
		updates.Phone = c.Username
		var count int64
		if err := e.CountByPhone(c.Username, &count); err != nil {
			return err
		}
		if count > 0 {
			return errors.New("该手机号已存在")
		}
	} else if regexps.CheckEmail(c.Username) {
		updates.Email = c.Username
		var count int64
		if err := e.CountByEmail(c.Username, &count); err != nil {
			return err
		}
		if count > 0 {
			return errors.New("该邮箱已存在")
		}
	} else {
		return errors.New("请输入正确的手机号或者邮箱")
	}

	db := core.DB().Model(user).Updates(updates)
	if err := db.Error; err != nil {
		core.Log.Error("UserService Save error", zap.Error(err))
		return err
	}
	return nil
}

// 修改用户信息
func (e *SysUser) ChangeUserinfo(userId int, user models.SysUser) error {
	if user.Password != "" {
		enPwd, err := user.GenPwd(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(enPwd)
	}
	user.UpdateBy = userId
	err := e.UpdateById(user)
	if err != nil {
		core.Log.Error("UserService Save error", zap.Error(err))
		return err
	}
	return nil
}
