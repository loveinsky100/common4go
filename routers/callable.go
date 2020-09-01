/**
 * @author leo
 * @date 2020/8/27 8:52 下午
 */
package routers

import "errors"

type Callable interface {
	// call method and return value or error
	Call() (interface{}, error)
}

type HandlerCallable struct {
	Handler func() (interface{}, error)
}

func (callable *HandlerCallable) Call() (interface{}, error) {
	if nil != callable.Handler {
		return callable.Handler()
	}

	return nil, errors.New("HandlerCallable handler is nil")
}