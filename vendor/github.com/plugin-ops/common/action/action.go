package action

import (
	"context"
)

type Action interface {
	Name() string
	Do(ctx context.Context, params ...Parameter) ([]Parameter, error)
}
