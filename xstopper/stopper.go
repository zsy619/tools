package xstopper

// Stopper 是可嵌入接口，用于拥有后台 goroutine 且需要被所属对象所有者停止的对象（称为 stoppee）。
type Stopper interface {
	// Stop 请求 stoppee 停止其 goroutine。
	Stop()
	// OnDone 注册一个回调，在停止完成后触发。
	OnDone(func())
	// Finish 由 stoppee 在自身关闭后调用。
	Finish()
	// Done 仅在 Finish 被调用之后返回 true。
	Done() bool
}
