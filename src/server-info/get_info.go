package server_info

import (
	"context"
	"errors"
	"github.com/plugin-ops/common/action"
)

var DontKnowWhatToDoError = errors.New("don't know what to do")
var IncompleteParametersError = errors.New("incomplete parameters")

func Name() string {
	return "get-server-info"
}

func Do(ctx context.Context, params ...action.Parameter) ([]action.Parameter, error) {
	if len(params) < 2 {
		return nil, DontKnowWhatToDoError
	}

	var info InfoGetter
	switch params[0].String() {
	case "local":
		info = &Local{}
	case "remote":
		return nil, errors.New("暂不支持")
	default:
		return nil, DontKnowWhatToDoError
	}

	result := []action.Parameter{}
	switch params[1].String() {
	case "cpu-name":
		names, err := info.GetCpuName()
		if err != nil {
			return nil, err
		}
		for _, name := range names {
			result = append(result, action.NewParameter(name))
		}
	case "disk-usage":
		if len(params) < 3 {
			return nil, IncompleteParametersError
		}
		usage, err := info.GetDiskUsage(params[2].String())
		if err != nil {
			return nil, err
		}
		result = append(result, action.NewParameter(usage))
	case "path-size":
		if len(params) < 3 {
			return nil, IncompleteParametersError
		}
		usage, err := info.GetPathSize(params[2].String())
		if err != nil {
			return nil, err
		}
		result = append(result, action.NewParameter(usage))
	case "read-dir":
		if len(params) < 3 {
			return nil, IncompleteParametersError
		}
		usage, err := info.ReadDir(params[2].String())
		if err != nil {
			return nil, err
		}
		result = append(result, action.NewParameter(usage))

	default:
		return nil, DontKnowWhatToDoError
	}

	return result, nil
}

type InfoGetter interface {
	GetCpuName() ([]string, error)
	GetDiskUsage(path string) (float64, error)
	GetPathSize(path string) (float64, error)
	ReadDir(dir string) ([]string, error)
}
