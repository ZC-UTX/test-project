package config

import (
	"context"
<<<<<<< HEAD
=======
	__ "github.com/zchengutx/testproject/topics"
	__2 "github.com/zchengutx/testproject/works"
>>>>>>> 6766ff67bf9a5af2a76cccacd8606666467b3bd0
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
<<<<<<< HEAD
	DB  *gorm.DB
	Log = zap.NewExample()
	Ctx = context.Background()
=======
	DB          *gorm.DB
	Config      AppConfig
	Log         = zap.NewExample()
	TopicClient __.TopicClient
	Ctx         = context.Background()
	WorksClient __2.WorksClient
>>>>>>> 6766ff67bf9a5af2a76cccacd8606666467b3bd0
)
