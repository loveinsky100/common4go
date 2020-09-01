/**
 * @author leo
 * @date 2020/8/18 6:21 下午
 */
package gotest

import (
	"encoding/json"
	"testing"
)

import "../collections"

func TestListAddGet(t *testing.T)  {
	list := collections.NewArrayList()
	list.Add(1)

	v, _ := list.Get(0)
	if v != 1 {
		t.Errorf("an error occured: v != 1")
	}
}

func TestListGetOutOfIndex(t *testing.T)  {
	list := collections.NewArrayList()
	list.Add(1)

	_, ok := list.Get(1)
	if ok {
		t.Errorf("an error occured: OutOfIndex")
	}

	_, ok = list.Get(-1)
	if ok {
		t.Errorf("an error occured: OutOfIndex")
	}
}

func TestListRemove(t *testing.T)  {
	list := collections.NewArrayList()
	list.Add(1)

	list.Get(1)
	if 1 != list.Size() {
		t.Errorf("an error occured: Remove")
	}

	list.Remove(0)
	if 0 != list.Size() {
		t.Errorf("an error occured: Remove")
	}
}

func TestForeach(t *testing.T)  {
	list := collections.NewArrayList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	max := 0
	list.Foreach(func(value interface{}, index int) bool {
		if value != index + 1 {
			t.Errorf("an error occured: Foreach")
		}

		max = index
		return true
	})

	if max != 3 {
		t.Errorf("an error occured: Foreach")
	}
}

func TestListJson(t *testing.T)  {
	list := collections.NewArrayList()
	list.Add(1)

	list2 := collections.NewArrayList()
	list.Add(list2)
	list2.Add(2)

	data, err := json.Marshal(list)
	if nil != err {
		t.Errorf("an error occured: Json %+v", err)
	}

	aList := collections.NewArrayList()
	json.Unmarshal(data, &aList)

	v, ok := aList.Get(0)
	if !ok || int(v.(float64)) != 1 {
		t.Errorf("an error occured: Json")
	}
}