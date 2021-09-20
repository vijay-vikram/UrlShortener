package redis

type RedisStorage interface {
	Set() int64
	GetUniqueRange() (int64, int64)
}
