package config

var Ext *Extend

type Extend struct {
	Ding     DingCfg  `mapstructure:"ding" json:"ding" yaml:"ding"`
	WechatMp WechatMp `mapstructure:"wechat-mp" json:"wechat-mp" yaml:"wechat-mp"`
	Ai       Ai       `mapstructure:"ai" json:"ai" yaml:"ai"`
}

type DingCfg struct {
	AgentId   string `mapstructure:"agent-id" json:"agent-id" yaml:"agent-id"`
	AppKey    string `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	AppSecret string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	CropId    string `mapstructure:"crop-id" json:"crop-id" yaml:"crop-id"`
}

type WechatMp struct {
	AppId          string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
	AppSecret      string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	WxToken        string `mapstructure:"wx-token" json:"wx-token" yaml:"wx-token"`
	EncodingAESKey string `mapstructure:"encoding-aes-key" json:"encoding-aes-key" yaml:"encoding-aes-key"`
}

type Ai struct {
	Ali Ali `mapstructure:"ali" json:"ali" yaml:"ali"`
}

type Ali struct {
	SK string `mapstructure:"sk" json:"sk" yaml:"sk"`
}
