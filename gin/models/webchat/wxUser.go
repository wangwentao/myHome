package webchat

import "myHome/gin/models"

type WxUser struct {
	models.PGBaseModel
	OpenID    string `json:"openId,omitempty" gorm:"primarykey"`
	UnionID   string `json:"unionId,omitempty"`
	NickName  string `json:"nickName"`
	AvatarURL string `json:"avatarUrl"`
	Gender    int    `json:"gender"` //0 未知,1 男性,2 女性
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Language  string `json:"language"` //en 英文, zh_CN 简体中文, zh_TW 繁体中文,
}

func (user *WxUser) TableModel() string {

	return "wx_user"
}
