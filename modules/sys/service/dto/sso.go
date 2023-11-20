package dto

import (
	"time"
)

type VerifyCodeReq struct {
	Username string `json:"username" comment:"手机号或者邮箱"` //手机号或者邮箱
	Code     string `json:"code" comment:"验证码"`         //验证码
}

type RegisterReq struct {
	Username   string `json:"username" comment:"手机号或者邮箱"` //手机号或者邮箱
	Password   string `json:"password" comment:"密码"`      //密码
	RePassword string `json:"rePassword" comment:"重复密码"`  //重复密码
	UUID       string `json:"uuid" comment:"uuid"`        //微信或者钉钉第一次登录绑定用
	Code       string `json:"code" comment:"验证码"`         //验证码
	Name       string `json:"name" comment:"真实姓名"`        //真实姓名
	// Client     string `json:"client" comment:"客户端类型"`        //客户端类型 未定义
	// Source     string `json:"source"`                        //来源（最长32位）
	// Os         string `json:"os"`                            //操作系统
	// Inviter    string `json:"inviter" comment:"邀请人"`         //邀请人id
	// InviteType int    `json:"inviteType" comment:"1 企业内部邀人"` //1企业内部邀请 2 用户邀请
}

type LoginReq struct {
	Username string `json:"username" comment:"手机号或者邮箱"`                //手机号或者邮箱
	Password string `json:"password" comment:"密码 密码不为空 为密码登录 否则验证码登录"` //密码 密码不为空 为密码登录 否则验证码登录
	UUID     string `json:"uuid" comment:"微信或者钉钉登录绑定用"`                //微信或者钉钉第一次登录绑定用
	Code     string `json:"code" comment:"验证码"`                        //验证码
	Name     string `json:"name" comment:"真实姓名"`                       //真实姓名
	// Client     string `json:"client" comment:"客户端类型"`                    //客户端类型 未定义
	// Source     string `json:"source"`                                    //来源（最长32位）
	// Os         string `json:"os"`                                        //操作系统
	// Inviter    string `json:"inviter" comment:"邀请人"`                     //邀请人id
	// InviteType int    `json:"inviteType" comment:"1 企业内部邀人"`             //1企业内部邀请 2 用户邀请
}

type LoginWechatReq struct {
	Code       string `json:"code" comment:"微信登录code"`       //微信登录获得code
	Client     string `json:"client" comment:"客户端类型"`        //客户端类型 未定义
	Device     string `json:"device" comment:"客户端设备类型"`      //设备类型 未定义
	Inviter    string `json:"inviter" comment:"邀请人"`         //邀请人id
	InviteType int    `json:"inviteType" comment:"1 企业内部邀人"` //1企业内部邀请 2 用户邀请
	Name       string `json:"name" comment:"真实姓名"`           //真实姓名
}

type LoginDingReq struct {
	Code       string `json:"code" comment:"钉钉登录code"`       //钉钉登录code
	State      string `json:"state" comment:"服务器端返回State"`   //服务器返回state
	Client     string `json:"client" comment:"客户端类型"`        //客户端类型 暂未定义
	Device     string `json:"device" comment:"客户端设备类型"`      //客户端设备类型
	Inviter    string `json:"inviter" comment:"邀请人"`         //邀请人id
	InviteType int    `json:"inviteType" comment:"1 企业内部邀人"` //邀请人类型 1企业内部
	Name       string `json:"name" comment:"真实姓名"`           //真实姓名（在该企业的名字）
}

type LoginOK struct {
	AccessToken  string    `json:"accessToken"`  //返回token
	Expire       time.Time `json:"expires"`      //有效期
	Username     string    `json:"username"`     //token有效期
	Roles        []string  `json:"roles"`        //角色
	RefreshToken string    `json:"refreshToken"` //刷新token
	Need         int       `json:"need"`         //（0|1|2）0：已绑定 1：需要绑定账号 token 是一个加密串，走login接口在uuid上带上 2：需要设置密码
}

type CodeSendReq struct {
	Username   string `json:"username" `  //手机号或者邮箱
	UUID       string `json:"uuid" `      //微信或者钉钉登录绑定用，已绑定传空，详见LoginOK
	Code       string `json:"code"`       //验证码
	CheckExist bool   `json:"checkExist"` //检验账号是否存在
}

type ChangePwdReq struct {
	OldPassword string `json:"oldPwd" comment:"旧密码"`     //老密码
	NewPassword string `json:"newPwd" comment:"新密码"`     //新密码
	RePassword  string `json:"rePwd" comment:"重复新密码"`    //重复新密码
	InviteCode  string `json:"inviteCode" comment:"邀请码"` //邀请码，首次有效
}

type ForgetPwdReq struct {
	Username string `json:"username" comment:"手机号|邮箱"` //手机号|邮箱
	Code     string `json:"code" comment:"验证码"`        //验证码
	Password string `json:"password" comment:"密码"`     //密码
}

type BindReq struct {
	Username string `json:"username" comment:"手机号|邮箱"` //手机号|邮箱
	Code     string `json:"code" comment:"验证码"`        //验证码
}

type UserinfoResp struct {
	UserId     int       `json:"userId"`     //用户详情
	Username   string    `json:"username"`   //用户名
	Phone      string    `json:"phone"`      //手机号
	Email      string    `json:"email"`      //邮箱
	FirstName  string    `json:"firstName"`  //名字
	LastName   string    `json:"lastName"`   //姓
	Nickname   string    `json:"nickname"`   //昵称
	Avatar     string    `json:"avatar"`     //头像
	Bio        string    `json:"bio"`        //签名
	Gender     int       `json:"gender"`     //1 男 2女  3 未知
	Birthday   time.Time `json:"birthday"`   //生日
	Inviter    string    `json:"inviter"`    //邀请人
	InviteType int       `json:"inviteType"` //邀请类型
	CreatedAt  time.Time `json:"createdAt"`  //创建时间
}

type MyUserinfoResp struct {
	UserId   int    `json:"userId"`   //主键
	Username string `json:"username"` //用户名
	Phone    string `json:"phone"`    //手机号
	Email    string `json:"email"`    //邮箱
	NickName string `json:"nickName"` //昵称
	Avatar   string `json:"avatar"`   //头像
	Bio      string `json:"bio"`      //签名
	Gender   int    `json:"gender"`   //性别 1男 2女 3未知
	RoleId   int    `json:"roleId"`   //角色id
	DeptId   int    `json:"deptId"`   //部门id
	Remark   string `json:"remark"`   //备注
	Status   int    `json:"status"`   //状态 1冻结 2正常 3默认
}

type ChangeUserinfoReq struct {
	Name     string `json:"name" `    //名
	Nickname string `json:"nickname"` //昵称
	Avatar   string `json:"avatar"`   //头像
	Bio      string `json:"bio" `     //签名
	Gender   int    `json:"gender"`   //性别 1 男 2女  3 未知
	Birthday string `json:"birthday"` //生日
}

type DingCfgResp struct {
	Appid       string `json:"appid"`       //钉钉Appid
	RedirectUri string `json:"redirectUrl"` //重定向地址
	State       string `json:"state"`       //状态码
}

type MpSceneReq struct {
	Scene           string `json:"scene"`                         //随机参数
	Client          string `json:"client" comment:"客户端类型"`        //客户端类型 未定义
	Source          string `json:"source"`                        //来源（最长32位）
	LastLoginPort   string `json:"last_login_port" `              //登录平台
	LastLoginDevice string `json:"last_login_device"`             //登录设备
	Os              string `json:"os"`                            //操作系统
	Inviter         string `json:"inviter" comment:"邀请人"`         //邀请人id
	InviteType      int    `json:"inviteType" comment:"1 企业内部邀人"` //1企业内部邀请 2 用户邀请
}
