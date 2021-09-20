package redis

type RedisStorage interface {
	GetByShortUrl(string) (string, error)
	SaveFullUrl(string, string) error
	GetUniqueRange() (int64, int64)
}
