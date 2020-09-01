/**
 * @author leo
 * @date 2020/8/18 5:51 下午
 */
package collections

const EMPTY = 0

type _HashSet struct {
	hashMap Map
}

func NewHashSet() Set {
	return &_HashSet{
		hashMap: NewHashMap(),
	}
}

func NewSyncHashSet() Set {
	return &_HashSet{
		hashMap: NewSyncHashMap(),
	}
}

func (hashSet *_HashSet) Add(key interface{}) {
	hashSet.hashMap.Put(key, EMPTY)
}

func (hashSet *_HashSet) Remove(key interface{}) {
	hashSet.hashMap.Remove(key)
}

func (hashSet *_HashSet) Contains(key interface{}) bool {
	_, ok := hashSet.hashMap.Get(key)
	return ok
}

func (hashSet *_HashSet) Foreach(consumer func(key interface{}) bool) {
	hashSet.hashMap.Foreach(func(key interface{}, value interface{}) bool {
		return consumer(key)
	})
}