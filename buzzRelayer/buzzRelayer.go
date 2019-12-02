package buzzRelayer

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// シューティングゲームの結果を格納する構造体
type Buzz struct {
	Num int `json:"Buzz"`
}

// バズり値算出のために必要な変数
var (
	buzz     Buzz
	clientId int
	buzzNum  int
)

// クライアントからシューティングのヒット数を受け取る関数
func RecvBuzz(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqClientId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param Frm ID.")
	}
	reqBuzzNum, err := strconv.Atoi(c.QueryParam("buzz"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param From Buzz.")
	}

	// クライアントから取得した値の代入
	clientId = reqClientId
	buzzNum = reqBuzzNum

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, "OK")
}

// クライアントにバズり値を送る関数
func SendBuzz(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param.")
	}

	// IDが一致すれば、送信する値の代入
	if reqId == clientId {
		buzz.Num = buzzNum
		clientId = 0
	} else {
		buzzNum = 0
	}

	// クライアントにバズり度を送信
	return c.JSON(http.StatusOK, buzz)
}
