package mascotRelayer

import (
	"buzzriKamiAPI/noticeRelayer"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// シューティングゲームの結果を格納する構造体
type Talk struct {
	Num int `json:"Num"`
}

// バズり値算出のために必要な変数
var (
	talk     Talk
	clientId int
	talkNum  int
)

// クライアントからシューティングのヒット数を受け取る関数
func RecvTalk(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqClientId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}
	reqTalkNum, err := strconv.Atoi(c.QueryParam("num"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for Num.")
	}

	// クライアントから取得した値の代入
	clientId = reqClientId
	talkNum = reqTalkNum

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, noticeRelayer.Result{Status: true})
}

// クライアントにバズり値を送る関数
func SendTalk(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}

	// IDが一致すれば、送信する値の代入
	if reqId == clientId {
		talk.Num = talkNum
		talkNum = 0
		clientId = 0
	} else {
		talk.Num = 0
	}

	// クライアントにバズり度を送信
	return c.JSON(http.StatusOK, talk)
}
