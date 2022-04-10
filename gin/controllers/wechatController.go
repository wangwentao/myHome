package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)
import (
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func ServeWeChat(c *gin.Context) {

	wc := wechat.NewWechat()
	memory := cache.NewMemory()

	// account config
	/*cfg := &offConfig.Config{
		AppID:     "wx3bafe99efb98e453",
		AppSecret: "841d559635b33fef1677c33ffb7f5ef7",
		Token:     "mysit123456",
		//EncodingAESKey: "rvEyZH8ZhQNFvJJtew1yomssPohKSDCUI3qm8AOnDUT",
		Cache: memory,
	}*/
	// Test
	cfg := &offConfig.Config{
		AppID:     "wx6aed735c1871fc49",
		AppSecret: "1429523ac188a50418b1a34a6e5b1785",
		Token:     "mysit123456",
		//EncodingAESKey: "rvEyZH8ZhQNFvJJtew1yomssPohKSDCUI3qm8AOnDUT",
		Cache: memory,
	}

	// account object
	officialAccount := wc.GetOfficialAccount(cfg)

	// 传入request和responseWriter
	server := officialAccount.GetServer(c.Request, c.Writer)

	// 自定义菜单
	m := officialAccount.GetMenu()

	btn001 := menu.NewClickButton("Demo", "BTN001")
	btn002 := menu.NewClickButton("Show", "BTN002")

	buttons := []*menu.Button{btn001, btn002}
	errM := m.SetMenu(buttons)

	/*errM := m.SetMenuByJSON(
			`{
			 "button":[
			 {
				  "type":"click",
				  "name":"今日歌曲",
				  "key":"V1001_TODAY_MUSIC"
			  },
			  {
				   "name":"菜单",
				   "sub_button":[
				   {
					   "type":"view",
					   "name":"搜索",
					   "url":"http://www.soso.com/"
					},
					{
					   "type":"click",
					   "name":"赞一下我们",
					   "key":"V1001_GOOD"
					}]
			   }]
	 		}`)*/

	if errM != nil {
		fmt.Println(errM)
		return
	}

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {

		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	errS := server.Send()
	if errS != nil {
		fmt.Println(errS)
		return
	}
}
