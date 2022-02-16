package test

import (
	"context"

	"github.com/plugin-ops/common/action"
)

type TestActionSet struct {
}

func (t *TestActionSet) Name() string {
	return "TestActionSet"
}

func (t *TestActionSet) Do(ctx context.Context, actionName string, params ...action.Parameter) ([]action.Parameter, error) {
	return (&TestAction{}).Do(ctx, params...)
}

func (t *TestActionSet) AddAction(action action.Action) {
	return
}

func (t *TestActionSet) GetAction(actionName string) action.Action {
	return &TestAction{}
}

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
