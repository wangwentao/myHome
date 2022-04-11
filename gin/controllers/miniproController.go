package controllers

import (
	"github.com/gin-gonic/gin"
	"myHome/gin/configs"
	"myHome/gin/models"
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

	au := configs.MiniPro.GetAuth()
	res, err := au.Code2Session(code)
	logs.Error(err).Msg("Use jsCode Call Code2Session function")

	sek := &models.WxUserSession{ResCode2Session: res}
	sid := services.MiniProLogin(c, sek)

	services.RemoveExpeiredSession(c, osid)

	logs.Info().Msg("Mini progrogram login end")
	c.JSON(http.StatusOK, gin.H{
		"sessionId": sid,
	})
}

func UserProfile(c *gin.Context) {

	sid := c.GetHeader("sessionId")
	rp := models.WxPutProfile{}
	err := c.BindJSON(&rp)
	logs.Error(err).Msg("Bind request json to struct")

	oid, sek := services.FindWxSessionKey(c, sid)

	// decrypt encrypted data
	enc := configs.MiniPro.GetEncryptor()
	data, deer := enc.Decrypt(sek, rp.UserProfile.EncryptedData, rp.UserProfile.Iv)
	logs.Error(deer).Msgf("Decrypt data: %+v", data)

	// save user profile
	user := models.WxUser{}
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
