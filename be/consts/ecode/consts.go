package ecode

import "net/http"

const (
	// E0000 認証エラー
	E0000 = "E0000"

	// E0001 リクエストエラー
	E0001 = "E0001"

	// E1001 アカウントロック中
	E1001 = "E1001"
)

var ErrorMap = map[string]int{
	E0000: http.StatusUnauthorized,
	E0001: http.StatusBadRequest,

	E1001: http.StatusUnauthorized,
}