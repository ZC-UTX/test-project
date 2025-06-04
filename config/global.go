package config

import (
	"context"
	__ "github.com/zchengutx/testproject/topics"
	__2 "github.com/zchengutx/testproject/works"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Config      AppConfig
	Log         = zap.NewExample()
	TopicClient __.TopicClient
	Ctx         = context.Background()
	WorksClient __2.WorksClient
)
