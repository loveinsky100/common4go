/**
 * @author leo
 * @date 2020/8/18 5:32 下午
 */
package collections

type Map interface {
	Put(key interface{}, value interface{})

	Remove(key interface{})

	Get(key interface{}) (interface{}, bool)

	Foreach(consumer func(key interface{}, value interface{}) bool)
}