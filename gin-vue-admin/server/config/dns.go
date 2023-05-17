package config

// 帝恩思 https://www.dns.com/
type Dns struct {
	ApiKey    string `mapstructure:"apikey" json:"apikey" yaml:"apikey"`          // apikey
	ApiSecret string `mapstructure:"apisecret" json:"apisecret" yaml:"apisecret"` // apisecret
}
