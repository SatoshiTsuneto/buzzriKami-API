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

type Result struct {
	Status bool `json:"Status"`
}

// データを一時保管する変数
var (
	notice   Notice
	clientId int
)

// クライアントから通知を受け取る関数
func RecvNotice(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}

	// クライアントから取得した値の代入
	clientId = reqId

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, Result{Status: true})
}

// クライアントに通知を送る関数
func SendNotice(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}

	// 送信する値の代入
	if reqId == clientId {
		notice.Flg = true
		clientId = 0
	} else {
		notice.Flg = false
	}

	// クライアントに真偽値を送信
	return c.JSON(http.StatusOK, notice)
}
