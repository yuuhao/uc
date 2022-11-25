package config

var Cfg *ConfigMap

type ConfigMap struct {
	App    App    `yaml:"app" json:"app"`
	Wechat Wechat `yaml:"wechat" json:"wechat"`
	Log    Log    `yaml:"log" json:"log"`
}
