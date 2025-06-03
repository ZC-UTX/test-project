package init

import (
	init_grpc "github.com/zchengutx/testproject/init-grpc"
	init_mysql "github.com/zchengutx/testproject/init-mysql"
	init_redis "github.com/zchengutx/testproject/init-redis"
	init_viper "github.com/zchengutx/testproject/init-viper"
)

func init() {
	init_viper.InitViper()
	init_mysql.InitMysql()
	init_grpc.InitGrpc()
	init_redis.ExampleNewClient()
}
