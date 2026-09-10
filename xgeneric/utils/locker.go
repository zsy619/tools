package utils

import (
	gosync "sync"
)

// Locker 定义了抽象的锁接口，提供 Lock/Unlock 与读锁 RLock/RUnlock 方法。
// 编译期断言 *gosync.RWMutex 实现了该接口。
type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*gosync.RWMutex)(nil)

// FakeLocker 是一个空操作（无同步）的假锁实现。
// 用于在不需要并发安全时作为 Locker 接口的占位实现。
type FakeLocker struct{}

// Lock 是空操作，不执行任何锁定。
func (l FakeLocker) Lock() {}

// Unlock 是空操作，不执行任何解锁。
func (l FakeLocker) Unlock() {}

// RLock 是空操作，不执行任何读锁定。
func (l FakeLocker) RLock() {}

// RUnlock 是空操作，不执行任何读解锁。
func (l FakeLocker) RUnlock() {}
