package wechat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/menu"
	"github.com/silenceper/wechat/message"
)

const (
	AppID="wxf1d7f3e6d462265f"
	AppSecret="23573a32c19bba33f299627875b2aaea"
	Token= "fatjiong"
)

var wx = &wechat.Wechat{}

//配置微信参数
func init(){
	config := &wechat.Config{
		AppID:          AppID,
		AppSecret:      AppSecret,
		Token:          Token,
	}
	wx = wechat.NewWechat(config)
}

	// 处理那啥
	func DoHandler(c *gin.Context){
		doMessage(c)
	}

	func doMessage(c *gin.Context){
		// 传入request和responseWriter
		server := wx.GetServer(c.Request, c.Writer)
		//设置接收消息的处理方法
		server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
			//回复消息：演示回复用户发送的消息
			text := message.NewText(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		})

		server.Serve()
		server.Send()
	}


	//创建自定义菜单
	func doMenu(c *gin.Context){
		mu := wx.GetMenu()
		buttons := make([]*menu.Button, 1)
		btn := new(menu.Button)

		//创建click类型菜单
		btn.SetClickButton("name", "key123")
		buttons[0] = btn

		//设置btn为二级菜单
		btn2 := new(menu.Button)
		btn2.SetSubButton("subButton", buttons)

		buttons2 := make([]*menu.Button, 1)
		buttons2[0] = btn2

		//发送请求
		err := mu.SetMenu(buttons2)
		if err != nil {
			fmt.Printf("err= %v", err)
			return
		}

	}