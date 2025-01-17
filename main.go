package main

import (
	"buzzriKamiAPI/bulletRelayer"
	"buzzriKamiAPI/buzzRelayer"
	"buzzriKamiAPI/fileUploader"
	"buzzriKamiAPI/hitRelayer"
	"buzzriKamiAPI/mascotRelayer"
	"buzzriKamiAPI/noticeRelayer"
	"buzzriKamiAPI/omikujiRelayer"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/saveFile", fileUploader.FileSave)

	e.GET("/recvBullet", bulletRelayer.SendBullet)
	e.GET("/sendBullet", bulletRelayer.RecvBullet)

	e.GET("/recvHit", hitRelayer.SendHit)
	e.GET("/sendHit", hitRelayer.RecvHit)

	e.GET("/recvNotice", noticeRelayer.SendNotice)
	e.GET("/sendNotice", noticeRelayer.RecvNotice)

	e.GET("/recvBuzz", buzzRelayer.SendBuzz)
	e.GET("/sendBuzz", buzzRelayer.RecvBuzz)

	e.GET("/recvTalk", mascotRelayer.SendTalk)
	e.GET("/sendTalk", mascotRelayer.RecvTalk)

	e.GET("/recvAction", mascotRelayer.SendAction)
	e.GET("/sendAction", mascotRelayer.RecvAction)

	e.GET("/recvOmikuji", omikujiRelayer.SendOmikuji)
	e.GET("/sendOmikuji", omikujiRelayer.RecvOmikuji)

	e.Logger.Fatal(e.Start(":9999"))
}
