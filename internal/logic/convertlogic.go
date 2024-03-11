package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/model"
	"shortener/pkg/base62"
	"shortener/pkg/connect"
	"shortener/pkg/md5"
	"shortener/pkg/urltool"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 校验 longUrl 是否为空: 在handler中实现 基于validator
	// 校验 longUrl 是否是可以访问的
	valid := connect.Connect(req.LongUrl)
	if !valid {
		return nil, errors.New("url无法正常访问")
	}
	// 校验 longUrl 是否已经转链过， 如果转链过就直接返回转链过的短链
	md5Val := md5.MD5String(req.LongUrl)
	_, err = l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{
		String: md5Val,
		Valid:  true,
	})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("该url已转链过")
		}
		return nil, err
	}
	// 校验 longUrl 是否是已经是短链， 避免循环转链
	basePath := urltool.BasePath(req.LongUrl)
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: basePath,
		Valid:  true,
	})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("该url已是短链")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl Failed",
			logx.Field("err", err))
		return nil, err
	}
	// 从基于Mysql实现的发号器中取号
	seq, err := l.svcCtx.Sequence.Next(l.ctx)
	if err != nil {
		logx.Errorw("从Mysql发号起取号失败",
			logx.Field("err", err))
		return nil, err
	}
	// 将取到的号转为 62进制
	short := base62.Base62Encode(seq)
	_, err = l.svcCtx.ShortUrlModel.Insert(l.ctx, &model.ShortUrlMap{
		Lurl: sql.NullString{req.LongUrl, true},
		Md5:  sql.NullString{md5Val, true},
		Surl: sql.NullString{short, true},
	})
	if err != nil {
		logx.Errorw("插入ShortUrlModel表失败",
			logx.Field("err", err))
		return nil, err
	}
	shortUrl := l.svcCtx.Config.ShortDomain + "/" + short
	return &types.ConvertResponse{
		ShortUrl: shortUrl,
	}, nil
}
