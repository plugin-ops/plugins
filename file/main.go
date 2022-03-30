package main

import (
	"errors"
	"os"

	"github.com/plugin-ops/pedestal/pedestal/action"
	"github.com/plugin-ops/pedestal/pedestal/plugin"
)

type P struct {
}

func (p *P) Name() string {
	return "file"
}

const (
	ActionRemove = "remove"
)

func (p *P) Do(params ...*action.Value) (map[string]*action.Value, error) {
	if len(params) < 2 {
		return nil, errors.New("incomplete parameters")
	}

	switch params[0].String() {
	case ActionRemove:
		err := os.RemoveAll(params[1].String())
		return nil, err
	default:
		return nil, errors.New("don't know what to do")
	}
}

func (p *P) Version() float32 {
	return 0.1
}

func (p *P) Description() string {
	return `
动作功能: 此插件用于操作文件
传入参数: 
	1. 动作类型(枚举)
		remove: 删除文件
	2. 目标路径(绝对路径)
返回参数(不同动作返回值不同):
	remove: 无
`
}

func main() {
	plugin.ServePlugin(&P{})
}
