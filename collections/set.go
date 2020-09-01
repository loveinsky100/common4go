/**
 * @author leo
 * @date 2020/8/18 5:54 下午
 */
package collections

type Set interface {
	
	Add(key interface{})

	Remove(key interface{})

	Contains(key interface{}) bool

	Foreach(consumer func(key interface{}) bool)
}