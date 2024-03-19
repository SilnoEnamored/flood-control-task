package client

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type CacheClient struct {
	Cache *redis.Client
}

func NewCacheClient(cache *redis.Client) CacheClient {
	return CacheClient{
		Cache: cache,
	}
}

// redis Check
func (c CacheClient) Check(ctx context.Context, userID int64) (bool, error) {
	var i int64
	userId := strconv.FormatInt(userID, 10)
	res := c.Cache.Get(ctx, userId)
	if err := res.Err(); err != nil {
		if !errors.Is(err, redis.Nil) {
			return false, err
		}
	}

	if res != nil {
		i, _ = res.Int64()
		if i >= Cfg.Limit {
			return false, nil
		}
		i++
	}

	c.Cache.Set(ctx, userId, i, Cfg.Timeout)

	return true, nil
}
