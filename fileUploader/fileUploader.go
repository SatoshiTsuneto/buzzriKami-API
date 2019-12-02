package fileUploader

import (
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"strconv"
	"subServer/noticeController"
)

// 取得したファイルを保存する関数
func FileSave(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param.")
	}

	// 完了通知を送る先のIDを取得
	noticeController.ClientId = reqId

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
	dst, err := os.Create("./fileUploader/file/" + file.Filename)
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
	return c.JSON(http.StatusOK, "OK.")
}