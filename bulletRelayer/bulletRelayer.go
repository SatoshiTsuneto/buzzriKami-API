package bulletRelayer

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// 送信するデータの構造体
type Bullet struct {
	Num int `json:"Bullet"`
}

// 弾の数を保存する変数
var (
	bullet    Bullet
	clientId  int
	bulletNum int
)

// 銃弾の数を受け取る関数
func RecvBullet(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqClientId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}
	reqBulletNum, err := strconv.Atoi(c.QueryParam("bullet"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for Bullet.")
	}

	// クライアントから取得した値の代入
	clientId = reqClientId
	bulletNum = reqBulletNum

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, "OK")
}

// クライアントに銃弾の数を返す関数
func SendBullet(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}

	// IDが一致すれば、送信する値の代入
	if reqId == clientId {
		bullet.Num = bulletNum
		bulletNum = 0
		clientId = 0
	} else {
		bullet.Num = 0
	}

	// クライアントに弾数を送信
	return c.JSON(http.StatusOK, bullet)
}
