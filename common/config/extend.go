package config

var Ext Extend

type Extend struct {
	Demo Demo `mapstructure:"demo" json:"demo" yaml:"demo"`
}

type Demo struct {
	DemoString string
	DemoInt    int
}
