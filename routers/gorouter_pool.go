/**
 * @author leo
 * @date 2020/8/27 8:52 下午
 */
package routers

import (
	"errors"
	"sync"
)

type GoRouterPool interface {
	Add(callable Callable) (Future, error)
	AddHandler(handler func() (interface{}, error)) (Future, error)
}

type DefaultGoRouterPool struct {
	mutex sync.Mutex
	poolSize int
	current int
	reject RejectedHandler
}

func NewGoRouterPool(poolSize int, reject RejectedHandler) GoRouterPool {
	pool := &DefaultGoRouterPool{
		poolSize: poolSize,
		reject: reject,
	}

	return pool
}

func (pool *DefaultGoRouterPool) Add(callable Callable) (Future, error) {
	pool.mutex.Lock()
	defer func() {
		pool.current++
		pool.mutex.Unlock()
	}()

	if pool.current >= pool.poolSize {
		if nil != pool.reject {
			pool.reject.Reject(callable)
		}

		return nil, errors.New("add pool failed due to pool full")
	}

	future := NewCallableFuture(callable)
	go pool.execute(future)
	return future, nil
}

func (pool *DefaultGoRouterPool) AddHandler(handler func() (interface{}, error)) (Future, error) {
	callable := &HandlerCallable {
		Handler: handler,
	}

	return pool.Add(callable)
}

func (pool *DefaultGoRouterPool) execute(future Future) {
	defer func() {
		pool.mutex.Lock()
		pool.current--
		pool.mutex.Unlock()
	}()

	future.Execute()
}