package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ShortUrlDB struct {
		DSN string
	}
	SequenceDB struct {
		DSN string
	}
	CacheRedis   cache.CacheConf
	ShortDomain  string
	Base64String string
}
