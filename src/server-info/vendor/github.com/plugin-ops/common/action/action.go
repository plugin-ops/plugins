package action

import (
	"context"
)

// TODO Action也需要能获取依赖
type Action interface {
	Name() string
	ActionType() ActionType
	Description() string
	Version() int
	OriginalContent() []byte
	Author() string
	GetRelyOn() map[string]string // map[dependency]recipient

	AddRelyOn(recipient string, dependency Action) error
	Compile() error
	Do(ctx context.Context, params ...Parameter) ([]Parameter, error)
}

type ActionType string

const (
	Action_Type_Unknown = "unknown"
)
