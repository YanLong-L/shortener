package sequence

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisSequence struct {
	client redis.Cmdable
}

func NewRedisSequence(client redis.Cmdable) Sequence {
	return &RedisSequence{
		client: client,
	}
}

func (r RedisSequence) Next(ctx context.Context) (int64, error) {
	id, err := r.client.Incr(ctx, "shortener:sequence").Result()
	if err != nil {
		return 0, err
	}
	return id, nil
}
