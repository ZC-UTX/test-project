package nacos

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Nacos struct {
		Namespace string `mapstructure:"namespace"`
		Addr      string `mapstructure:"addr"`
		Port      uint64 `mapstructure:"port"`
		Dataid    string `mapstructure:"dataid"`
		Group     string `mapstructure:"group"`
	} `mapstructure:"nacos"`
}

var Conf Config

func InitViper(configPath string) (Config, error) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("viper.ReadInConfig error:", err)
	}
	err = viper.Unmarshal(&Conf)

	return Conf, err
}
