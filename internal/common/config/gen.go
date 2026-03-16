package config

type GenCfg struct {
	Enable       bool              `mapstructure:"enable" json:"enable" yaml:"enable"`
	GenFront     bool              `mapstructure:"gen-front" json:"gen-front" yaml:"gen-front"`
	FrontPath    string            `mapstructure:"front-path" json:"front-path" yaml:"front-path"`
	TemplatePath string            `mapstructure:"template-path" json:"template-path" yaml:"template-path"`
	ModuleMap    map[string]string `mapstructure:"module-map" json:"module-map" yaml:"module-map"`
}
