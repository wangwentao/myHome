package stores

import (
	"context"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"myHome/gin/configs"
	"myHome/gin/models/webchat"
	"myHome/gin/utils"
	"myHome/gin/utils/logs"
)

const (
	wxOpenID     = "openid"
	wxSessionKey = "sessionkey"
)

func StoreWxSession(ctx context.Context, s *webchat.WxUserSession) string {

	sid := utils.GetUUID()

	m := make(map[string]interface{})
	m[wxOpenID] = s.OpenID
	m[wxSessionKey] = s.SessionKey

	_, err := configs.RedisStore.HMSet(ctx, sid, m).Result()
	logs.Error(err).Msg("Store wechat openid and sessionkey")

	return sid
}

func FindWxSession(ctx context.Context, sid string) *webchat.WxUserSession {

	openid, err := configs.RedisStore.HGet(ctx, sid, wxOpenID).Result()
	utils.CheckErr(err)
	sek, serr := configs.RedisStore.HGet(ctx, sid, wxSessionKey).Result()
	utils.CheckErr(serr)

	wxse := auth.ResCode2Session{
		OpenID:     openid,
		SessionKey: sek,
	}

	return &webchat.WxUserSession{ResCode2Session: wxse}
}

func DelExpeiredSession(ctx context.Context, sid string) {

	err := configs.RedisStore.HDel(ctx, sid, wxOpenID).Err()
	utils.CheckErr(err)
	err = configs.RedisStore.HDel(ctx, sid, wxSessionKey).Err()
	utils.CheckErr(err)
}

func ExistSessionID(ctx context.Context, sid string) bool {
	res, err := configs.RedisStore.Exists(ctx, sid).Result()
	utils.CheckErr(err)

	return res > 0
}
