package redis

import (
	"fmt"
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

var script = `local key = KEYS[1]

local start = 100000

local result = redis.call('GET', key)
if false == result then
   redis.call('SET', key, 200000)
   return start
end

local newNu = tonumber(result)
local start2  =  newNu + start
redis.call('SET', key, start2)

return newNu`

func (r redisStorageImpl) Set() int64 {
	eval := r.redisClient.Eval(script, []string{"lhhhtytdsssdddddddyh"})
	result, err := eval.Result()
	if err != nil {
		fmt.Print("error*", err)
		panic(err)
	}
	fmt.Println(result)
	return result.(int64)
}

func (r redisStorageImpl) Get() string {
	//val, err := r.redisClient.Get("asss").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key 1" , val)
	//fmt.Println(reflect.TypeOf(val))
	//val, err = r.redisClient.Set("zx", "ab", 0).Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key 2" , val)
	//fmt.Println(reflect.TypeOf(val))

	return "val"
}
