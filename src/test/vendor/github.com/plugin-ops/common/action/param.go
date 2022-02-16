package action

import (
	"fmt"
	"reflect"
)

type Parameter struct {
	value interface{}
}

func NewParameter(value interface{}) Parameter {
	// TODO 应当根据value类型返回不同的实现
	return Parameter{value: value}
}

// basis

func (p Parameter) String() string {
	return fmt.Sprintf("%v", p.value)
}

func (p Parameter) IsSlice() bool {
	return reflect.TypeOf(p.value).Kind() == reflect.Array || reflect.TypeOf(p.value).Kind() == reflect.Slice
}

func (p Parameter) IsMap() bool {
	return reflect.TypeOf(p.value).Kind() == reflect.Map
}

// iterator

func (p Parameter) Next() bool {
	return false
}

func (p Parameter) Key() Parameter {
	return NewParameter(p.value)
}

func (p Parameter) Value() Parameter {
	return NewParameter(p.value)
}

// slice or map

func (p Parameter) Get(key Parameter) (Parameter, error) {
	return p, nil
}

func (p Parameter) Set(Key, value Parameter) error {
	return nil
}
