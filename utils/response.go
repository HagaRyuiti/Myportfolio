package utils

import (
	"encoding/json"
	"net/http"
)

// JSONレスポンスを返すためのヘルパー関数
func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "レスポンスのエンコードに失敗しました", http.StatusInternalServerError)
	}
}
