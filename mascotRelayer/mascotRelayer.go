package mascotRelayer

import (
	"buzzriKamiAPI/noticeRelayer"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// マスコットのトーク番号を格納する構造体
type Talk struct {
	Num int `json:"Talk"`
}

// マスコットのトーク番号を格納する構造体
type Action struct {
	Num int `json:"Action"`
}

// バズり値算出のために必要な変数
var (
	talk           Talk
	action         Action
	clientTalkId   int
	talkNum        int
	clientActionId int
	actionNum      int
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
	reqTalkNum, err := strconv.Atoi(c.QueryParam("talk"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for Talk.")
	}

	// クライアントから取得した値の代入
	clientTalkId = reqClientId
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
	if reqId == clientTalkId {
		talk.Num = talkNum
		talkNum = 0
		clientTalkId = 0
	} else {
		talk.Num = 0
	}

	// クライアントにバズり度を送信
	return c.JSON(http.StatusOK, talk)
}

// クライアントからシューティングのヒット数を受け取る関数
func RecvAction(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqClientId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}
	reqActionNum, err := strconv.Atoi(c.QueryParam("action"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for Action.")
	}

	// クライアントから取得した値の代入
	clientActionId = reqClientId
	actionNum = reqActionNum

	// 完了をクライアントに送信
	return c.JSON(http.StatusOK, noticeRelayer.Result{Status: true})
}

// クライアントにバズり値を送る関数
func SendAction(c echo.Context) error {
	// JSのためのヘッダー設定
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	// クエリーの取得および、数値への変換
	reqId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Missing Query Param for ID.")
	}

	// IDが一致すれば、送信する値の代入
	if reqId == clientActionId {
		action.Num = actionNum
		actionNum = 0
		clientActionId = 0
	} else {
		action.Num = 0
	}

	// クライアントにバズり度を送信
	return c.JSON(http.StatusOK, action)
}
