package xcas

import "sync"

// SessionStore 用于存储会话与票据的映射关系。
// SessionID 通常从 Cookie 中读取。
type SessionStore interface {
	// Get 根据会话 ID 获取对应的票据；不存在时返回空字符串与 false。
	Get(sessionID string) (string, bool)

	// Set 将会话 ID 与票据绑定并保存。
	Set(sessionID, ticket string) error

	// Delete 删除指定会话 ID 的映射。
	Delete(sessionID string) error
}

// NewMemorySessionStore 创建一个使用内存存储的默认 SessionStore 实现。
func NewMemorySessionStore() SessionStore {
	return &memorySessionStore{
		sessions: make(map[string]string),
	}
}

// memorySessionStore 内存版 SessionStore 实现。
type memorySessionStore struct {
	mu       sync.RWMutex
	sessions map[string]string
}

// Get 根据 sessionID 获取对应的票据，不存在时第二个返回值为 false。
func (m *memorySessionStore) Get(sessionID string) (string, bool) {
	m.mu.RLock()
	ticket, ok := m.sessions[sessionID]
	m.mu.RUnlock()

	return ticket, ok
}

// Set 记录 sessionID 与 ticket 的映射。
func (m *memorySessionStore) Set(sessionID, ticket string) error {
	m.mu.Lock()
	m.sessions[sessionID] = ticket
	m.mu.Unlock()

	return nil
}

// Delete 删除指定 sessionID 的映射；sessionID 不存在时 delete 为 no-op。
func (m *memorySessionStore) Delete(sessionID string) error {
	m.mu.Lock()
	delete(m.sessions, sessionID)
	m.mu.Unlock()

	return nil
}
