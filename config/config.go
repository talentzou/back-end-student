package config

type Server struct {
	Captcha   Captcha      `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mysql     Mysql        `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT       JWT          `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
