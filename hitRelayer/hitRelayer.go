package hitRelayer

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// 送信するデータの構造体
type Hit struct {
	Num int `json:"Hit"`
}

// データを一時保管する変数
var (
	hit      Hit
	clientId int
	hitNum   int
)

// クライアントから通知を受け取る関数
func RecvHit(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}
	reqHitNum, err := strconv.Atoi(c.QueryParam("hit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for Hit.")
	}

	// クライアントから取得した値の代入
	clientId = reqId
	hitNum = reqHitNum

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, "OK")
}

// クライアントに通知を送る関数
func SendHit(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}

	// 送信する値の代入
	if reqId == clientId {
		hit.Num = hitNum
		clientId = 0
	} else {
		hit.Num = 0
	}

	// クライアントに真偽値を送信
	return c.JSON(http.StatusOK, hit)
}
