/**
 * @author leo
 * @date 2020/9/1 11:00 上午
 */
package gotest

import "testing"
import "../collections"

func TestSyncHashSetContain(t *testing.T)  {
	hashSet := collections.NewSyncHashSet()
	hashSet.Add("Hello")

	ok := hashSet.Contains("Hello")
	if !ok {
		t.Errorf("an error occured")
	}
}

func TestSyncHashSetRemove(t *testing.T)  {
	hashSet := collections.NewSyncHashSet()
	hashSet.Add("Hello")

	ok := hashSet.Contains("Hello")
	if !ok {
		t.Errorf("an error occured")
	}

	hashSet.Remove("Hello")
	ok = hashSet.Contains("Hello")
	if ok {
		t.Errorf("an error occured")
	}
}

func TestSyncHashSetForeach(t *testing.T)  {
	hashSet := collections.NewSyncHashSet()
	hashSet.Add(0)
	hashSet.Add(1)
	hashSet.Add(2)
	hashSet.Add(3)

	total := 0 + 1 + 2 + 3
	count := 0
	hashSet.Foreach(func(key interface{}) bool {
		count += key.(int)
		return true
	})

	if count != total {
		t.Errorf("an error occured")
	}
}
