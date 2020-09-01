/**
 * @author leo
 * @date 2020/8/18 5:11 下午
 */
package collections

import (
	"encoding/json"
	"sync"
)

type _ArrayList struct {
	items []interface{}
}

func NewArrayList() List {
	return &_ArrayList{
		items: make([]interface{}, 0),
	}
}

func (arrayList *_ArrayList) Add(value interface{}) {
	arrayList.items = append(arrayList.items, value)
}

func (arrayList *_ArrayList) Get(index int) (interface{}, bool) {
	if index < 0 || index >= len(arrayList.items) {
		return nil, false
	}

	return arrayList.items[index], true
}

func (arrayList *_ArrayList) Remove(index int) bool {
	if index < 0 || index >= len(arrayList.items) {
		return false
	}

	arrayList.items = append(arrayList.items[:index], arrayList.items[index+1:]...)
	return true
}

func (arrayList *_ArrayList) Foreach(consumer func(value interface{}, index int) bool) {
	for i, v := range arrayList.items {
		r := consumer(v, i)
		if !r {
			break
		}
	}
}

func (arrayList *_ArrayList) Size() int {
	return len(arrayList.items)
}

func (arrayList *_ArrayList) MarshalJSON() ([]byte, error) {
	return json.Marshal(arrayList.items)
}

func (arrayList *_ArrayList) UnmarshalJSON(data []byte) error {
	var item []interface{}
	err := json.Unmarshal(data, &item)
	if nil != err {
		return err
	}

	arrayList.items = item
	return nil
}

type _SyncArrayList struct {
	ArrayList List
	rw sync.RWMutex
}

func NewSyncArrayList() List {
	return &_SyncArrayList {
		ArrayList: NewArrayList(),
	}
}

func NewSyncList(list List) List {
	return &_SyncArrayList {
		ArrayList: list,
	}
}

func (arrayList *_SyncArrayList) Add(value interface{}) {
	arrayList.rw.Lock()
	defer arrayList.rw.Unlock()

	arrayList.ArrayList.Add(value)
}

func (arrayList *_SyncArrayList) Get(index int) (interface{}, bool) {
	arrayList.rw.RLock()
	defer arrayList.rw.RUnlock()

	return arrayList.ArrayList.Get(index)
}

func (arrayList *_SyncArrayList) Remove(index int) bool {
	arrayList.rw.Lock()
	defer arrayList.rw.Unlock()

	return arrayList.ArrayList.Remove(index)
}

func (arrayList *_SyncArrayList) Foreach(consumer func(value interface{}, index int) bool) {
	arrayList.rw.RLock()
	defer arrayList.rw.RUnlock()

	arrayList.ArrayList.Foreach(consumer)
}

func (arrayList *_SyncArrayList) Size() int {
	arrayList.rw.RLock()
	defer arrayList.rw.RUnlock()
	return arrayList.ArrayList.Size()
}


func (arrayList *_SyncArrayList) MarshalJSON() ([]byte, error) {
	arrayList.rw.RLock()
	defer arrayList.rw.RUnlock()
	return json.Marshal(arrayList.ArrayList)
}

func (arrayList *_SyncArrayList) UnmarshalJSON(data []byte) error {
	arrayList.rw.Lock()
	defer arrayList.rw.Unlock()
	return json.Unmarshal(data, &arrayList.ArrayList)
}