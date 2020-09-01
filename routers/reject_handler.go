/**
 * @author leo
 * @date 2020/8/27 9:43 下午
 */
package routers

type RejectedHandler interface {
	Reject(callable Callable)
}

type CallableRejectedHandler struct {
	Handler func(callable Callable)
}

func (rejectedHandler *CallableRejectedHandler) Reject(callable Callable) {
	if nil != rejectedHandler.Handler {
		rejectedHandler.Handler(callable)
	}
}

