package xiaomipush

// 发送给Android设备的Message对象
func NewAndroidMessage(title, description string) *Message {
	return &Message{
		Payload:     "",
		Title:       title,
		Description: description,
		PassThrough: 0,
		NotifyType:  -1, // default notify type
		TimeToLive:  0,
		TimeToSend:  0,
		NotifyID:    0,
		Extra:       make(map[string]string),
	}
}

// 打开当前app对应的Launcher Activity。
func (m *Message) SetLauncherActivity() *Message {
	m.Extra["notify_effect"] = "1"
	return m
}

// 打开当前app内的任意一个Activity。
func (m *Message) SetJumpActivity(value string) *Message {
	m.Extra["notify_effect"] = "2"
	m.Extra["intent_uri"] = value
	return m
}

// 打开网页
func (m *Message) SetJumpWebURL(value string) *Message {
	m.Extra["notify_effect"] = "3"
	m.Extra["web_uri"] = value
	return m
}

//
func (m *Message) SetPayload(payload string) *Message {
	m.Payload = payload
	return m
}

// 增加通知类别（Channel）是 Android O 引入的新功能
func (m *Message) SetChannel(channelId, channelName, channelDescription string) *Message {
	// https://dev.mi.com/console/doc/detail?pId=1163#_12
	m.Extra["channel_id"] = channelId
	m.Extra["channel_name"] = channelName
	m.Extra["channel_description"] = ""
	return m
}
