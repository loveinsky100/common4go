/**
 * @author leo
 * @date 2020/8/28 3:56 下午
 */
package gotest

import (
	"../numbers"
	"../routers"
	"runtime"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	pool := routers.NewGoRouterPool(10, &routers.CallableRejectedHandler {
		Handler: func(callable routers.Callable) {
		},
	})

	future, err := pool.AddHandler(func() (interface{}, error) {
		time.Sleep(time.Millisecond * 100)
		return 1, nil
	})

	if err != nil {
		t.Errorf("an error occured: %+v", err)
	}

	v, err := future.GetWithTimeout(time.Millisecond * 200)
	if err != nil {
		t.Errorf("an error occured: %+v", err)
	}

	if 1 != v {
		t.Errorf("value not equal")
	}
}

func TestCancel(t *testing.T) {
	pool := routers.NewGoRouterPool(10, &routers.CallableRejectedHandler {
		Handler: func(callable routers.Callable) {
		},
	})

	for index := 0; index < 10; index ++ {
		future, err := pool.AddHandler(func() (interface{}, error) {
			time.Sleep(time.Millisecond * 100)
			return 1, nil
		})

		if err != nil {
			t.Errorf("an error occured: %+v", err)
		}

		if index % 2 == 0 {
			time.Sleep(time.Millisecond * 10)
		}

		cancelSuccess := future.Cancel()
		_, err = future.GetWithTimeout(time.Millisecond * 200)
		if cancelSuccess {
			if index % 2 == 0 {
				t.Errorf("an error occured: cancel failed %s", numbers.ToString(cancelSuccess))
				break
			}

			if err == nil {
				t.Errorf("an error occured: cancel failed %s", numbers.ToString(cancelSuccess))
			}
		} else {
			if err != nil {
				t.Errorf("an error occured: cancel failed %s", numbers.ToString(cancelSuccess))
			}
		}

	}
}

func TestReject(t *testing.T) {
	pool := routers.NewGoRouterPool(1, &routers.CallableRejectedHandler {
		Handler: func(callable routers.Callable) {

		},
	})

	_, err := pool.AddHandler(func() (interface{}, error) {
		time.Sleep(time.Millisecond * 100)
		return 1, nil
	})

	if err != nil {
		t.Errorf("an error occured: %+v", err)
	}

	_, err = pool.AddHandler(func() (interface{}, error) {
		time.Sleep(time.Millisecond * 100)
		return 1, nil
	})

	if err == nil {
		t.Errorf("an error occured: reject failed")
	}
}

func TestTimeout(t *testing.T) {
	pool := routers.NewGoRouterPool(1, &routers.CallableRejectedHandler {
		Handler: func(callable routers.Callable) {

		},
	})

	future, err := pool.AddHandler(func() (interface{}, error) {
		time.Sleep(time.Millisecond * 100)
		return 1, nil
	})

	if err != nil {
		t.Errorf("an error occured: %+v", err)
	}

	_, err = future.GetWithTimeout(time.Millisecond * 50)
	if err == nil {
		t.Errorf("an error occured: timeout failed")
	}
}

func TestIsDone(t *testing.T) {
	pool := routers.NewGoRouterPool(1, &routers.CallableRejectedHandler {
		Handler: func(callable routers.Callable) {

		},
	})

	future, err := pool.AddHandler(func() (interface{}, error) {
		time.Sleep(time.Millisecond * 100)
		return 1, nil
	})

	if err != nil {
		t.Errorf("an error occured: %+v", err)
	}

	_, err = future.GetWithTimeout(time.Millisecond * 200)
	if err != nil {
		t.Errorf("an error occured: %+v", err)
	}

	if future.Status() != routers.FINISH {
		t.Errorf("an error occured: is done error")
	}
}

func TestGoRouterCounter(t *testing.T)  {
	startCount := runtime.NumGoroutine()
	pool := routers.NewGoRouterPool(10, &routers.CallableRejectedHandler {
		Handler: func(callable routers.Callable) {

		},
	})

	for index := 0; index < 15; index ++ {
		pool.AddHandler(func() (interface{}, error) {
			time.Sleep(time.Millisecond * 1000)
			return 1, nil
		})

		// sleep make sure router add into runtime
		time.Sleep(time.Millisecond * 10)
		addCount := runtime.NumGoroutine() - startCount
		if index < 10 {
			if addCount != (1 + index) {
				t.Errorf("an error occured: add count error: %d", runtime.NumGoroutine())
			}
		} else {
			if addCount != 10 {
				t.Errorf("an error occured: add count error %d", runtime.NumGoroutine())
			}
		}
	}
}