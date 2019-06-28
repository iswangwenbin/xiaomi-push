package xiaomipush

import "strconv"

// 发送给IOS设备的Message对象
func NewIOSMessage(description string) *Message {
	return &Message{
		Payload:     "",
		Title:       "",
		Description: description,
		PassThrough: 0,
		NotifyType:  -1, // default notify type
		TimeToLive:  0,
		TimeToSend:  0,
		NotifyID:    0,
		Extra:       make(map[string]string),
	}
}

// 可选项，自定义通知数字角标。
func (m *Message) SetBadge(badge int64) *Message {
	m.Extra["badge"] = strconv.FormatInt(badge, 10)
	return m
}

// 可选项，iOS8推送消息快速回复类别。
func (m *Message) SetCategory(category string) *Message {
	m.Extra["category"] = category
	return m
}

// 可选项，自定义消息铃声。
func (m *Message) SetSoundURL(soundURL string) *Message {
	m.Extra["sound_url"] = soundURL
	return m
}
