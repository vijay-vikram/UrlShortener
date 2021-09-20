package server

import "github.com/UrlShortener/storage/redis"

// BeforeStart ...
func BeforeStart() {
	//	TODO: add logic for range distribution
	redis.Initialize()
}
