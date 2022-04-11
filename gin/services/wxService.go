package services

import (
	"context"
	"myHome/gin/models"
	"myHome/gin/services/stores"
	"myHome/gin/utils/logs"
)

func MiniProLogin(ctx context.Context, s *models.WxUserSession) string {

	sid := stores.StoreWxSession(ctx, s)

	wxUser := &models.WxUser{
		OpenID:  s.OpenID,
		UnionID: s.UnionID,
	}

	err := stores.NewModel(wxUser)
	logs.Error(err).Msg("Create new user")

	return sid
}

func SaveUserProfile(user *models.WxUser) error {

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
