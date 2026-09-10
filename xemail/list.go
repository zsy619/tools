package xemail

import "time"

// Status 表示邮件列表成员在服务商中的订阅状态。
type Status string

const (
	// StatusSubscribed 成员已订阅。
	StatusSubscribed Status = "subscribed"

	// StatusUnsubscribed 成员已退订。
	StatusUnsubscribed = "unsubscribed"

	// StatusInvalid 成员状态无效（例如被清理）。
	StatusInvalid = "invalid"

	// StatusPending 成员处于待确认状态，等待 opt-in。
	StatusPending = "pending"
)

// Tag 表示邮件列表成员上的一个标签。
type Tag string

// ListMember 表示邮件列表中一个成员的基础数据。
type ListMember struct {
	ServiceID       string    // 服务商侧的成员 ID
	EmailAddress    string    // 邮箱地址
	Status          Status    // 当前订阅状态
	SignupDate      time.Time // 注册时间
	LastUpdatedDate time.Time // 最近更新时间
	Tags            []Tag     // 关联标签
}

// SubscriptionParams 描述订阅时可附带的其他参数。
type SubscriptionParams struct {
	Tags []Tag // 订阅时要附带的标签
}

// ListMemberManager 定义邮件列表成员管理服务需要实现的方法集合。
type ListMemberManager interface {
	// ServiceName 返回底层邮件列表服务商名称（如 sendgrid、mailchimp）。
	ServiceName() string
	// GetListMember 返回指定 list 中的指定邮箱成员。
	GetListMember(listID string, email string) (*ListMember, error)
	// IsSubscribedToList 判断邮箱是否已订阅指定 list。
	IsSubscribedToList(listID string, email string) (bool, error)
	// SubscribeToList 将邮箱订阅到指定 list，并附带 params。
	SubscribeToList(listID string, email string, params *SubscriptionParams) error
	// UnsubscribeFromList 将邮箱从指定 list 退订；若 delete 为 true 则永久删除。
	UnsubscribeFromList(listID string, email string, delete bool) error
}
