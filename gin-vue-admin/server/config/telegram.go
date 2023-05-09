package config

type Telegram struct {
	BotToken  string `mapstructure:"bottoken" json:"bottoken" yaml:"bottoken"`
	ChannelId int64  `mapstructure:"channelid" json:"channelid" yaml:"channelid"`
}
