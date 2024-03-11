package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/pkg/urltool"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Show 用户访问短链接， 返回重定向响应跳转到对应的longUrl
func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 拿到用户传的 shortUrl
	shortUrl := req.ShortUrl
	// 取base path, 即数据库中保存的shortUrl的值
	basePath := urltool.BasePath(shortUrl)
	// 拿到对应的 原链接
	res, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{basePath, true})
	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.Errorw("svcCtx.ShortUrlModel.FindOneBySurl Failed", logx.Field("err", err))
		}
		return nil, errors.New("404")
	}
	return &types.ShowResponse{
		LongUrl: res.Lurl.String,
	}, nil
}
