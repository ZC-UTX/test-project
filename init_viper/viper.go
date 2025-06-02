package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zchengutx/testproject/config"
	"go.uber.org/zap"
)

func main() {
	viper.SetConfigFile("../config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		config.Log.Error("viper.ReadInConfig failed", zap.Error(err))
	}
	err = viper.Unmarshal(&config.Config)
	if err != nil {
		config.Log.Error("viper.Unmarshal failed", zap.Error(err))

	}
	config.Log.Info(fmt.Sprintf("Video Config %+v", config.Config))
}
