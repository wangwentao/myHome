package controllers

import (
	"github.com/gin-gonic/gin"
	"myHome/gin/configs"
	"myHome/gin/models/webchat"
	"myHome/gin/services"
	"myHome/gin/utils"
	"myHome/gin/utils/logs"
	"net/http"
)

func MiniLogin(c *gin.Context) {

	osid := c.GetHeader("sessionId")
	logs.Trace().Msgf("header old sessionId : %s", osid)
	code := c.Query("code")
	logs.Trace().Msgf("miniprogram jsCOde : %s", code)

	sid := services.MiniProLogin(c, code, osid)

	logs.Info().Msg("Mini progrogram login end")
	c.JSON(http.StatusOK, gin.H{
		"sessionId": sid,
	})
}

func UserProfile(c *gin.Context) {

	sid := c.GetHeader("sessionId")
	rp := webchat.WxPutProfile{}
	err := c.BindJSON(&rp)
	logs.Error(err).Msg("Bind request json to struct")

	oid, sek := services.FindWxSessionKey(c, sid)

	// decrypt encrypted data
	enc := configs.MiniPro.GetEncryptor()
	data, deer := enc.Decrypt(sek, rp.UserProfile.EncryptedData, rp.UserProfile.Iv)
	logs.Error(deer).Msgf("Decrypt data: %+v", data)

	// save user profile
	user := webchat.WxUser{}
	utils.Copy(&user, &data)
	if len(user.OpenID) == 0 {
		user.OpenID = oid
	}

	err = services.SaveUserProfile(&user)
	logs.Error(err).Msg("Save user profile.")

	c.JSON(http.StatusOK, gin.H{
		"Post Message": "Successful!",
	})

}
