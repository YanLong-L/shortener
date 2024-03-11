package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map
	Sequence      sequence.Sequence
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	seq := sequence.NewMysqlSequence(c.SequenceDB.DSN)
	return &ServiceContext{
		Config:        c,
		Sequence:      seq,
		ShortUrlModel: model.NewShortUrlMapModel(conn, c.CacheRedis),
	}
}
