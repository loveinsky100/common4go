/**
 * @author leo
 * @date 2020/8/18 11:24 上午
 */
package numbers

import "reflect"

type Operate func(value interface{}, targetType Type, defaultValue interface{}) interface{}

type Type string
type Value interface {}

type Transform interface {
	operate(value interface{}, targetType Type, defaultValue interface{}) interface{}
}

type Transformer struct {
	// current type
	Source Type

	// current to target transform function map
	TargetOperateMap map[Type]Operate
}

func (transformer *Transformer) operate(value interface{}, targetType Type, defaultValue interface{}) interface{} {
	if len(transformer.TargetOperateMap) == 0 {
		return defaultValue
	}

	operate, ok := transformer.TargetOperateMap[targetType]
	if !ok {
		return defaultValue
	}

	return operate(value, targetType, defaultValue)
}

// type transform context
type TransformContext interface {
	Transform

	// register into context
	register(transformer *Transformer)
}

type TransformManager struct {
	// source type to transformer
	TransformerMap map[Type]*Transformer
}

func (manager *TransformManager) register(transformer *Transformer) {
	if nil == manager.TransformerMap {
		manager.TransformerMap = make(map[Type]*Transformer)
	}

	manager.TransformerMap[transformer.Source] = transformer
}

func (manager *TransformManager) transform(value Value, target Type, defaultValue Value) Value {
	if nil == manager.TransformerMap {
		return defaultValue
	}

	valueType := reflect.TypeOf(value)
	if valueType == nil {
		return defaultValue
	}

	transformer, ok := manager.TransformerMap[Type(valueType.Name())]
	if !ok {
		return defaultValue
	}

	return transformer.operate(value, target, defaultValue)
}