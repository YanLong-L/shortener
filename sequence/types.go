package sequence

import "context"

type Sequence interface {
	Next(ctx context.Context) (int64, error)
}
