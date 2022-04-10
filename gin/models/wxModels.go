package models

import "github.com/silenceper/wechat/v2/miniprogram/auth"

type WxUserSession struct {
	auth.ResCode2Session
}

type WxUserProfile struct {
	UserInfo      WxUser `json:"userInfo"`
	RawData       string `json:"rawData,omitempty"`
	Signature     string `json:"signature,omitempty"`
	EncryptedData string `json:"encryptedData,omitempty"`
	Iv            string `json:"iv,omitempty"`
	CloudID       string `json:"cloudID,omitempty"`
}

type WxPutProfile struct {
	UserProfile WxUserProfile `json:"userprofile"`
}
