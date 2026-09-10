package xcas

import (
	"errors"
)

// TicketStore 相关错误。
var (
	// ErrInvalidTicket 给定的票据未关联到任何 AuthenticationResponse。
	ErrInvalidTicket = errors.New("cas: ticket store: invalid ticket")
)

// TicketStore 提供对 service ticket 与 AuthenticationResponse 数据的存取接口。
type TicketStore interface {
	// Read 返回与票据 ID 关联的 AuthenticationResponse。
	Read(id string) (*AuthenticationResponse, error)

	// Write 存储一次票据校验得到的 AuthenticationResponse。
	Write(id string, ticket *AuthenticationResponse) error

	// Delete 删除与票据 ID 关联的 AuthenticationResponse。
	Delete(id string) error

	// Clear 清空存储中的全部 AuthenticationResponse 数据。
	Clear() error
}
