package noticeRelayer

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// 送信するデータの構造体
type Notice struct {
	Flg bool `json:"Flg"`
}

// データを一時保管する変数
var (
	notice   Notice
	ClientId int
)

// クライアントから通知を受け取る関数
func RecvNotice(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param.")
	}

	// クライアントから取得した値の代入
	ClientId = reqId

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, "OK")
}

// クライアントに通知を送る関数
func SendNotice(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param.")
	}

	// 送信する値の代入
	if reqId == ClientId {
		notice.Flg = true
		ClientId = 0
	} else {
		notice.Flg = false
	}

	// クライアントに真偽値を送信
	return c.JSON(http.StatusOK, notice)
}
