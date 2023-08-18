package config

type FSCfg struct {
	Bucket           string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                                        //桶
	Region           string `mapstructure:"region" json:"region" yaml:"region"`                                        //区域
	Endpoint         string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`                                  //
	DisableSSL       bool   `mapstructure:"disable-ssl" json:"disable-ssl" yaml:"disable-ssl"`                         //是否开启ssl
	SecretID         string `mapstructure:"secret-id" json:"secret-id" yaml:"secret-id"`                               //访问id accesskey
	SecretKey        string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`                            //安全码
	BaseURL          string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`                                  //基础url
	PathPrefix       string `mapstructure:"path-prefix" json:"path-prefix" yaml:"path-prefix"`                         //路径前缀
	StorePath        string `mapstructure:"store-path" json:"store-path" yaml:"store-path"`                            //本地文件存储路径
	S3ForcePathStyle bool   `mapstructure:"s3-force-path-style" json:"s3-force-path-style" yaml:"s3-force-path-style"` //aws用
	UseCdnDomains    bool   `mapstructure:"use-cdn-domains" json:"use-cdn-domains" yaml:"use-cdn-domains"`             // 七牛 上传是否使用CDN上传加速
}
