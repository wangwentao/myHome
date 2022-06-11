package services

import (
	"context"
	"myHome/gin/configs"
	"myHome/gin/models/webchat"
	"myHome/gin/services/stores"
	"myHome/gin/utils/logs"
)

func MiniProLogin(ctx context.Context, jsCode string, osid string) string {

	au := configs.MiniPro.GetAuth()
	res, err := au.Code2Session(jsCode)
	logs.Error(err).Msg("Use jsCode Call Code2Session function")

	sek := &webchat.WxUserSession{ResCode2Session: res}
	sid := stores.StoreWxSession(ctx, sek)

	wxUser := &webchat.WxUser{
		OpenID:  sek.OpenID,
		UnionID: sek.UnionID,
	}

	err = stores.NewModel(wxUser)
	logs.Error(err).Msg("Create new user")

	//delete expeired user session id from redis
	if len(osid) > 0 {
		RemoveExpeiredSession(ctx, osid)
	}

	return sid
}

func SaveUserProfile(user *webchat.WxUser) error {

	err := stores.UpdateModel(user)
	return err
}

func SessionExist(ctx context.Context, sid string) bool {
	res := stores.ExistSessionID(ctx, sid)

	return res
}

func RemoveExpeiredSession(ctx context.Context, sid string) {

	logs.Trace().Msg("Remove mini program experied session from reids")
	if SessionExist(ctx, sid) {
		stores.DelExpeiredSession(ctx, sid)
	}
}

func FindWxSessionKey(ctx context.Context, sid string) (o string, s string) {

	wxse := stores.FindWxSession(ctx, sid)

	return wxse.OpenID, wxse.SessionKey
}
