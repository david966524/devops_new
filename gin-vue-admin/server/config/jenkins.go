package config

type Jenkins struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`                //jenkins地址
	Username string `mapstructure:"username" json:"username" yaml:"username"` //用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
