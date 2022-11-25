package config

type Wechat struct {
	AppID      string `yaml:"appid" json:"appid"`
	MerchantID string `json:"merchantId" yaml:"merchantId"`
}
