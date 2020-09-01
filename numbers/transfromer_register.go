/**
 * @author leo
 * @date 2020/8/18 3:07 下午
 */
package numbers

var Context = &TransformManager {}

type MethodTransform func() Operate

// define transform method
type Operator interface {
	// transform to string
	Transform2String(value Value, defaultValue Value) interface{}

	// transform to int64
	Transform2Int64(value Value, defaultValue Value) interface{}

	// transform to bool
	Transform2Bool(value Value, defaultValue Value) interface{}

	// transform ro float64
	Transform2Float64(value Value, defaultValue Value) interface{}
}

type TypeOperator struct {
	// target
	target Type

	// transform method
	methodTransform MethodTransform
}

func RegisterOperator(source Type, operator Operator)  {
	transformer := registerTransformer(source, transform2TypeOperators(operator))
	if nil == transformer {
		return
	}

	Context.register(transformer)
}

func transform2TypeOperators(operator Operator) []*TypeOperator {
	operators := make([]*TypeOperator, 0)
	operators = appendOperator(operators, "string", operator.Transform2String)
	operators = appendOperator(operators, "int64", operator.Transform2Int64)
	operators = appendOperator(operators, "bool", operator.Transform2Bool)
	operators = appendOperator(operators, "float64", operator.Transform2Float64)
	return operators
}

func appendOperator(operators []*TypeOperator, target string, method func(value Value, defaultValue Value) interface{}) []*TypeOperator {
	operators = append(operators, transform2TypeOperator(Type(target), func() Operate {
		return func(value interface{}, targetType Type, defaultValue interface{}) interface{} {
			return method(value, defaultValue)
		}
	}))

	return operators
}

func transform2TypeOperator(target Type, transform MethodTransform) *TypeOperator {
	return &TypeOperator{
		target: target,
		methodTransform: transform,
	}
}

func registerTransformer(source Type, operators []*TypeOperator) *Transformer {
	if len(operators) == 0 {
		return nil
	}

	transformer := Transformer {
		Source: source,
	}

	operateMap := make(map[Type]Operate)
	for _, operator := range operators {
		operateMap[operator.target] = operator.methodTransform()
	}

	transformer.TargetOperateMap = operateMap
	return &transformer
}

