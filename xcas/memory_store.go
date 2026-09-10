package xcas

import (
	"sync"
)

// MemoryStore 在内存中实现 TicketStore 接口，存储票据对应的认证响应。
// 全部操作使用读写锁保护，适合单机场景；进程重启后数据会丢失。
type MemoryStore struct {
	mu    sync.RWMutex
	store map[string]*AuthenticationResponse
}

// Read 根据票据 ID 返回对应的 AuthenticationResponse；
// 当 store 未初始化或 ID 不存在时返回 ErrInvalidTicket。
func (s *MemoryStore) Read(id string) (*AuthenticationResponse, error) {
	s.mu.RLock()

	if s.store == nil {
		s.mu.RUnlock()
		return nil, ErrInvalidTicket
	}

	t, ok := s.store[id]
	s.mu.RUnlock()

	if !ok {
		return nil, ErrInvalidTicket
	}

	return t, nil
}

// Write 将 AuthenticationResponse 与票据 ID 关联写入存储；store 未初始化时会懒加载。
func (s *MemoryStore) Write(id string, ticket *AuthenticationResponse) error {
	s.mu.Lock()

	if s.store == nil {
		s.store = make(map[string]*AuthenticationResponse)
	}

	s.store[id] = ticket

	s.mu.Unlock()
	return nil
}

// Delete 删除指定票据；ID 不存在时 delete 是 no-op，函数始终返回 nil。
func (s *MemoryStore) Delete(id string) error {
	s.mu.Lock()
	delete(s.store, id)
	s.mu.Unlock()
	return nil
}

// Clear 清空所有票据数据：将内部 map 置为 nil，后续写入时会重新创建。
func (s *MemoryStore) Clear() error {
	s.mu.Lock()
	s.store = nil
	s.mu.Unlock()
	return nil
}
