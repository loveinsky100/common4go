# common4go
Common Go lang data structures, utilities, Future Task etc  
If you are java programmer, It will help you.   

List, HashMap, HashSet  
gorouter pool(thread pool), Future Task  
basic number conversion  
common logs(in dev)

# Installation
With Go and git installed:

```
go get github.com/loveinsky100/common4go
````

# Data Structures
## List

1.ArrayList  
2.SyncArrayList  

```golang
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
``` 

```golang
	// create array list
	list := collections.NewArrayList()

	// create sync array list
	list := collections.NewSyncArrayList()

	// create sync list base on array list
	list := collections.NewSyncList(collections.NewArrayList())
```

## Map
1.HashMap  
2.SyncHashMap

```golang
type Map interface {
	Put(key interface{}, value interface{})

	Remove(key interface{})

	Get(key interface{}) (interface{}, bool)

	Foreach(consumer func(key interface{}, value interface{}) bool)
}
```

```golang
	// create hash map
	hashMap := collections.NewHashMap()

	// create sync hash map
	hashMap := collections.NewSyncHashMap()
```

## Set
1.HashSet

```golang
type Set interface {
	
	Add(key interface{})

	Remove(key interface{})

	Contains(key interface{}) bool

	Foreach(consumer func(key interface{}) bool)
}
```

```golang
	// create hash set
	hashSet := collections.NewHashSet()

	// create sync hash set
	hashSet := collections.NewSyncHashSet()
```
# Future Task
## Future
you can you use future task like java, and also you need use future task in a go router pool, such like thread pool in java, but go router pool is diffrent between thread pool, it only provide max pool size and reject handler, there are no min pool size or wait queue.

```golang
type GoRouterPool interface {
	Add(callable Callable) (Future, error)
	AddHandler(handler func() (interface{}, error)) (Future, error)
}

type Future interface {
	// wait until task done or panic
	Get() (interface{}, error)

	// wait until task done or timeout, notice: when task timeout it will still running in background until return method
	GetWithTimeout(timeout time.Duration) (interface{}, error)

	// cancel task, when task not enter, you can cancel it, true: cancel success false:
	Cancel() bool

	// check is done or in other status
	Status() FutureStatus

	// execute task, you need not call is when you put it in routers pool
	Execute()
}
```

how to use 
```golang
	pool := routers.NewGoRouterPool(10, &routers.CallableRejectedHandler {
		Handler: func(callable routers.Callable) {
			// do some thing
		},
	})

	future, err := pool.AddHandler(func() (interface{}, error) {
		time.Sleep(time.Millisecond * 100)
		return 1, nil
	})

	if err != nil {
		// reject
		t.Errorf("an error occured: reject", err)
	}

	v, err := future.GetWithTimeout(time.Millisecond * 200)
	if err != nil {
		t.Errorf("an error occured: %+v", err)
	}
```

# number conversion
```golang
numbers.ToInt64(int8(120))
numbers.ToFloat64("2.198900")
numbers.ToString(2.198900)
numbers.ToBool(1)
numbers.ToBool("FALSE")
```

# dev
### schedule task like EventLoop & EventLoopGroup
### linked list
### queue & stack & cache
### logs