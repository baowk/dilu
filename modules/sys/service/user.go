package service

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/go-admin-team/go-admin-core/sdk/config"
// 	"github.com/go-admin-team/go-admin-core/sdk/service"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"

// 	"go-admin/app/sso/models"
// 	"go-admin/app/sso/service/dto"
// 	"go-admin/common"
// 	"go-admin/common/actions"
// 	"go-admin/common/constn"
// 	"go-admin/common/dingtalk/dingbot"
// 	cDto "go-admin/common/dto"
// 	"go-admin/common/middleware"
// 	"go-admin/common/utils"
// 	cryptoutil "go-admin/common/utils/crypto_util"
// 	"go-admin/common/utils/oapi"
// 	regexputil "go-admin/common/utils/regexps"
// 	extConf "go-admin/config"
// )

// type User struct {
// 	service.Service
// }

// // GetPage 获取User列表
// func (e *User) GetPage(c *dto.UserGetPageReq, p *actions.DataPermission, list *[]models.User, count *int64) error {
// 	var err error
// 	var data models.User

// 	err = e.Orm.Model(&data).
// 		Scopes(
// 			cDto.MakeCondition(c.GetNeedSearch()),
// 			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
// 			actions.Permission(data.TableName(), p),
// 		).
// 		Find(list).Limit(-1).Offset(-1).
// 		Count(count).Error
// 	if err != nil {
// 		e.Log.Errorf("UserService GetPage error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取User对象
// func (e *User) Get(d *dto.UserGetReq, p *actions.DataPermission, model *models.User) error {
// 	var data models.User

// 	err := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).
// 		First(model, d.GetId()).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		err = errors.New("查看对象不存在或无权查看")
// 		e.Log.Errorf("Service GetUser error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Insert 创建User对象
// func (e *User) Insert(c *dto.UserInsertReq) error {
// 	var err error
// 	var data models.User
// 	c.Generate(&data)
// 	err = e.Orm.Create(&data).Error
// 	if err != nil {
// 		e.Log.Errorf("UserService Insert error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Update 修改User对象
// func (e *User) Update(c *dto.UserUpdateReq, p *actions.DataPermission) error {
// 	var err error
// 	var data = models.User{}
// 	e.Orm.Scopes(
// 		actions.Permission(data.TableName(), p),
// 	).First(&data, c.GetId())
// 	c.Generate(&data)

// 	db := e.Orm.Save(&data)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("UserService Save error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")
// 	}
// 	return nil
// }

// // Remove 删除User
// func (e *User) Remove(d *dto.UserDeleteReq, p *actions.DataPermission) error {
// 	var data models.User

// 	db := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).Delete(&data, d.GetId())
// 	if err := db.Error; err != nil {
// 		e.Log.Errorf("Service RemoveUser error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权删除该数据")
// 	}
// 	return nil
// }

// var ErrPwdNotExist = errors.New("密码不存在")

// // 通过密码登录
// func (e *User) LoginPwd(c *dto.LoginReq, ip string) (dto.LoginOK, error) {
// 	model := models.User{}
// 	lok := dto.LoginOK{}
// 	if regexputil.CheckMobile(c.Username) {
// 		if err := e.GetByMobile(c.Username, &model); err != nil {
// 			return lok, err
// 		}

// 	} else if regexputil.CheckEmail(c.Username) { //是否邮箱
// 		if err := e.GetByEmail(c.Username, &model); err != nil {
// 			return lok, err
// 		}
// 	} else { //用户名密码登录
// 		if err := e.GetByUsername(c.Username, &model); err != nil {
// 			return lok, err
// 		}
// 	}
// 	if model.Password == "" {
// 		return lok, ErrPwdNotExist
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(c.Password)); err != nil {
// 		return lok, errors.New("手机号或密码错误，请重新输入")
// 	}
// 	if c.UUID != "" {
// 		e.bindById(c.UUID, model)
// 	}
// 	return e.loginOK(&model, 0)
// }

// // 通过验证码
// func (e *User) LoginCode(c *dto.LoginReq, ip string) (dto.LoginOK, error) {
// 	var model models.User
// 	lok := dto.LoginOK{}
// 	var name string
// 	if regexputil.CheckMobile(c.Username) {
// 		if err := e.GetByMobile(c.Username, &model); err != nil {
// 			if !errors.Is(err, gorm.ErrRecordNotFound) {
// 				return lok, err
// 			} else {
// 				model.Mobile = c.Username
// 				name = c.Username
// 				//model.Bind.Mobile = 1
// 			}
// 		}
// 	} else if regexputil.CheckEmail(c.Username) { //是否邮箱
// 		if err := e.GetByEmail(c.Username, &model); err != nil {
// 			if !errors.Is(err, gorm.ErrRecordNotFound) {
// 				return lok, err
// 			} else {
// 				model.Email = c.Username
// 				arr := strings.Split(c.Username, "@")
// 				name = arr[0]
// 				//model.Bind.Email = 1
// 			}
// 		}
// 	} else {
// 		return lok, errors.New("请输入正确的手机号或者邮箱")
// 	}
// 	if model.Id == 0 {
// 		model.UserId = utils.GenUUid()
// 		model.CreatedAt = time.Now()
// 		model.InviteType = c.InviteType
// 		model.Inviter = c.Inviter
// 		model.UpdatedAt = model.CreatedAt
// 		model.Source = c.Client
// 		if c.Source != "" {
// 			model.Source = c.Source
// 		}
// 		model.LastLoginDevice = c.LastLoginDevice
// 		model.LastLoginPort = c.LastLoginPort
// 		model.Ip = ip
// 		var loc common.IPLocationData
// 		if err := common.GetLocationByIp(ip, &loc); err == nil {
// 			model.IpLocation = loc.Country + loc.Province + loc.City
// 		}

// 		model.Os = c.Os

// 		if c.Name != "" {
// 			model.FirstName = c.Name
// 		} else {
// 			model.FirstName = name
// 		}
// 		err := e.Orm.Create(&model).Error
// 		if err != nil {
// 			e.Log.Errorf("UserService Insert error:%s \r\n", err)
// 			return lok, err
// 		}
// 		if c.UUID != "" {
// 			e.bindById(c.UUID, model)
// 		}
// 		go e.SendRegDingBot(model)
// 		return e.loginOK(&model, 2)
// 	}
// 	if c.UUID != "" {
// 		e.bindById(c.UUID, model)
// 	}
// 	return e.loginOK(&model, 0)
// }

// // 通过验证码
// func (e *User) UniReg(c *dto.UniRegReq, ip string) (models.User, error) {

// 	var model models.User
// 	var name string
// 	if regexputil.CheckMobile(c.Username) {
// 		if err := e.GetByMobile(c.Username, &model); err != nil {
// 			if !errors.Is(err, gorm.ErrRecordNotFound) {
// 				return model, err
// 			}
// 			name = c.Username
// 			model.Mobile = c.Username
// 		}

// 		// } else if regexputil.CheckEmail(c.Username) { //是否邮箱
// 		// 	if err := e.CountByEmail(c.Username, &count); err != nil {
// 		// 		return lok, err
// 		// 	}
// 		// 	if count == 0 {
// 		// 		model.Email = c.Username
// 		// 		arr := strings.Split(c.Username, "@")
// 		// 		name = arr[0]
// 		// 		//model.Bind.Email = 1
// 		// 	} else {
// 		// 		if err := e.GetByEmail(c.Username, &model); err != nil {
// 		// 			return lok, err
// 		// 		}
// 		// 	}
// 	} else {
// 		return model, errors.New("请输入正确的手机号")
// 	}
// 	if model.Id == 0 {
// 		model.UserId = utils.GenUUid()
// 		model.CreatedAt = time.Now()
// 		model.UpdatedAt = model.CreatedAt
// 		model.Source = c.Source

// 		if c.Name != "" {
// 			model.FirstName = c.Name
// 		} else {
// 			model.FirstName = name
// 		}
// 		err := e.Orm.Create(&model).Error
// 		if err != nil {
// 			e.Log.Errorf("UserService Insert error:%s \r\n", err)
// 			return model, err
// 		}
// 		go e.SendRegDingBot(model)
// 		return model, nil
// 	}
// 	return model, nil
// }

// func (e *User) bindById(enCode string, user models.User) error {
// 	dstr := cryptoutil.RSA_Decrypt(enCode, constn.PriKey)
// 	arr := strings.Split(dstr, "-")
// 	if len(arr) != 2 {
// 		return errors.New("参数错误")
// 	}
// 	if arr[0] == "1" && user.Bind.Wechat == 1 {
// 		return errors.New("该账号已绑定微信")
// 	} else {
// 		user.Bind.Wechat = 1
// 	}
// 	if arr[0] == "2" && user.Bind.Ding == 1 {
// 		return errors.New("该账号已绑定钉钉")
// 	} else {
// 		user.Bind.Ding = 1
// 	}
// 	tls := ThirdLogin{}
// 	tls.Orm = e.Orm
// 	var tlm models.ThirdLogin
// 	id, err := strconv.Atoi(arr[1])
// 	if err != nil {
// 		return errors.New("参数错误")
// 	}
// 	err = tls.GetById(id, &tlm)
// 	if err != nil {
// 		return err
// 	}
// 	if tlm.IsEmpty() {
// 		return errors.New("参数错误")
// 	}

// 	if err := tls.UpdateUserId(user.UserId, tlm); err != nil {
// 		return err
// 	}
// 	if err := e.updateBind(user); err != nil {
// 		return err
// 	}
// 	return nil

// }

// // Get 获取User对象
// func (e *User) Register(loginType int, c *dto.RegisterReq, ip string) (dto.LoginOK, error) {
// 	model := models.User{}
// 	lok := dto.LoginOK{}
// 	var count int64
// 	if loginType == 1 {
// 		if err := e.CountByMobile(c.Username, &count); err != nil {
// 			return lok, err
// 		}
// 		if count > 0 {
// 			return lok, errors.New("手机号已存在")
// 		}
// 		model.Mobile = c.Username
// 		if c.Name == "" {
// 			c.Name = c.Username
// 		}
// 		//model.Bind.Mobile = 1
// 	} else {
// 		if err := e.CountByEmail(c.Username, &count); err != nil {
// 			return lok, err
// 		}
// 		if count > 0 {
// 			return lok, errors.New("邮箱已存在")
// 		}
// 		model.Email = c.Username
// 		if c.Name == "" {
// 			arr := strings.Split(c.Username, "@")
// 			c.Name = arr[0]
// 		}
// 		//model.Bind.Email = 1
// 	}
// 	model.UserId = utils.GenUUid()
// 	if enPwd, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost); err != nil {
// 		return lok, errors.New("生成密码错误")
// 	} else {
// 		model.Password = string(enPwd)
// 	}
// 	model.CreatedAt = time.Now()
// 	model.UpdatedAt = model.CreatedAt
// 	model.InviteType = c.InviteType
// 	model.Source = c.Client
// 	if c.Inviter != "" {
// 		model.Inviter = c.Inviter
// 	}
// 	if c.Name != "" {
// 		model.FirstName = c.Name
// 		model.Nickname = c.Name
// 	}
// 	if c.Source != "" {
// 		model.Source = c.Source
// 	}
// 	model.LastLoginDevice = c.LastLoginDevice
// 	model.LastLoginPort = c.LastLoginPort
// 	model.Ip = ip
// 	var loc common.IPLocationData
// 	if err := common.GetLocationByIp(ip, &loc); err == nil {
// 		model.IpLocation = loc.Country + loc.Province + loc.City
// 	}
// 	model.Os = c.Os
// 	err := e.Orm.Create(&model).Error
// 	if err != nil {
// 		e.Log.Errorf("UserService Insert error:%s \r\n", err)
// 		return lok, err
// 	}
// 	go e.SendRegDingBot(model)
// 	return e.loginOK(&model, 0)
// }

// func (e *User) SendRegDingBot(model models.User) {
// 	var username string
// 	if model.Mobile != "" {
// 		username = model.Mobile
// 	} else if model.Email != "" {
// 		username = model.Email
// 	}
// 	var user models.User
// 	inviter := model.Inviter
// 	if err := e.GetByUserId(model.Inviter, &user); err == nil {
// 		if user.Mobile != "" {
// 			inviter = user.Mobile
// 		} else if user.Email != "" {
// 			inviter = user.Email
// 		}
// 		// else if user.Username != "" {
// 		// 	inviter = user.Username
// 		// }
// 	}
// 	dingbot.SendReg(extConf.ExtConfig.Dingbot, username, model.UserId, model.Source, inviter)
// 	oA := oapi.New(extConf.ExtConfig.Yunlogin.CrmUrl)
// 	name := model.FirstName
// 	if name == "" {
// 		name = model.Nickname
// 	}
// 	oA.UserToCrm(model.Id, name, username, model.CreatedAt)
// }

// func (e *User) loginOK(u *models.User, need int) (dto.LoginOK, error) {
// 	expT := time.Now().Add(time.Duration(config.JwtConfig.Timeout) * time.Second)
// 	token, err := middleware.GenToken(u.UserId, "", expT, config.JwtConfig.Secret)
// 	lok := dto.LoginOK{}
// 	if err != nil {
// 		return lok, err
// 	}
// 	lok.Expire = expT
// 	lok.Token = token
// 	lok.Need = need
// 	return lok, nil
// }

// // Get 获取User对象
// func (e *User) GetByUserId(userId string, model *models.User) error {
// 	var data models.User
// 	err := e.Orm.Model(&data).Where("user_id = ?", userId).First(model).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		err = errors.New("查看对象不存在或无权查看")
// 		e.Log.Errorf("Service GetUser error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取User对象
// func (e *User) GetByEmail(email string, model *models.User) error {
// 	var data models.User
// 	err := e.Orm.Model(&data).Where("email = ?", email).First(model).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		//err = errors.New("当前账号不存在，请先注册")
// 		e.Log.Errorf("Service GetUser error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取User对象
// func (e *User) GetByUsername(username string, model *models.User) error {
// 	var data models.User
// 	err := e.Orm.Model(&data).Where("username = ?", username).First(model).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		err = errors.New("当前账号不存在，请先注册")
// 		e.Log.Errorf("Service GetUser error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取User对象
// func (e *User) CountByEmail(email string, count *int64) error {
// 	var data models.User
// 	err := e.Orm.Model(&data).Where("email = ?", email).Count(count).Error
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取User对象
// func (e *User) GetByMobile(mobile string, model *models.User) error {
// 	var data models.User
// 	err := e.Orm.Model(&data).Where("mobile = ?", mobile).First(model).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		//err = errors.New("当前账号不存在，请先注册")
// 		e.Log.Errorf("Service GetUser error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // 该手机注册次数
// func (e *User) CountByMobile(mobile string, count *int64) error {
// 	var data models.User
// 	err := e.Orm.Model(&data).Where("mobile = ?", mobile).Count(count).Error
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // 通过验证码
// func (e *User) ChangePwd(mobile, email, password string) error {
// 	enPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return errors.New("生成密码错误")
// 	}
// 	var user models.User
// 	if mobile != "" {
// 		if err := e.GetByMobile(mobile, &user); err != nil {
// 			return errors.New("用户不存在")
// 		}
// 	}
// 	if email != "" {
// 		if err := e.GetByEmail(email, &user); err != nil {
// 			return errors.New("用户不存在")
// 		}
// 	}
// 	updates := models.User{
// 		Password:  string(enPwd),
// 		UpdatedAt: time.Now(),
// 	}
// 	db := e.Orm.Model(user).Updates(updates)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("UserService Save error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // 通过老密码修改
// func (e *User) ChangePwdByOld(userId, oldPwd, newPwd, inviteCode string) error {
// 	var user models.User
// 	if userId != "" {
// 		if err := e.GetByUserId(userId, &user); err != nil {
// 			return errors.New("用户不存在")
// 		}
// 	}
// 	enPwd, err := bcrypt.GenerateFromPassword([]byte(newPwd), bcrypt.DefaultCost)
// 	if err != nil {
// 		return errors.New("生成密码错误")
// 	}
// 	updates := models.User{
// 		Password:  string(enPwd),
// 		UpdatedAt: time.Now(),
// 	}

// 	if user.Password != "" { //已设置密码
// 		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPwd)); err != nil {
// 			return errors.New("请输入正确的旧密码")
// 		}
// 	} else {
// 		//首次设置密码，可设置邀请码
// 		updates.Inviter = inviteCode
// 	}
// 	db := e.Orm.Model(user).Updates(updates)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("UserService Save error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // 通过老密码修改
// func (e *User) Bind(userId string, c *dto.BindReq) error {
// 	var user models.User
// 	if userId != "" {
// 		if err := e.GetByUserId(userId, &user); err != nil {
// 			return errors.New("用户不存在")
// 		}
// 	}

// 	updates := models.User{
// 		UpdatedAt: time.Now(),
// 	}

// 	if regexputil.CheckMobile(c.Username) {
// 		updates.Mobile = c.Username
// 		var count int64
// 		if err := e.CountByMobile(c.Username, &count); err != nil {
// 			return err
// 		}
// 		if count > 0 {
// 			return errors.New("该手机号已存在")
// 		}
// 	} else if regexputil.CheckEmail(c.Username) {
// 		updates.Email = c.Username
// 		var count int64
// 		if err := e.CountByEmail(c.Username, &count); err != nil {
// 			return err
// 		}
// 		if count > 0 {
// 			return errors.New("该邮箱已存在")
// 		}
// 	} else {
// 		return errors.New("请输入正确的手机号或者邮箱")
// 	}

// 	db := e.Orm.Model(user).Updates(updates)
// 	if err := db.Error; err != nil {
// 		e.Log.Errorf("UserService Save error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // 通过老密码修改
// func (e *User) ChangeUserinfo(userId string, c *dto.ChangeUserinfoReq) error {
// 	var user models.User
// 	if userId != "" {
// 		if err := e.GetByUserId(userId, &user); err != nil {
// 			return errors.New("用户不存在")
// 		}
// 	}

// 	updates := models.User{
// 		UpdatedAt: time.Now(),
// 	}
// 	if c.Avatar != "" {
// 		updates.Avatar = c.Avatar
// 	}
// 	if c.Bio != "" {
// 		updates.Bio = c.Bio
// 	}
// 	if c.Birthday != "" {
// 		if d, err := time.Parse(constn.FmtDate, c.Birthday); err == nil {
// 			updates.Birthday = d
// 		} else {
// 			return errors.New("生日格式错误")
// 		}
// 	}
// 	if c.FirstName != "" {
// 		updates.FirstName = c.FirstName
// 	}
// 	if c.LastName != "" {
// 		updates.LastName = c.LastName
// 	}
// 	if c.Nickname != "" {
// 		updates.Nickname = c.Nickname
// 	}

// 	db := e.Orm.Model(user).Updates(updates)
// 	if err := db.Error; err != nil {
// 		e.Log.Errorf("UserService Save error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // 微信登录
// func (e *User) LoginWechat(c *dto.LoginWechatReq, ip string) (dto.LoginOK, error) {
// 	lok := dto.LoginOK{}
// 	ret, err := AuthWechat(c.Code)
// 	if err != nil {
// 		return lok, err
// 	}
// 	var ld WechatLogin
// 	err = json.Unmarshal([]byte(ret), &ld)
// 	if err != nil {
// 		return lok, err
// 	}

// 	s := ThirdLogin{}
// 	s.Service.Orm = e.Service.Orm

// 	var tlModel models.ThirdLogin
// 	if err = s.GetTL(1, ld.OpenId, ld.Unionid, &tlModel); err != nil {
// 		return lok, err
// 	}
// 	var user models.User

// 	if tlModel.IsEmpty() {
// 		// user.Bind.Wechat = 1
// 		// user.UserId = utils.GenUUid()
// 		// user.CreatedAt = time.Now()
// 		// user.UpdatedAt = user.CreatedAt
// 		// err = e.Orm.Create(&user).Error
// 		// if err != nil {
// 		// 	e.Log.Errorf("UserService Insert error:%s \r\n", err)
// 		// 	return lok, err
// 		// }
// 		tlModel.UnionId = ld.Unionid
// 		tlModel.OpenId = ld.OpenId
// 		//tlModel.UserId = user.UserId
// 		tlModel.Platform = 1
// 		tlModel.CreatedAt = time.Now()
// 		tlModel.UpdatedAt = tlModel.CreatedAt
// 		tlModel.ThirdData = ret
// 		if err = s.Create(tlModel); err != nil {
// 			return lok, err
// 		}
// 		needMobile(tlModel.Platform, tlModel.Id, &lok)
// 		return lok, nil
// 	} else {
// 		if tlModel.UserId == "" {
// 			needMobile(tlModel.Platform, tlModel.Id, &lok)
// 			return lok, nil
// 		}
// 		if err = e.GetByUserId(tlModel.UserId, &user); err != nil {
// 			return lok, err
// 		}
// 	}
// 	return e.loginOK(&user, 0)
// }

// // 钉钉登录
// func (e *User) LoginDing(c *dto.LoginDingReq, ip string) (dto.LoginOK, error) {
// 	lok := dto.LoginOK{}

// 	data, err := AuthDing(c.Code)
// 	if err != nil {
// 		return lok, err
// 	}

// 	var duser DingUser

// 	err = json.Unmarshal(data, &duser)
// 	if err != nil {
// 		log.Println(err)
// 		return lok, errors.New("auth code is null")
// 	}

// 	s := ThirdLogin{}
// 	s.Service.Orm = e.Service.Orm

// 	var user models.User
// 	var tlModel models.ThirdLogin
// 	if err = s.GetTL(2, duser.Userid, "", &tlModel); err != nil {
// 		return lok, err
// 	}

// 	if tlModel.IsEmpty() {
// 		tlModel.OpenId = duser.Userid
// 		tlModel.Platform = 2
// 		tlModel.CreatedAt = time.Now()
// 		tlModel.UpdatedAt = tlModel.CreatedAt
// 		tlModel.ThirdData = string(data)
// 		if err = s.Create(tlModel); err != nil {
// 			return lok, err
// 		}
// 		needMobile(tlModel.Platform, tlModel.Id, &lok)
// 		return lok, nil
// 	} else {
// 		if tlModel.UserId == "" {
// 			needMobile(tlModel.Platform, tlModel.Id, &lok)
// 			return lok, nil
// 		}
// 		if err = e.GetByUserId(tlModel.UserId, &user); err != nil {
// 			return lok, err
// 		}
// 	}
// 	return e.loginOK(&user, 0)
// }

// func needMobile(platform, id int, lod *dto.LoginOK) {
// 	enS := cryptoutil.RSA_Encrypt(fmt.Sprintf("%d-%d", platform, id), constn.PubKey)
// 	lod.Need = 1
// 	lod.Token = enS

// }

// // 微信登录
// func (e *User) BindWechat(c *dto.LoginWechatReq, userId string) error {
// 	var user models.User
// 	err := e.GetByUserId(userId, &user)
// 	if err != nil {
// 		return err
// 	}
// 	if user.Bind.Wechat == 1 {
// 		return errors.New("该账号已绑定微信")
// 	}
// 	ret, err := AuthWechat(c.Code)
// 	if err != nil {
// 		return err
// 	}

// 	var ld WechatLogin
// 	err = json.Unmarshal([]byte(ret), &ld)
// 	if err != nil {
// 		return err
// 	}

// 	s := ThirdLogin{}
// 	s.Service.Orm = e.Service.Orm

// 	var tlModel models.ThirdLogin
// 	if err = s.GetTL(1, ld.OpenId, ld.Unionid, &tlModel); err != nil {
// 		return err
// 	}

// 	if tlModel.IsEmpty() {
// 		tlModel.UnionId = ld.Unionid
// 		tlModel.OpenId = ld.OpenId
// 		tlModel.UserId = user.UserId
// 		tlModel.Platform = 1
// 		tlModel.CreatedAt = time.Now()
// 		tlModel.UpdatedAt = tlModel.CreatedAt
// 		tlModel.ThirdData = ret
// 		if err = s.Create(tlModel); err != nil {
// 			return err
// 		}
// 		user.Bind.Wechat = 1
// 		err := e.updateBind(user)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		return errors.New("该微信已被被其他账号绑定")
// 	}
// 	return nil
// }

// func (e *User) updateBind(user models.User) error {
// 	bind := user.Bind
// 	updates := models.User{
// 		UpdatedAt: time.Now(),
// 		Bind:      bind,
// 	}
// 	return e.Orm.Model(user).Updates(updates).Error
// }

// // 钉钉登录
// func (e *User) BindDing(c *dto.LoginDingReq, userId string) error {
// 	var user models.User
// 	err := e.GetByUserId(userId, &user)
// 	if err != nil {
// 		return err
// 	}
// 	if user.Bind.Ding == 1 {
// 		return errors.New("该账号已绑定钉钉")
// 	}

// 	data, err := AuthDing(c.Code)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("res:%s", data)

// 	var duser DingUser

// 	err = json.Unmarshal(data, &duser)
// 	if err != nil {
// 		fmt.Println(err)
// 		return errors.New("auth code is null")
// 	}

// 	s := ThirdLogin{}
// 	s.Service.Orm = e.Service.Orm

// 	var tlModel models.ThirdLogin
// 	if err = s.GetTL(2, duser.Userid, "", &tlModel); err != nil {
// 		return err
// 	}

// 	if tlModel.IsEmpty() {
// 		tlModel.OpenId = duser.Userid
// 		tlModel.UserId = user.UserId
// 		tlModel.Platform = 2
// 		tlModel.CreatedAt = time.Now()
// 		tlModel.UpdatedAt = tlModel.CreatedAt
// 		tlModel.ThirdData = string(data)
// 		if err = s.Create(tlModel); err != nil {
// 			return err
// 		}

// 		bind := user.Bind
// 		bind.Ding = 1
// 		updates := models.User{
// 			UpdatedAt: time.Now(),
// 			Bind:      bind,
// 		}
// 		err := e.Orm.Model(user).Updates(updates).Error
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		return err
// 	}
// 	return nil
// }
