package config

import (
	__ "github.com/zchengutx/testproject/topics"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Config      AppConfig
	Log         = zap.NewExample()
	TopicClient __.TopicClient
)
