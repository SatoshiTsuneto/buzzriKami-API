package fileUploader

import (
	"buzzriKamiAPI/noticeRelayer"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

// 取得したファイルを保存する関数
func FileSave(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// フォームから送られてきたファイルを取得
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Get File")
	}

	// ファイルを開く
	data, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Open File.")
	}
	defer data.Close()

	// ファイルの作成
	dst, err := os.Create("./BuzzriKamiWeb/html/img/" + file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing File Create.")
	}
	defer dst.Close()

	// 取得したファイルの保存
	_, err = io.Copy(dst, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Save Picture.")
	}

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, noticeRelayer.Result{Status: true})
}
