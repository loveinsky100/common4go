/**
 * @author leo
 * @date 2020/8/18 5:24 下午
 */
package collections

type List interface {
	// add item to list
	Add(value interface{})

	// get index item
	Get(index int) (interface{}, bool)

	// remove item
	Remove(index int) bool

	// loop item
	Foreach(consumer func(value interface{}, index int) bool)

	// get item len
	Size() int
}