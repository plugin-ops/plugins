package test

import (
	"context"

	"github.com/plugin-ops/common/action"
)

type TestAction struct {
}

func (t *TestAction) Name() string {
	return "TestAction"
}

func (t *TestAction) Do(ctx context.Context, params ...action.Parameter) ([]action.Parameter, error) {
	return []action.Parameter{
		action.NewParameter("Do Test Action"),
	}, nil
}
