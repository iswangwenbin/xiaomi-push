package xiaomipush

import (
	"context"
	"encoding/json"
)

// 根据account，发送消息到指定account上
func (m *MiPush) SendToUserAccount(ctx context.Context, msg *Message, userAccount string) (*SendResult, error) {
	params := m.assembleSendToUserAccountParams(msg, userAccount)
	bytes, err := m.doPost(ctx, m.host+MessageUserAccountURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
