package cache

import (
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-go-common-shared/cache"
	"time"
)

// ProvideRedisCache is a function to provide a redis cache
func ProvideRedisCache() (r *cache.RedisCache) {
	wire.Build(
		cache.NewRedisCache(cache.WithTimeout(3 * time.Second)),
	)
	return
}
