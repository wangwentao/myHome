package services

import (
	"context"
	"myHome/gin/models"
	"myHome/gin/services/stores"
	"myHome/gin/utils"
	"myHome/gin/utils/logs"
)

func MiniProLogin(ctx context.Context, s *models.WxUserSession) string {

	sid := stores.StoreWxSession(ctx, s)

	wxUser := &models.WxUser{
		OpenID:  s.OpenID,
		UnionID: s.UnionID,
	}
	err := stores.NewWxUser(wxUser)
	utils.CheckErr(err)

	return sid
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

func SaveUserProfile(user *models.WxUser) error {

	err := stores.UpdateWxUser(user)
	return err
}

func FindWxSessionKey(ctx context.Context, sid string) (o string, s string) {

	wxse := stores.FindWxSession(ctx, sid)

	return wxse.OpenID, wxse.SessionKey
}
