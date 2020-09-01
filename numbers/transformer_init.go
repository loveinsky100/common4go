/**
 * @author leo
 * @date 2020/8/18 3:57 下午
 */
package numbers

func init()  {
	RegisterOperator("string", &TransformString{})
	RegisterOperator("uint8", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(uint8))
	}})

	RegisterOperator("uint16", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(uint16))
	}})

	RegisterOperator("uint32", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(uint32))
	}})

	RegisterOperator("uint64", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(uint64))
	}})

	RegisterOperator("int", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(int))
	}})

	RegisterOperator("uint", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(uint))
	}})

	RegisterOperator("int8", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(int8))
	}})

	RegisterOperator("int16", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(int16))
	}})

	RegisterOperator("int32", &TransformInt64{convert: func(value Value) int64 {
		return int64(value.(int32))
	}})

	RegisterOperator("int64", &TransformInt64{convert: func(value Value) int64 {
		return value.(int64)
	}})


	RegisterOperator("bool", &TransformBool{})
	RegisterOperator("float64", &TransformFloat64{})
}