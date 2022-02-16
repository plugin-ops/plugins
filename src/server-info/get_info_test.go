package server_info

import (
	"context"
	"fmt"
	"testing"

	"github.com/plugin-ops/common/action"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	var i []action.Parameter
	_, err := Do(context.TODO(), action.NewParameter("local"), action.NewParameter("cpu-name"))
	assert.NoError(t, err)
	_, err = Do(context.TODO(), action.NewParameter("local"), action.NewParameter("disk-usage"), action.NewParameter("./"))
	assert.NoError(t, err)
	_, err = Do(context.TODO(), action.NewParameter("local"), action.NewParameter("path-size"), action.NewParameter("./"))
	assert.NoError(t, err)
	_, err = Do(context.TODO(), action.NewParameter("local"), action.NewParameter("read-dir"), action.NewParameter("./"))
	assert.NoError(t, err)
	if i != nil {
		fmt.Println(i)
	}
}
