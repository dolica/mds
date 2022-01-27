package config

type Mongo struct {
	Url     string `mapstructure:"url"`
	DB      string `mapstructure:"db"`
	Timeout int    `mapstructure:"timeout"`
}
