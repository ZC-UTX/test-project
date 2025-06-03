package init_redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/zchengutx/testproject/config"
	"go.uber.org/zap"
)

func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,     // use default Addr
		Password: config.Config.Redis.Password, // no password set
		DB:       config.Config.Redis.DB,       // use default DB
	})

	pong, err := rdb.Ping(config.Ctx).Result()
	if err != nil {
		config.Log.Error("redis ping error", zap.Error(err))
		return
	}

	config.Log.Info("redis ping success", zap.String("pong", pong))
}
