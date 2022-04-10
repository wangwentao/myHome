package stores

import (
	"context"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"gorm.io/gorm/clause"
	"myHome/gin/configs"
	"myHome/gin/models"
	"myHome/gin/utils"
)

const (
	wxOpenID     = "openid"
	wxSessionKey = "sessionkey"
)

// Postresql

func NewWxUser(user *models.WxUser) error {

	db := configs.PGStore.Table(user.TableModel()).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(user)

	return db.Error
}

func UpdateWxUser(user *models.WxUser) error {

	err := configs.PGStore.Table(user.TableModel()).Save(user).Error
	utils.CheckErr(err)
	return err
}

// Redis

func StoreWxSession(ctx context.Context, s *models.WxUserSession) string {

	sid := utils.GetUUID()

	m := make(map[string]interface{})
	m[wxOpenID] = s.OpenID
	m[wxSessionKey] = s.SessionKey

	_, err := configs.RedisStore.HMSet(ctx, sid, m).Result()
	utils.PanicError(err)

	return sid
}

func FindWxSession(ctx context.Context, sid string) *models.WxUserSession {

	openid, err := configs.RedisStore.HGet(ctx, sid, wxOpenID).Result()
	utils.CheckErr(err)
	sek, serr := configs.RedisStore.HGet(ctx, sid, wxSessionKey).Result()
	utils.CheckErr(serr)

	wxse := auth.ResCode2Session{
		OpenID:     openid,
		SessionKey: sek,
	}

	return &models.WxUserSession{ResCode2Session: wxse}
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
