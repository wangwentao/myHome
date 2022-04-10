package main

import (
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"myHome/gin/models"
	"myHome/gin/utils"
)

func main() {

	/*fmt.Println("Utils Test")
	fmt.Println("UUID: "+utils.GenUUID(16))
	fmt.Println("GET UUID: "+utils.GetUUID())*/

	data := encryptor.PlainData{
		OpenID:    "123",
		UnionID:   "456",
		NickName:  "wwt",
		Gender:    2,
		Country:   "China",
		City:      "Shanghai",
		AvatarURL: "---",
	}

	user := models.WxUser{}

	utils.Copy(&user, &data)

	fmt.Println(data)
	fmt.Println(user)
}
