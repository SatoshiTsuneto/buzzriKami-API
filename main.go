package main

import (
	"buzzriKamiAPI/bulletRelayer"
	"buzzriKamiAPI/buzzRelayer"
	"buzzriKamiAPI/noticeRelayer"
	"github.com/labstack/echo"
	"subServer/fileUploader"
)

func main() {
	e := echo.New()

	e.POST("/saveFile", fileUploader.FileSave)

	e.GET("/recvBullet", bulletRelayer.SendBullet)
	e.GET("/sendBullet", bulletRelayer.RecvBullet)

	e.GET("/recvNotice", noticeRelayer.SendNotice)
	e.GET("/sendNotice", noticeRelayer.RecvNotice)

	e.GET("/recvBuzz", buzzRelayer.SendBuzz)
	e.GET("/sendBuzz", buzzRelayer.RecvBuzz)

	e.Logger.Fatal(e.Start(":9999"))
}
