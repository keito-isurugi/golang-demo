package middleware

import (
	"net/http"
)

// CORSエラー回避用ミドルウェア
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// クロスオリジン用にセット
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")
    w.Header().Set("Content-Type", "application/json")

		// preflight用に200でいったん返す
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}