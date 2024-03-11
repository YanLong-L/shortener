package sequence

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const ReplaceSQL = `REPLACE INTO sequence (stub) VALUES ('a')`

type MysqlSequence struct {
	conn sqlx.SqlConn
}

func NewMysqlSequence(dsn string) Sequence {
	return &MysqlSequence{
		conn: sqlx.NewMysql(dsn),
	}
}

// Next 取下一个号
func (s *MysqlSequence) Next(ctx context.Context) (int64, error) {
	res, err := s.conn.ExecCtx(ctx, ReplaceSQL)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}
