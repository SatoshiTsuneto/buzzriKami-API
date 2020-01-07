package omikujiRelayer

import (
	"buzzriKamiAPI/noticeRelayer"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// シューティングゲームの結果を格納する構造体
type Omikuji struct {
	Result int `json:"Result"`
}

// バズり値算出のために必要な変数
var (
	omikuji    Omikuji
	clientId   int
	omikujiNum int
)

// クライアントからシューティングのヒット数を受け取る関数
func RecvOmikuji(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqClientId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}
	reqOmikujiNum, err := strconv.Atoi(c.QueryParam("result"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for Result.")
	}

	// クライアントから取得した値の代入
	clientId = reqClientId
	omikujiNum = reqOmikujiNum

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, noticeRelayer.Result{Status: true})
}

// クライアントにバズり値を送る関数
func SendOmikuji(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}

	// IDが一致すれば、送信する値の代入
	if reqId == clientId {
		omikuji.Result = omikujiNum
		omikujiNum = 0
		clientId = 0
	} else {
		omikuji.Result = 0
	}

	// クライアントにバズり度を送信
	return c.JSON(http.StatusOK, omikuji)
}
