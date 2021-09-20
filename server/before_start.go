package server

import (
	"github.com/UrlShortener/internal/keyprovider"
	"github.com/UrlShortener/storage/redis"
)

// BeforeStart ...
func BeforeStart() {
	//	TODO: add logic for range distribution
	redis.Initialize()
	uniqueRangeStart, uniqueRangeEnd := redis.GetRedisStorage().GetUniqueRange()
	keyprovider.Initialize(uniqueRangeStart, uniqueRangeEnd)
}
