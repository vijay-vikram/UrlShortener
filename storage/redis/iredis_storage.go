package redis

type RedisStorage interface {
	Set() int64
	Get() string
}
