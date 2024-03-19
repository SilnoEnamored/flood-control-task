package main

import (
	"cache_service/client"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

const (
	goroutineNumber = 15
)

func main() {
	ctx := context.Background()

	cache := redis.NewClient(&redis.Options{
		Addr: client.Cfg.Addr,
	})

	cacheClient := client.NewCacheClient(cache)
	postgresClient, err := client.NewPostgresClient()
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(goroutineNumber)
	for goroutine := int64(0); goroutine < goroutineNumber; goroutine++ {
		go func(userId int64) {
			defer wg.Done()
			userId = userId % 5
			for i := 0; i < 150; i++ {
				ok, err := cacheClient.Check(ctx, userId)
				if err != nil {
					panic(err)
				}
				fmt.Printf("redis: user_id: %d limit is %v \n", userId, ok)

				ok, err = postgresClient.Check(ctx, userId)
				if err != nil {
					panic(err)
				}
				fmt.Printf("postgres: user_id: %d limit is %v \n", userId, ok)
			}
		}(goroutine)
	}
	wg.Wait()

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
