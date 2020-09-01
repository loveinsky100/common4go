/**
 * @author leo
 * @date 2020/8/18 6:21 下午
 */
package gotest

import (
	"testing"
)

import "../numbers"

func Asset(expression func()bool, t *testing.T) {
	v := expression()
	if !v {
		t.Errorf("an error occured")
	}
}

func TestTransformer(t *testing.T)  {
	Asset(func() bool {
		return numbers.ToInt64(int8(120)) == 120
	}, t)

	Asset(func() bool {
		return numbers.ToInt(int8(120)) == 120
	}, t)

	Asset(func() bool {
		return numbers.ToInt64("120") == 120
	}, t)

	Asset(func() bool {
		return numbers.ToInt64(120) == 120
	}, t)

	Asset(func() bool {
		return numbers.ToInt64(-999) == -999
	}, t)

	Asset(func() bool {
		return !numbers.ToBool(-999)
	}, t)

	Asset(func() bool {
		return numbers.ToBool("TRUE")
	}, t)

	Asset(func() bool {
		return !numbers.ToBool("TRUEE")
	}, t)

	Asset(func() bool {
		return !numbers.ToBool("FALSE")
	}, t)

	Asset(func() bool {
		return numbers.ToBool("1")
	}, t)

	Asset(func() bool {
		return numbers.ToBool(1)
	}, t)

	Asset(func() bool {
		return numbers.ToString(1) == "1"
	}, t)

	Asset(func() bool {
		return numbers.ToString(2.198900) == "2.198900"
	}, t)

	Asset(func() bool {
		return numbers.ToFloat64("2.198900") == 2.198900
	}, t)
}