package init_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

type Redis struct {
	Addr     string // use default Addr
	Password string
	DB       int
}

func NewRedis(addr, password string, db int) *Redis {
	return &Redis{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
}

func (r *Redis) ExampleNewClient() (*redis.Client, string) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})

	err := Rdb.Ping(context.Background()).Err()

	if err != nil {
		return nil, "redis init failed"
	}
	return Rdb, ""
}
