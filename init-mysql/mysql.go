package init_mysql

import (
	"fmt"
	"github.com/zchengutx/testproject/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (m *Mysql) InitMysql() *gorm.DB {

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		config.Log.Error("gorm.Open failed", zap.Error(err))
	}

	config.Log.Info(fmt.Sprintf("Mysql connect success, %+v", config.DB))

	sqlDB, err := config.DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB
}
