package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/middleware"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/common/utils/crypto_util"
	"github.com/baowk/dilu-core/common/utils/regexp_util"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/errs"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	var err error

	err = core.DB().Debug().Preload("Dept").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	return nil
}

// Get 获取SysUser对象
func (e *SysUser) Get(userId int, model *models.SysUser) error {
	var data models.SysUser

	err := core.DB().Model(&data).Debug().
		First(model, userId).Error
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

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = core.DB().Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	c.Generate(&data)
	err = core.DB().Create(&data).Error
	if err != nil {
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	return nil
}

// Update 修改SysUser对象
func (e *SysUser) Update(c *dto.SysUserUpdateReq) error {
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
	c.Generate(&model)
	update := core.DB().Model(&model).Where("user_id = ?", &model.UserId).Omit("password", "salt").Updates(&model)
	if err = update.Error; err != nil {
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		core.Log.Warn("db update error")
		return err
	}
	return nil
}

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
	err = core.DB().Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
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
	err = core.DB().Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
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
func (e *SysUser) Remove(userId int) error {
	var err error
	var data models.SysUser

	db := core.DB().Model(&data).Delete(&data, userId)
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
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = core.DB().Model(c).Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("无权更新该数据")
		}
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(oldPassword))
	if err != nil {
		core.Log.Error("CompareHashAndPassword error", zap.Error(err))
		return err
	}
	c.Password = newPassword
	db := core.DB().Model(c).Where("user_id = ?", id).
		Select("Password", "Salt").
		Updates(c)
	if err = db.Error; err != nil {
		core.Log.Error("db error: %s", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		core.Log.Warn("db update error")
		return err
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
	if enPwd, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost); err != nil {
		return lok, codes.ErrSys(err)
	} else {
		model.Password = string(enPwd)
	}
	model.NickName = c.Name
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
	claims := middleware.NewClaims(u.UserId, exp, core.Cfg.JWT.Issuer, core.Cfg.JWT.Subject)
	claims.Phone = u.Phone
	claims.Nickname = u.NickName
	claims.RoleId = u.RoleId
	token, err := middleware.Generate(claims, core.Cfg.JWT.SignKey)
	lok := dto.LoginOK{}
	if err != nil {
		return lok, errs.Err(codes.FAILURE, "", err)
	}
	lok.Expire = exp
	lok.Token = token
	lok.Need = need
	return lok, nil
}

// 通过密码登录
func (e *SysUser) LoginPwd(c *dto.LoginReq, ip string) (dto.LoginOK, errs.IError) {
	model := models.SysUser{}
	lok := dto.LoginOK{}
	if regexp_util.CheckMobile(c.Username) {
		if err := e.GetByPhone(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
			}
			return lok, errs.ErrWithCode(codes.FAILURE)
		}
	} else if regexp_util.CheckEmail(c.Username) { //是否邮箱
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
	if err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(c.Password)); err != nil {
		core.Log.Error("sysuser", zap.Error(err))
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
	if regexp_util.CheckMobile(c.Username) {
		if err := e.GetByPhone(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, codes.ErrSys(err)
			} else {
				model.Phone = c.Username
				name = c.Username
			}
		}
	} else if regexp_util.CheckEmail(c.Username) { //是否邮箱
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
	if model.UserId == 0 {
		model.CreatedAt = time.Now()
		model.UpdatedAt = model.CreatedAt
		model.NickName = name

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
	dstr := crypto_util.RSA_Decrypt(enCode, consts.PriKey)
	arr := strings.Split(dstr, "-")
	if len(arr) != 2 {
		return errors.New("参数错误")
	}

	var tlm models.ThirdLogin
	id, err := strconv.Atoi(arr[1])
	if err != nil {
		return errors.New("参数错误")
	}
	err = ThirdLoginS.GetById(id, &tlm)
	if err != nil {
		return err
	}
	if tlm.Id < 1 {
		return errors.New("参数错误")
	}

	if err := ThirdLoginS.UpdateUserId(user.UserId, tlm); err != nil {
		return err
	}
	return nil

}

// 通过验证码
func (e *SysUser) ChangePwd(mobile, email, password string) errs.IError {
	enPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return codes.ErrSys(err)
	}
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
	if err := ThirdLoginS.GetTL(3, openId, "", &tl); err != nil {
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
		if err := ThirdLoginS.Create(&tl); err != nil {
			return lok, codes.ErrSys(err)
		}
		needMobile(tl.Platform, tl.Id, &lok)
		return lok, nil
	} else {
		if tl.UserId == 0 {
			needMobile(tl.Platform, tl.Id, &lok)
			return lok, nil
		}
		if err := e.Get(tl.UserId, &user); err != nil {
			return lok, codes.ErrSys(err)
		}
	}
	return e.loginOK(&user, 0)

}

// 钉钉登录
func (e *SysUser) LoginDing(c *dto.LoginDingReq, userId string) (dto.LoginOK, errs.IError) {
	lok := dto.LoginOK{}

	if userId == "" {
		return lok, errs.ErrWithCode(codes.ThirdNotScan)
	}
	var user models.SysUser
	var tlModel models.ThirdLogin
	if err := ThirdLoginS.GetTL(2, userId, "", &tlModel); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return lok, codes.ErrSys(err)
		}
	}

	if tlModel.Id < 1 {
		tlModel.OpenId = userId
		tlModel.Platform = 2
		tlModel.CreatedAt = time.Now().Unix()
		tlModel.UpdatedAt = tlModel.CreatedAt
		if err := ThirdLoginS.Create(&tlModel); err != nil {
			return lok, codes.ErrSys(err)
		}
		needMobile(tlModel.Platform, tlModel.Id, &lok)
		return lok, nil
	} else {
		if tlModel.UserId == 0 {
			needMobile(tlModel.Platform, tlModel.Id, &lok)
			return lok, nil
		}
		if err := e.Get(tlModel.UserId, &user); err != nil {
			return lok, codes.ErrSys(err)
		}
	}
	return e.loginOK(&user, 0)
}

// Get 获取User对象
func (e *SysUser) GetByUsername(username string, model *models.SysUser) errs.IError {
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
	return nil
}

func needMobile(platform, id int, lod *dto.LoginOK) {
	enS := crypto_util.RSA_Encrypt(fmt.Sprintf("%d-%d", platform, id), consts.PubKey)
	lod.Need = 1
	lod.Token = enS
}
