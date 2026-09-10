// build +integration

package xemail

import (
	"crypto/md5" // nolint: gosec
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
	"github.com/beeker1121/mailchimp-go/lists/members"

	"github.com/zsy619/tools/xjson"
)

// NOTE(PN): 此 Mailchimp 库主要用于将地址加入邮件列表；
// 实际的邮件投递请通过 Sendgrid 的 Emailer 完成。

// 实现 MailingListMemberManager 接口。

// NewMailchimpAPI 是一个便捷函数，使用指定的 apiKey 实例化一个新的 MailchimpAPI。
func NewMailchimpAPI(apiKey string) *MailchimpAPI {
	return &MailchimpAPI{
		apiKey: apiKey,
	}
}

// MailchimpTag 表示一个 Mailchimp 标签。
type MailchimpTag string

const (
	// ErrTitleResourceNotFound Mailchimp 返回的资源不存在错误标题。
	ErrTitleResourceNotFound = "resource not found"
	// ErrTitleMemberExists Mailchimp 返回的成员已存在错误标题。
	ErrTitleMemberExists = "member exists"
)

const (
	// ServiceNameMailchimp Mailchimp 在 ListMemberManager 中注册的服务名。
	ServiceNameMailchimp = "mailchimp"
)

// MailchimpAPI 是对 Mailchimp API 的封装，底层使用 beeker1121/mailchimp-go 库。
type MailchimpAPI struct {
	apiKey string
}

// ServiceName 返回底层邮件列表服务商名称。
func (m *MailchimpAPI) ServiceName() string {
	return ServiceNameMailchimp
}

// GetListMember 返回 Mailchimp 指定列表中的成员信息。
// 邮箱不存在或网络失败时返回错误；成功时填充 ServiceID/EmailAddress/Status/Tags 等字段。
func (m *MailchimpAPI) GetListMember(listID string, email string) (*ListMember, error) {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return nil, err
	}

	mcMember, err := m.getMember(listID, m.mailchimpUserHash(email), nil)
	if err != nil {
		return nil, err
	}

	var status Status

	switch mcMember.Status {
	case members.StatusSubscribed:
		status = StatusSubscribed
	case members.StatusUnsubscribed:
		status = StatusUnsubscribed
	case members.StatusCleaned:
		status = StatusInvalid
	default:
		status = StatusPending
	}

	var tags []Tag
	if len(mcMember.Tags) > 0 {
		tags = make([]Tag, len(mcMember.Tags))
		for ind, t := range mcMember.Tags {
			tags[ind] = Tag(t.Name)
		}
	}

	return &ListMember{
		ServiceID:       mcMember.ID,
		EmailAddress:    mcMember.EmailAddress,
		Status:          status,
		SignupDate:      mcMember.TimestampSignup,
		LastUpdatedDate: mcMember.LastChanged,
		Tags:            tags,
	}, nil
}

// IsSubscribedToList 判断指定邮箱是否在指定列表中处于已订阅状态。
// 当成员不存在或状态非 subscribed 时返回 (false, nil)。
func (m *MailchimpAPI) IsSubscribedToList(listID string, email string) (bool, error) {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return false, err
	}

	member, err := m.getMember(listID, m.mailchimpUserHash(email), nil)
	if err != nil {
		if m.isMemberNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	if member == nil {
		return false, nil
	}

	if member.Status != members.StatusSubscribed {
		return false, nil
	}

	return true, nil
}

// SubscribeToList 将指定邮箱加入指定列表。
// 注意：仅在首次创建时附加 params.Tags，重新订阅不会再次写入标签。
// 若成员已存在则会将其状态更新为已订阅。
func (m *MailchimpAPI) SubscribeToList(listID string, email string, params *SubscriptionParams) error {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	newParams := &NewMemberParams{
		EmailAddress:    strings.ToLower(email),
		Status:          members.StatusSubscribed,
		TimestampSignup: now,
		TimestampOpt:    now,
	}

	var mcTags []MailchimpTag

	if params != nil {
		if len(params.Tags) > 0 {
			mcTags = make([]MailchimpTag, len(params.Tags))
			for ind, val := range params.Tags {
				mcTags[ind] = MailchimpTag(val)
			}

			newParams.Tags = mcTags
		}
	}

	member, err := m.newMember(listID, newParams)
	if err != nil {
		// 若成员已在列表中，则更新其状态为已订阅。
		if !m.isMemberAlreadyOnListError(err) {
			return err
		}

		updateParams := &members.UpdateParams{
			Status: members.StatusSubscribed,
		}
		member, err := members.Update(listID, m.mailchimpUserHash(email), updateParams)
		if err != nil {
			return err
		}

		if member == nil {
			return fmt.Errorf("No member returns on update")
		}
		return nil
	}

	if member == nil {
		return fmt.Errorf("No member returns on subscribe")
	}

	return nil
}

// UnsubscribeFromList 将指定邮箱从指定列表中退订。
// 当 delete 为 true 时会从列表中永久删除该成员；否则仅更新状态为 unsubscribed。
func (m *MailchimpAPI) UnsubscribeFromList(listID string, email string, delete bool) error {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return err
	}

	if delete {
		return members.Delete(listID, m.mailchimpUserHash(email))
	}

	params := &members.UpdateParams{
		Status: members.StatusUnsubscribed,
	}

	_, err = members.Update(listID, m.mailchimpUserHash(email), params)
	return err
}

// mailchimpUserHash 计算邮箱小写形式的 MD5 哈希，作为 Mailchimp 用户标识。
func (m *MailchimpAPI) mailchimpUserHash(emailAddress string) string {
	// 邮箱地址小写形式的 md5 哈希
	h := md5.New()                                        // nolint: gosec
	_, _ = h.Write([]byte(strings.ToLower(emailAddress))) // nolint: gosec, errcheck
	return hex.EncodeToString(h.Sum(nil))
}

// isMemberNotFoundError 判断错误是否为 Mailchimp 的“资源不存在”错误。
func (m *MailchimpAPI) isMemberNotFoundError(err error) bool {
	e, ok := err.(*mailchimp.APIError)
	if ok {
		if e.Status == 404 && strings.ToLower(e.Title) == ErrTitleResourceNotFound {
			return true
		}
	}
	return false
}

// isMemberAlreadyOnListError 判断错误是否为 Mailchimp 的“成员已存在”错误。
func (m *MailchimpAPI) isMemberAlreadyOnListError(err error) bool {
	e, ok := err.(*mailchimp.APIError)
	if ok {
		if e.Status == 400 && strings.ToLower(e.Title) == ErrTitleMemberExists {
			return true
		}
	}
	return false
}

// NOTE(PN): 以下内容扩展了 mailchimp-go 库，加入了对 Tags 的支持。
// 待有空时将尝试向上游提交 PR 合入这些改动。

// NewMemberParams 是增加了 Tags 字段的成员创建参数版本。
type NewMemberParams struct {
	EmailType       members.EmailType      `json:"email_type,omitempty"`
	Status          members.Status         `json:"status"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	Location        *members.Location      `json:"location,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup time.Time              `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    time.Time              `json:"timestamp_opt,omitempty"`
	EmailAddress    string                 `json:"email_address"`
	Tags            []MailchimpTag         `json:"tags,omitempty"`
}

// MarshalJSON 自定义 NewMemberParams 的 JSON 序列化，
// 用于修正 Mailchimp 期望的时间戳格式。
func (np *NewMemberParams) MarshalJSON() ([]byte, error) {
	var timestampSignup string
	var timestampOpt string

	if !np.TimestampSignup.IsZero() {
		timestampSignup = np.TimestampSignup.Format(time.RFC3339)
	}
	if !np.TimestampOpt.IsZero() {
		timestampOpt = np.TimestampOpt.Format(time.RFC3339)
	}

	type alias NewMemberParams
	return xjson.Marshal(&struct {
		*alias
		TimestampSignup string `json:"timestamp_signup,omitempty"`
		TimestampOpt    string `json:"timestamp_opt,omitempty"`
	}{
		alias:           (*alias)(np),
		TimestampSignup: timestampSignup,
		TimestampOpt:    timestampOpt,
	})
}

// MemberTag 表示 Member 上的一个标签。
type MemberTag struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Member 表示邮件列表中增加了 Tags 字段的一个成员。
type Member struct {
	ID              string                 `json:"id"`
	EmailAddress    string                 `json:"email_address"`
	UniqueEmailID   string                 `json:"unique_email_id"`
	EmailType       members.EmailType      `json:"email_type,omitempty"`
	Status          members.Status         `json:"status"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Stats           *members.Stats         `json:"stats,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup time.Time              `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    time.Time              `json:"timestamp_opt,omitempty"`
	MemberRating    uint8                  `json:"member_rating,omitempty"`
	LastChanged     time.Time              `json:"last_changed,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	EmailClient     string                 `json:"email_client,omitempty"`
	Location        *members.Location      `json:"location,omitempty"`
	LastNote        *members.Note          `json:"last_note,omitempty"`
	ListID          string                 `json:"list_id"`
	Tags            []*MemberTag           `json:"tags,omitempty"`
}

// newMember 向指定 list 添加一个新成员，支持 tags。
func (m *MailchimpAPI) newMember(listID string, params *NewMemberParams) (*Member, error) {
	res := &Member{}
	path := fmt.Sprintf("lists/%s/members", listID)

	if params == nil {
		if err := mailchimp.Call("POST", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("POST", path, nil, params, res); err != nil {
		return nil, err
	}
	return res, nil
}

// getMember 获取指定 list 中指定哈希成员的信息，支持 tags。
func (m *MailchimpAPI) getMember(listID, hash string, params *members.GetMemberParams) (*Member, error) {
	res := &Member{}
	path := fmt.Sprintf("lists/%s/members/%s", listID, hash)

	if params == nil {
		if err := mailchimp.Call("GET", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("GET", path, params, nil, res); err != nil {
		return nil, err
	}
	return res, nil
}
