/**
 * @author leo
 * @date 2020/8/18 5:32 下午
 */
package collections

import "sync"

type _HashMap struct {
	itemMap map[interface{}]interface{}
}

func NewHashMap() Map {
	return &_HashMap{
		itemMap: make(map[interface{}]interface{}),
	}
}

func (hashMap *_HashMap) Put(key interface{}, value interface{}) {
	hashMap.itemMap[key] = value
}

func (hashMap *_HashMap) Remove(key interface{}) {
	delete(hashMap.itemMap, key)
}

func (hashMap *_HashMap) Get(key interface{}) (interface{}, bool) {
	v, ok := hashMap.itemMap[key]
	return v, ok
}

func (hashMap *_HashMap) Foreach(consumer func(key interface{}, value interface{}) bool) {
	for k, v := range hashMap.itemMap {
		r := consumer(k, v)
		if !r {
			break
		}
	}
}

type _SyncHashMap struct {
	itemMap sync.Map
}

func NewSyncHashMap() Map {
	return &_SyncHashMap {
		itemMap: sync.Map{},
	}
}

func (hashMap *_SyncHashMap) Put(key interface{}, value interface{}) {
	hashMap.itemMap.Store(key, value)
}

func (hashMap *_SyncHashMap) Remove(key interface{}) {
	hashMap.itemMap.Delete(key)
}

func (hashMap *_SyncHashMap) Get(key interface{}) (interface{}, bool) {
	return hashMap.itemMap.Load(key)
}

func (hashMap *_SyncHashMap) Foreach(consumer func(key interface{}, value interface{}) bool) {
	hashMap.itemMap.Range(consumer)
}