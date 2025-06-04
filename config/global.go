package config

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Log = zap.NewExample()
	Ctx = context.Background()
)
