/**
 * @author leo
 * @date 2020/8/18 4:17 下午
 */
package numbers

import (
	"fmt"
	"strconv"
)

type Int64Convert func(value Value) int64

type TransformInt64 struct {
	convert Int64Convert
}

func (ts *TransformInt64) Transform2String(value Value, defaultValue Value) interface{} {
	v := ts.convert(value)
	return strconv.FormatInt(v, 10)
}

func (ts *TransformInt64) Transform2Int64(value Value, defaultValue Value) interface{} {
	return ts.convert(value)
}

func (ts *TransformInt64) Transform2Bool(value Value, defaultValue Value) interface{} {
	v := ts.convert(value)
	if v >= 1 {
		return true
	}

	return false
}

func (ts *TransformInt64) Transform2Float64(value Value, defaultValue Value) interface{} {
	return float64(ts.convert(value))
}

const (
	TRUE = "true"
	FALSE = "false"
)

type TransformBool struct {
}

func (ts *TransformBool) Transform2String(value Value, defaultValue Value) interface{} {
	v := value.(bool)
	if v {
		return TRUE
	}

	return FALSE
}

func (ts *TransformBool) Transform2Int64(value Value, defaultValue Value) interface{} {
	if value.(bool) {
		return int64(1)
	}

	return int64(0)
}

func (ts *TransformBool) Transform2Bool(value Value, defaultValue Value) interface{} {
	return value
}

func (ts *TransformBool) Transform2Float64(value Value, defaultValue Value) interface{} {
	if value.(bool) {
		return float64(1)
	}

	return float64(0)
}

type TransformFloat64 struct {
}

func (ts *TransformFloat64) Transform2String(value Value, defaultValue Value) interface{} {
	// return strconv.FormatFloat(value.(float64), 'E', -1, 64)
	return fmt.Sprintf("%f", value.(float64))
}

func (ts *TransformFloat64) Transform2Int64(value Value, defaultValue Value) interface{} {
	return int64(value.(float64))
}

func (ts *TransformFloat64) Transform2Bool(value Value, defaultValue Value) interface{} {
	v := value.(float64)
	if v >= 1 {
		return true
	}

	return false
}

func (ts *TransformFloat64) Transform2Float64(value Value, defaultValue Value) interface{} {
	return value
}


