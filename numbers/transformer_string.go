/**
 * @author leo
 * @date 2020/8/18 4:17 下午
 */
package numbers

import (
	"strconv"
	"strings"
)

var boolMap = map[string]bool{
	"1" : true,
	"true" : true,
	"yes" : true,
	"y" : true,
	"0" : false,
	"false" : false,
	"no" : false,
	"n" : false,
}

type TransformString struct {
}

func (ts *TransformString) Transform2String(value Value, defaultValue Value) interface{} {
	return value
}

func (ts *TransformString) Transform2Int64(value Value, defaultValue Value) interface{} {
	v, err := strconv.ParseInt(value.(string), 10, 64)
	if err != nil {
		return defaultValue.(int64)
	}

	return v
}

func (ts *TransformString) Transform2Bool(value Value, defaultValue Value) interface{} {
	str := value.(string)
	str = strings.ToLower(str)
	r, ok := boolMap[str]
	if !ok {
		return defaultValue
	}

	return r
}

func (ts *TransformString) Transform2Float64(value Value, defaultValue Value) interface{} {
	str := value.(string)
	v, err := strconv.ParseFloat(str, 64)
	if nil != err {
		return defaultValue
	}

	return v
}

