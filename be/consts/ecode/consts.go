package ecode

import "net/http"

const (
	// E0000 認証エラー
	E0000 = "E0000"

	// E0001 リクエストエラー
	E0001 = "E0001"
)

var ErrorMap = map[string]int{
	E0000: http.StatusUnauthorized,
	E0001: http.StatusBadRequest,
}