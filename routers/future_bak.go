/**
 * @author leo
 * @date 2020/8/26 3:57 下午
 */
package routers

// import (
// 	"context"
// 	"fmt"
// 	"errors"
// 	"time"
// )
//
// type FutureStatus int
//
// const (
// 	NEW 	FutureStatus = iota
// 	RUNNING	FutureStatus = iota
// 	FINISH	FutureStatus = iota
// 	CANCEL 	FutureStatus = iota
// 	PANIC 	FutureStatus = iota
// )
//
// type Future interface {
// 	// wait until task done
// 	Get() (interface{}, error)
//
// 	// wait until task done or timeout
// 	GetWithTimeout(timeout time.Duration) (interface{}, error)
//
// 	// cancel task
// 	Cancel()
//
// 	// check is done
// 	IsDone() bool
//
// 	// execute task
// 	Execute()
// }
//
// type CallableFuture struct {
// 	ctx 		  context.Context
// 	callable      Callable
// 	statusChannel chan FutureStatus
// 	returnChannel chan FutureStatus
// 	cancel 		  func()
// 	status 		  FutureStatus
// 	errorMessage  error
// 	result        interface{}
// }
//
// func NewCallableFuture(ctx context.Context, callable Callable) Future {
// 	var cancel func()
// 	if _, ok := ctx.Deadline(); !ok && ctx.Err() == nil {
// 		ctx, cancel = context.WithCancel(ctx)
// 	}
//
// 	statusChan := make(chan FutureStatus, 1)
// 	returnChan := make(chan FutureStatus, 1)
// 	return &CallableFuture {
// 		ctx: 		   ctx,
// 		callable:      callable,
// 		statusChannel: statusChan,
// 		returnChannel: returnChan,
// 		cancel: 	   cancel,
// 		status: 	   NEW,
// 		errorMessage:  nil,
// 		result:        nil,
// 	}
// }
//
// func(future *CallableFuture) Get() (interface{}, error) {
// 	if future.status > RUNNING {
// 		return future.result, future.errorMessage
// 	}
//
// 	select {
// 	case status := <-future.statusChannel:
// 		future.status = status
// 		return future.result, future.errorMessage
// 	}
// }
//
// func(future *CallableFuture) GetWithTimeout(timeout time.Duration) (interface{}, error) {
// 	if future.status > RUNNING {
// 		return future.result, future.errorMessage
// 	}
//
// 	timer := time.NewTimer(timeout)
// 	select {
// 	case <-timer.C:
// 		return nil, errors.New(fmt.Sprintf("future task timeout(%s)", timeout.String()))
// 	case status := <-future.statusChannel:
// 		future.status = status
// 		return future.result, future.errorMessage
// 	}
// }
//
// func(future *CallableFuture) Cancel() {
// 	if nil != future.cancel {
// 		future.cancel()
// 	}
// }
//
// func(future *CallableFuture) IsDone() bool {
// 	return future.status > RUNNING
// }
//
// func(future *CallableFuture) Execute() {
// 	future.status = RUNNING
// 	go future.innerExecute()
// 	select {
// 	case <-future.ctx.Done():
// 		future.statusChannel <- CANCEL
// 		close(future.statusChannel)
// 		return
// 	case status := <-future.returnChannel:
// 		future.statusChannel <- status
// 		close(future.statusChannel)
// 	}
// }
//
// func(future *CallableFuture) innerExecute() {
// 	defer func() {
// 		status := FINISH
// 		err := recover()
// 		if err != nil {
// 			status = PANIC
// 			future.errorMessage = errors.New(fmt.Sprintf("an error occured when execute task: %+v, error: %+v", future.callable, err))
// 		}
//
// 		future.returnChannel <- status
// 		close(future.returnChannel)
// 	}()
//
// 	if nil == future.callable {
// 		return
// 	}
//
// 	result, err := future.callable.Call()
// 	future.result = result
// 	future.errorMessage = err
// }