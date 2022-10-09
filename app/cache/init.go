package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"yifan/configs"
)

type CacheRepo struct {
	Cache *redis.Client
}

func New() (*CacheRepo, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     configs.GetConfig().Cache.Ip,
		Password: configs.GetConfig().Cache.Password,
	})
	return &CacheRepo{
		Cache: client,
	}, client.Ping(context.Background()).Err()
}
func (d *CacheRepo) GetCache() *redis.Client {
	return d.Cache
}
func (d *CacheRepo) Close() {
	d.Cache.Close()
}
