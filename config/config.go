package config

type Server struct {
	Amap  Amap  `mapstructure:"amap"`
	Mongo Mongo `mapstructure:"mongo"`
	Http  Http  `mapstructure:"http"`
}
