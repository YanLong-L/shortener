package sequence

import (
	"context"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

type SnowFlakeSequece struct {
	DCId int64 //数据中心ID
	DId  int64 //机器ID
}

func NewSnowFlakeSequece(dcId int64, dId int64) Sequence {
	return &SnowFlakeSequece{DCId: dcId, DId: dId}
}

func (s *SnowFlakeSequece) Next(ctx context.Context) (int64, error) {
	// 参数1 (int64): 数据中心ID (可用范围:0-31)
	// 参数2 (int64): 机器ID    (可用范围:0-31)
	// 返回1 (*Snowflake): Snowflake对象 | nil
	// 返回2 (error): 错误码
	snow, err := snowflake.NewSnowflake(s.DCId, s.DId)
	if err != nil {
		return 0, err
	}
	return snow.NextVal(), nil
}
