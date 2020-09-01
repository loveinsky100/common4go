/**
 * @author leo
 * @date 2020/8/18 4:40 下午
 */
package numbers

import "math"

const EMTYP_STRING = ""

func To(value interface{}, targetType string, defaultValue interface{}) interface{} {
	return Context.transform(value, Type(targetType), defaultValue)
}

func ToInt64(value interface{}) int64 {
	return ToInt64OrDefault(value, 0)
}

func ToInt64OrDefault(value interface{}, defaultValue int64) int64 {
	return To(value, "int64", defaultValue).(int64)
}

func ToInt32(value interface{}) int32 {
	return ToInt32OrDefault(value, 0)
}

func ToInt32OrDefault(value interface{}, defaultValue int32) int32 {
	v := ToInt64OrDefault(value, int64(defaultValue))
	if math.MaxInt32 < v || v < math.MinInt32 {
		return defaultValue
	}

	return int32(v)
}

func ToInt(value interface{}) int {
	return ToIntOrDefault(value, 0)
}

func ToIntOrDefault(value interface{}, defaultValue int) int {
	v := ToInt64OrDefault(value, int64(defaultValue))
	if math.MaxInt32 < v || v < math.MinInt32 {
		return defaultValue
	}

	return int(v)
}

func ToFloat64(value interface{}) float64 {
	return ToFloat64OrDefault(value, 0)
}
func ToFloat64OrDefault(value interface{}, defaultValue float64) float64 {
	return To(value, "float64", defaultValue).(float64)
}

func ToFloat32(value interface{}) float32 {
	return ToFloat32OrDefault(value, 0)
}

func ToFloat32OrDefault(value interface{}, defaultValue float32) float32 {
	v := ToFloat64OrDefault(value, float64(defaultValue))
	if v > math.MaxFloat32 || v < math.SmallestNonzeroFloat32 {
		return defaultValue
	}

	return float32(v)
}

func ToString(value interface{}) string {
	return ToStringOrDefault(value, EMTYP_STRING)
}

func ToStringOrDefault(value interface{}, defaultValue string) string {
	return To(value, "string", defaultValue).(string)
}

func ToBool(value interface{}) bool {
	return ToBoolOrDefault(value, false)
}

func ToBoolOrDefault(value interface{}, defaultValue bool) bool {
	return To(value, "bool", defaultValue).(bool)
}