package config

import (
	"github.com/baowk/dilu-core/config"
	rdConfig "github.com/baowk/dilu-rd/config"
)

var cfg *Config

func init() {
	cfg = new(Config)
}

type Config struct {
	// 明确指定常用的配置字段，避免继承访问问题
	Server      config.ServerCfg     `mapstructure:"server" json:"server" yaml:"server"`
	Remote      config.RemoteCfg     `mapstructure:"remote" json:"remote" yaml:"remote"`
	Logger      config.LogCfg        `mapstructure:"logger" json:"logger" yaml:"logger"`
	DBCfg       config.DBCfg         `mapstructure:"dbcfg" json:"dbcfg" yaml:"dbcfg"`
	Cache       config.CacheCfg      `mapstructure:"cache" json:"cache" yaml:"cache"`
	JWT         config.JWT           `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Cors        config.CORS          `mapstructure:"cors" json:"cors" yaml:"cors"`
	Gen         config.GenCfg        `mapstructure:"gen" json:"gen" yaml:"gen"`
	AccessLimit config.AccessLimit   `mapstructure:"access-limit" json:"access-limit" yaml:"access-limit"`
	GrpcServer  config.GrpcServerCfg `mapstructure:"grpc-server" json:"grpc-server" yaml:"grpc-server"`
	// 扩展配置
	RdConfig rdConfig.Config `mapstructure:"rd-config" json:"rd-config" yaml:"rd-config"`
}

func (c *Config) GetServerCfg() *config.ServerCfg {
	return &c.Server
}

func (c *Config) GetLogCfg() *config.LogCfg {
	return &c.Logger
}
func (c *Config) GetDBCfg() *config.DBCfg {
	return &c.DBCfg
}
func (c *Config) GetCacheCfg() *config.CacheCfg {
	return &c.Cache
}

func SaveConfig(local, remote *Config) {
	if remote != nil {
		*cfg = *remote
		cfg.Server.Mode = local.Server.Mode
		cfg.Server.RemoteEnable = local.Server.RemoteEnable
		cfg.Server.Name = local.Server.Name
		cfg.Server.Port = local.Server.Port
		cfg.Server.Host = local.Server.Host
		cfg.Server.Node = local.Server.Node
	} else {
		*cfg = *local
	}
}

func Get() *Config {
	return cfg
}
