package redis

import (
	"github.com/UrlShortener/logger"
	"github.com/go-redis/redis"
	"sync"
)

var (
	initRedisStorageOnce sync.Once
	redisStorage         RedisStorage
)

type redisStorageImpl struct {
	redisClient *redis.Client
}

func Initialize() {
	initRedisStorageOnce.Do(func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		redisStorage = &redisStorageImpl{
			redisClient: redisClient,
		}
	})
}

func GetRedisStorage() RedisStorage {
	return redisStorage
}

func (r redisStorageImpl) Set() int64 {
	return 0
}

func (r redisStorageImpl) GetUniqueRange() (int64, int64) {
	eval := r.redisClient.Eval(getUniqueRangeLuaScript, []string{uniqueRangeKey})
	result, err := eval.Result()
	if err != nil {
		logger.Handler.Error(logTag, "Error while getting unique range fro redis %v", err)
		panic(err)
	}
	rangeStart := result.(int64)
	// TODO:// get uniqueKeyRangeLength from config
	rangeEnd := result.(int64) + uniqueKeyRangeLength
	return rangeStart, rangeEnd
}
