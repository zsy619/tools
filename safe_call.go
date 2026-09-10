package tools

import (
	"errors"
	"fmt"
)

// ErrMissingCallback 当 SafeCall 收到 nil 回调函数时返回的错误。
//
// 该错误用于提示调用方传入的回调为空，避免静默成功造成歧义。
var ErrMissingCallback = errors.New("tools: 缺乏回调函数，请传入有效的 func() error")

// SafeCall 同步执行给定的回调函数 f，并在 f 抛出 panic 时自动恢复。
//
// 设计要点：
//  1. 仅在 f 为 nil 时返回 ErrMissingCallback；其它场景调用结果与直接 f() 等价。
//  2. 若 f 内部发生 panic，函数会捕获并将 panic 值包装为 error 返回，
//     从而避免 panic 传播到上层调用栈中断业务。
//  3. 该函数不会持有任何全局互斥锁，调用之间完全并发安全。
//
// 注意：原实现使用全局 sync.Mutex 序列化所有调用，在 xdatabase 等场景下
// 会把并发的数据库操作变成单点串行执行，已经不再适用。SafeCall 现在只
// 提供 "panic -> error" 的语义转换，如需互斥请使用 sync.Mutex 或
// sync.RWMutex 等显式原语。
//
// 典型用法：
//
//	if err := tools.SafeCall(func() error {
//	    return doSomethingDangerous()
//	}); err != nil {
//	    log.Println("safe call failed:", err)
//	}
func SafeCall(f func() error) (err error) {
	if f == nil {
		return ErrMissingCallback
	}
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = fmt.Errorf("tools: SafeCall 捕获到 panic: %w", v)
			default:
				err = fmt.Errorf("tools: SafeCall 捕获到 panic: %v", v)
			}
		}
	}()
	return f()
}
