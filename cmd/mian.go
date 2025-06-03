package main

import (
	init_grpc "github.com/zchengutx/testproject/init-grpc"
	init_mysql "github.com/zchengutx/testproject/init-mysql"
	init_viper "github.com/zchengutx/testproject/init-viper"
)

func main() {
	init_viper.InitViper()
	init_mysql.InitMysql()
	init_grpc.InitGrpc()
}
