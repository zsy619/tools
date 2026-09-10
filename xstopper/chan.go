package xstopper

import "sync"

// ChanStopper 是 Stopper 接口的实现，适用于在 select 块中循环监听的场景。
// 关闭 Chan 字段即可向 goroutine 发出停止信号。
type ChanStopper struct {
	sync.Mutex
	Chan   chan struct{} // 用于广播停止信号；Stop 会关闭它
	onDone func()        // 停止完成时触发的回调（OnDone 注册）
	done   bool          // 是否已完成停止（Finish 之后置为 true）
}

// NewChanStopper 创建一个带缓冲大小为 1 的 ChanStopper。
// 返回的 *ChanStopper 上调用任何方法都是安全的。
func NewChanStopper() *ChanStopper {
	return &ChanStopper{Chan: make(chan struct{}, 1)}
}

// Stop 关闭 Chan 以通知所有监听者停止。
// 在已停止的 ChanStopper 上重复调用会对已关闭的 channel 进行 close，从而 panic。
// nil 接收者下为 no-op。
func (s *ChanStopper) Stop() {
	if s == nil {
		return
	}
	s.Lock()
	defer s.Unlock()
	close(s.Chan)
}

// OnDone 注册停止完成时的回调函数 f。
// 后注册的回调会覆盖之前的回调。nil 接收者下为 no-op。
func (s *ChanStopper) OnDone(f func()) {
	if s == nil {
		return
	}
	s.Lock()
	defer s.Unlock()
	s.onDone = f
}

// Finish 触发已注册的 OnDone 回调，并将 done 置为 true。
// 通常由 stoppee 在确认自身已停止后调用。nil 接收者下为 no-op。
func (s *ChanStopper) Finish() {
	if s == nil {
		return
	}
	s.Lock()
	defer s.Unlock()
	if s.onDone != nil {
		s.onDone()
	}
	s.done = true
}

// Done 报告停止流程是否已完成（Finish 是否已被调用）。
// nil 接收者会因解引用 panic，因此调用方需保证 s 非 nil。
func (s *ChanStopper) Done() bool {
	return s.done
}
