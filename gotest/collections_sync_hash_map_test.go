/**
 * @author leo
 * @date 2020/9/1 11:00 上午
 */
package gotest

import "testing"
import "../collections"

func TestSyncHashMapAddGet(t *testing.T)  {
	hashMap := collections.NewSyncHashMap()
	hashMap.Put("Hello", "World")

	v, _ := hashMap.Get("Hello")
	if v != "World" {
		t.Errorf("an error occured: v != World")
	}
}

func TestSyncHashMapRemove(t *testing.T)  {
	hashMap := collections.NewSyncHashMap()
	hashMap.Put("Hello", "World")

	v, _ := hashMap.Get("Hello")
	if v != "World" {
		t.Errorf("an error occured: v != World")
	}

	hashMap.Remove("Hello")
	_, ok := hashMap.Get("Hello")
	if ok {
		t.Errorf("an error occured")
	}
}

func TestSyncHashMapForeach(t *testing.T)  {
	hashMap := collections.NewSyncHashMap()
	hashMap.Put(0, "0")
	hashMap.Put(1, "1")
	hashMap.Put(2, "2")

	hashMap.Foreach(func(key interface{}, value interface{}) bool {
		if key == 0 && value != "0" {
			t.Errorf("an error occured: v != 0")
		}

		if key == 1 && value != "1" {
			t.Errorf("an error occured: v != 1")
		}

		if key == 2 && value != "2" {
			t.Errorf("an error occured: v != 2")
		}

		return true
	})
}
