package test

import (
	"context"
	"fmt"
	"github.com/plugin-ops/common/action"
)

func Name() string {
	return "TestAction"
}

func Do(ctx context.Context, params ...action.Parameter) ([]action.Parameter, error) {
	fmt.Println(params)

	return []action.Parameter{
		action.NewParameter("Do Test Action"),
	}, nil
}
