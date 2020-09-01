/**
 * @author leo
 * @date 2020/8/26 3:57 下午
 */
package routers

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type FutureStatus int

const (
	NEW 	FutureStatus = iota
	CANCEL 	FutureStatus = iota
	RUNNING	FutureStatus = iota
	FINISH	FutureStatus = iota
	PANIC 	FutureStatus = iota
)

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

type CallableResult struct {
	result        interface{}
	errorMessage  error
	status 		  FutureStatus
}

type CallableFuture struct {
	RW	  		  sync.RWMutex
	callable      Callable
	respChannel	  chan *CallableResult
	status 		  FutureStatus
	errorMessage  error
	result        interface{}
	startTime	  time.Time
	endTime		  time.Time
}

func NewCallableFuture(callable Callable) Future {
	statusChan := make(chan *CallableResult, 1)
	return &CallableFuture {
		callable:      callable,
		respChannel:   statusChan,
		status: 	   NEW,
		errorMessage:  nil,
		result:        nil,
		startTime: 	   time.Now(),
		endTime: 	   time.Now(),
	}
}

func (future *CallableFuture) Get() (interface{}, error) {
	return future.GetWithTimeout(0)
}

func (future *CallableFuture) GetWithTimeout(timeout time.Duration) (interface{}, error) {
	future.RW.RLock()
	if future.status == CANCEL || future.status > RUNNING {
		future.RW.RUnlock()
		return future.result, future.errorMessage
	}

	future.RW.RUnlock()

	if 0 == timeout {
		select {
		case resp := <-future.respChannel:
			future.innerDone(resp)
			return future.innerReadResult()
		}
	} else {
		distance := time.Until(future.startTime)
		currentTimeout := timeout + distance
		timer := time.NewTimer(currentTimeout)
		select {
		case resp := <-future.respChannel:
			future.innerDone(resp)
			return future.innerReadResult()
		case <-timer.C:
			return nil, errors.New(fmt.Sprintf("future task timeout(%s)", timeout.String()))
		}
	}
}

func (future *CallableFuture) Cancel() bool {
	future.RW.RLock()
	if future.status > CANCEL {
		future.RW.RUnlock()
		return false
	}

	future.RW.RUnlock()

	future.RW.Lock()
	defer future.RW.Unlock()

	// double check, when get write lock, status maybe changed
	if future.status > CANCEL {
		return false
	}

	future.status = CANCEL
	future.errorMessage = errors.New("future task canceled")
	future.endTime = time.Now()
	close(future.respChannel)
	return true
}

func (future *CallableFuture) Status() FutureStatus {
	future.RW.RLock()
	defer future.RW.RUnlock()

	return future.status
}

func (future *CallableFuture) Execute() {
	// check if it is canceled or in other status
	future.RW.RLock()
	status := future.status
	if status > NEW {
		future.RW.RUnlock()
		return
	}

	future.RW.RUnlock()
	future.RW.Lock()
	if future.status > NEW {
		future.RW.Unlock()
		return
	}

	future.status = RUNNING
	future.RW.Unlock()

	future.innerExecute()
}

func (future *CallableFuture) innerExecute() () {
	callableResult := &CallableResult{}

	defer func() {
		err := recover()
		if err != nil {
			callableResult.status = PANIC
			callableResult.errorMessage = errors.New(fmt.Sprintf("an error occured when execute task: %+v, error: %+v", future.callable, err))
		}

		future.respChannel <- callableResult
		close(future.respChannel)
	}()

	if nil == future.callable {
		return
	}

	resp, err := future.callable.Call()
	callableResult.result = resp
	callableResult.errorMessage = err
	callableResult.status = FINISH
}

func (future *CallableFuture) innerDone(resp *CallableResult) {
	future.RW.Lock()
	defer future.RW.Unlock()

	future.status = resp.status
	future.errorMessage = resp.errorMessage
	future.result = resp.result
	future.endTime = time.Now()
}

func (future *CallableFuture) innerReadResult() (interface{}, error) {
	future.RW.RLock()
	defer future.RW.RUnlock()

	return future.result, future.errorMessage
}